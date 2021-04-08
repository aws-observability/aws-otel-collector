// Copyright  OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ecsobserver

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/hashicorp/golang-lru/simplelru"

	"go.uber.org/zap"
)

type Fetcher interface {
	// FetcherAndDecorate fetches all the tasks and attach addation information
	// like definition, serivces and container instances.
	FetchAndDecorate(ctx context.Context) ([]*Task, error)
}

const (
	// ECS Service Quota: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-quotas.html
	taskDefCacheSize               = 2000
	ec2CacheSize                   = 2000
	describeServiceLimit           = 10
	describeContainerInstanceLimit = 100
	// NOTE: these constants are not defined in go sdk, there are three values for deployment status.
	deploymentStatusActive   = "ACTIVE"
	deploymentStatusPrimary  = "PRIMARY"
	deploymentStatusInActive = "INACTIVE"
)

// ec2Client includes API required by TaskFetcher.
type ec2Client interface {
	DescribeInstancesWithContext(ctx context.Context, input *ec2.DescribeInstancesInput, opts ...request.Option) (*ec2.DescribeInstancesOutput, error)
}

// ecsClient includes API required by TaskFetcher.
type ecsClient interface {
	ListTasksWithContext(ctx context.Context, input *ecs.ListTasksInput, opts ...request.Option) (*ecs.ListTasksOutput, error)
	DescribeTasksWithContext(ctx context.Context, input *ecs.DescribeTasksInput, opts ...request.Option) (*ecs.DescribeTasksOutput, error)
	DescribeTaskDefinitionWithContext(ctx context.Context, input *ecs.DescribeTaskDefinitionInput, opts ...request.Option) (*ecs.DescribeTaskDefinitionOutput, error)
	ListServicesWithContext(ctx context.Context, input *ecs.ListServicesInput, opts ...request.Option) (*ecs.ListServicesOutput, error)
	DescribeServicesWithContext(ctx context.Context, input *ecs.DescribeServicesInput, opts ...request.Option) (*ecs.DescribeServicesOutput, error)
	DescribeContainerInstancesWithContext(ctx context.Context, input *ecs.DescribeContainerInstancesInput, opts ...request.Option) (*ecs.DescribeContainerInstancesOutput, error)
}

type TaskFetcher struct {
	logger            *zap.Logger
	ecs               ecsClient
	ec2               ec2Client
	cluster           string
	serviceNameFilter ServiceNameFilter
	taskDefCache      simplelru.LRUCache
	ec2Cache          simplelru.LRUCache
}

type TaskFetcherOptions struct {
	Logger            *zap.Logger
	Cluster           string
	Region            string
	ServiceNameFilter ServiceNameFilter
}

func NewTaskFetcher(opts TaskFetcherOptions) (*TaskFetcher, error) {
	logger := opts.Logger
	// Init cache
	taskDefCache, err := simplelru.NewLRU(taskDefCacheSize, nil)
	if err != nil {
		return nil, err
	}
	ec2Cache, err := simplelru.NewLRU(ec2CacheSize, nil)
	if err != nil {
		return nil, err
	}

	// AWS client
	if opts.Region == "" {
		return nil, fmt.Errorf("missing aws region for task fetcher")
	}
	logger.Debug("Init TaskFetcher", zap.String("Region", opts.Region), zap.String("Cluster", opts.Cluster))
	awsCfg := aws.NewConfig().WithRegion(opts.Region).WithCredentialsChainVerboseErrors(true)
	sess, err := session.NewSession(awsCfg)
	if err != nil {
		return nil, fmt.Errorf("create aws session failed: %w", err)
	}

	return &TaskFetcher{
		logger:            opts.Logger,
		ecs:               ecs.New(sess, awsCfg),
		ec2:               ec2.New(sess, awsCfg),
		cluster:           opts.Cluster,
		serviceNameFilter: opts.ServiceNameFilter,
		taskDefCache:      taskDefCache,
		ec2Cache:          ec2Cache,
	}, nil
}

func (f *TaskFetcher) FetchAndDecorate(ctx context.Context) ([]*Task, error) {
	// Task
	rawTasks, err := f.GetAllTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("getAllTasks failed: %w", err)
	}
	tasks, err := f.AddTaskDefinition(ctx, rawTasks)
	if err != nil {
		return nil, fmt.Errorf("addTaskDefinition failed: %w", err)
	}

	// EC2
	if err = f.AddContainerInstance(ctx, tasks); err != nil {
		return nil, fmt.Errorf("addContainerInstance failed: %w", err)
	}

	// Service
	services, err := f.GetAllServices(ctx, f.serviceNameFilter)
	if err != nil {
		return nil, fmt.Errorf("getAllServices failed: %w", err)
	}
	if err := f.AddService(tasks, services); err != nil {
		return nil, fmt.Errorf("addService failed: %w", err)
	}
	f.logger.Info("FetchAndDecorate completed", zap.Int("TasksCount", len(tasks)))
	return tasks, nil
}

// GetAllTasks get arns of all running tasks and describe those tasks.
func (f *TaskFetcher) GetAllTasks(ctx context.Context) ([]*ecs.Task, error) {
	svc := f.ecs
	cluster := aws.String(f.cluster)
	req := ecs.ListTasksInput{Cluster: cluster}
	var tasks []*ecs.Task
	for {
		listRes, err := svc.ListTasksWithContext(ctx, &req)
		if err != nil {
			return nil, fmt.Errorf("ecs.ListTasks failed: %w", err)
		}
		// NOTE: the limit for list task response and describe task request are both 100.
		descRes, err := svc.DescribeTasksWithContext(ctx, &ecs.DescribeTasksInput{
			Cluster: cluster,
			Tasks:   listRes.TaskArns,
		})
		if err != nil {
			return nil, fmt.Errorf("ecs.DescribeTasks failed: %w", err)
		}
		tasks = append(tasks, descRes.Tasks...)
		if listRes.NextToken == nil {
			break
		}
		req.NextToken = listRes.NextToken
	}
	return tasks, nil
}

// GetAllServices lists all the serivces but only describes those pass the filter.
// The filter is derived from ServiceConfig.
func (f *TaskFetcher) GetAllServices(ctx context.Context, nameFilter ServiceNameFilter) ([]*ecs.Service, error) {
	if f.serviceNameFilter == nil {
		return nil, fmt.Errorf("serviceNameFilter is nil")
	}
	svc := f.ecs
	cluster := aws.String(f.cluster)
	// List and filter out services we need to desribe.
	listReq := ecs.ListServicesInput{Cluster: cluster}
	var servicsToDescribe []*string
	for {
		res, err := svc.ListServicesWithContext(ctx, &listReq)
		if err != nil {
			return nil, err
		}
		for _, arn := range res.ServiceArns {
			segs := strings.Split(aws.StringValue(arn), "/")
			name := segs[len(segs)-1]
			// Only keep service specified in config.
			if nameFilter(name) {
				servicsToDescribe = append(servicsToDescribe, arn)
			}
		}
		if res.NextToken == nil {
			break
		}
		listReq.NextToken = res.NextToken
	}

	// DescribeServices size limit is 10 so we need to do paging on client side.
	var services []*ecs.Service
	for i := 0; i < len(servicsToDescribe); i += describeServiceLimit {
		end := minInt(i+describeServiceLimit, len(servicsToDescribe))
		desc := &ecs.DescribeServicesInput{
			Cluster:  cluster,
			Services: servicsToDescribe[i:end],
		}
		res, err := svc.DescribeServicesWithContext(ctx, desc)
		if err != nil {
			return nil, fmt.Errorf("ecs.DescribeServices failed %w", err)
		}
		services = append(services, res.Services...)
	}
	return services, nil
}

func (f *TaskFetcher) AddTaskDefinition(ctx context.Context, tasks []*ecs.Task) ([]*Task, error) {
	svc := f.ecs
	// key is task definition arn
	arn2Def := make(map[string]*ecs.TaskDefinition)
	for _, t := range tasks {
		arn2Def[aws.StringValue(t.TaskDefinitionArn)] = nil
	}

	for arn := range arn2Def {
		if arn == "" {
			continue
		}
		var def *ecs.TaskDefinition
		if cached, ok := f.taskDefCache.Get(arn); ok {
			def = cached.(*ecs.TaskDefinition)
		} else {
			res, err := svc.DescribeTaskDefinitionWithContext(ctx, &ecs.DescribeTaskDefinitionInput{
				TaskDefinition: aws.String(arn),
			})
			if err != nil {
				return nil, err
			}
			f.taskDefCache.Add(arn, res.TaskDefinition)
			def = res.TaskDefinition
		}
		arn2Def[arn] = def
	}

	var tasksWithDef []*Task
	for _, t := range tasks {
		tasksWithDef = append(tasksWithDef, &Task{
			Task:       t,
			Definition: arn2Def[aws.StringValue(t.TaskDefinitionArn)],
		})
	}
	return tasksWithDef, nil
}

func (f *TaskFetcher) AddService(tasks []*Task, services []*ecs.Service) error {
	// Map deployment ID to service name
	idToService := make(map[string]*ecs.Service)
	for _, svc := range services {
		for _, deployment := range svc.Deployments {
			status := aws.StringValue(deployment.Status)
			if status == deploymentStatusActive || status == deploymentStatusPrimary {
				idToService[aws.StringValue(deployment.Id)] = svc
			}
		}
	}

	// Attach service to task
	for _, t := range tasks {
		// Task is created using RunTask i.e. not manged by a service.
		if t.Task.StartedBy == nil {
			continue
		}
		deploymentId := aws.StringValue(t.Task.StartedBy)
		svc := idToService[deploymentId]
		// NOTE: This could happen because we only fetch services defiend in ServiceConfig.
		// However, we fetch all the tasks, which could be started by other services
		// or started using RunTasks directly.
		if svc == nil {
			continue
		}
		t.Service = svc
	}
	return nil
}

// AddContainerInstance fetches all the container instances's underlying EC2 vms
// and attach EC2 info to tasks.
func (f *TaskFetcher) AddContainerInstance(ctx context.Context, tasks []*Task) error {
	// Map container instance to EC2, key is container instance id.
	ciToEC2 := make(map[string]*ec2.Instance)
	// Only EC2 instance type need to fetch EC2 info
	for _, t := range tasks {
		if aws.StringValue(t.Task.LaunchType) != ecs.LaunchTypeEc2 {
			continue
		}
		ciToEC2[aws.StringValue(t.Task.ContainerInstanceArn)] = nil
	}
	// All fargate, skip
	if len(ciToEC2) == 0 {
		return nil
	}

	// Describe container instances that do not have cached EC2 info.
	var instanceList []*string
	for instanceArn := range ciToEC2 {
		cached, ok := f.ec2Cache.Get(instanceArn)
		if ok {
			ciToEC2[instanceArn] = cached.(*ec2.Instance) // use value from cache
		} else {
			instanceList = append(instanceList, aws.String(instanceArn))
		}
	}
	sortStringPointers(instanceList)

	// DescibeCotnainerInstance size limit is 100, do it in batch.
	for i := 0; i < len(instanceList); i += describeContainerInstanceLimit {
		end := minInt(i+describeContainerInstanceLimit, len(instanceList))
		if err := f.describeContainerInstances(ctx, instanceList[i:end], ciToEC2); err != nil {
			return fmt.Errorf("describe container instanced failed offset=%d: %w", i, err)
		}
	}

	// Assign the info back to task
	for _, t := range tasks {
		// NOTE: we need to skip fargate here because we are looping all tasks again.
		// TODO: test it in fetcher by using a mixed cluster
		if aws.StringValue(t.Task.LaunchType) != ecs.LaunchTypeEc2 {
			continue
		}
		containerInstance := aws.StringValue(t.Task.ContainerInstanceArn)
		ec2Info, ok := ciToEC2[containerInstance]
		if !ok {
			return fmt.Errorf("container instance ec2 info not found containerInstnace=%q", containerInstance)
		}
		t.EC2 = ec2Info
	}

	// Update the cache
	for ci, ec2Info := range ciToEC2 {
		f.ec2Cache.Add(ci, ec2Info)
	}
	return nil
}

// Run ecs.DescribeContainerInstances and ec2.DescibeIsntances for a batch (less than 100 container instances).
func (f *TaskFetcher) describeContainerInstances(ctx context.Context, instanceList []*string,
	ci2EC2 map[string]*ec2.Instance) error {
	// Get container instances
	res, err := f.ecs.DescribeContainerInstancesWithContext(ctx, &ecs.DescribeContainerInstancesInput{
		Cluster:            aws.String(f.cluster),
		ContainerInstances: instanceList,
	})
	if err != nil {
		return fmt.Errorf("ecs.DescribeContainerInstance faile: %w", err)
	}

	// Create the index to map ec2 id back to container instance id.
	var ec2Ids []*string
	ec2IdToCI := make(map[string]string)
	for _, containerInstance := range res.ContainerInstances {
		ec2Id := containerInstance.Ec2InstanceId
		ec2Ids = append(ec2Ids, ec2Id)
		ec2IdToCI[aws.StringValue(ec2Id)] = aws.StringValue(containerInstance.ContainerInstanceArn)
	}

	// Fetch all ec2 instances and update mapping from container instance id to ec2 info.
	req := ec2.DescribeInstancesInput{InstanceIds: ec2Ids}
	for {
		res, err := f.ec2.DescribeInstancesWithContext(ctx, &req)
		if err != nil {
			return fmt.Errorf("ec2.DescribeInstances failed: %w", err)
		}
		for _, reservation := range res.Reservations {
			for _, instance := range reservation.Instances {
				if instance.InstanceId == nil {
					continue
				}
				ec2Id := aws.StringValue(instance.InstanceId)
				ci, ok := ec2IdToCI[ec2Id]
				if !ok {
					return fmt.Errorf("mapping from ec2 to container instance not found ec2=%s", ec2Id)
				}
				ci2EC2[ci] = instance // update mapping
			}
		}
		if res.NextToken == nil {
			break
		}
		req.NextToken = res.NextToken
	}
	return nil
}

// Util Start

func sortStringPointers(ps []*string) {
	var ss []string
	for _, p := range ps {
		ss = append(ss, aws.StringValue(p))
	}
	sort.Strings(ss)
	for i := range ss {
		ps[i] = aws.String(ss[i])
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Util End
