package main

import (
	"context"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/cenkalti/backoff"
	docker "github.com/docker/docker/client"
	"gopkg.in/yaml.v2"
)

const ecrPublicRegion = "us-east-1"

var (
	config          Config
	backoffSettings *backoff.ExponentialBackOff
)

type Config struct {
	SourceRepos []Repository `yaml:"sourceRepos"`
	TargetRepos []Repository `yaml:"targetRepos"`
}

type Repository struct {
	Registry    string   `yaml:"registry"`
	Name        string   `yaml:"name"`
	Host        string   `yaml:"host"`
	AllowedTags []string `yaml:"allowed_tags"`
}

// worker is the mirror working function.
func worker(ctx context.Context, wg *sync.WaitGroup, workerCh chan []Repository, dc *docker.Client, em *ecrManager) {
	for repos := range workerCh {
		log.Printf("Starting to mirror repo %s/%s/%s to repo %s/%s", repos[0].Host, repos[0].Registry, repos[0].Name, repos[1].Registry, repos[1].Name)

		m := mirror{
			ctx:          ctx,
			dockerClient: dc,
			ecrManager:   em,
		}
		if err := m.setup(repos); err != nil {
			log.Printf("Failed to setup mirror for repo %s/%s/%s: %v", repos[0].Host, repos[0].Registry, repos[0].Name, err)
			wg.Done()
			continue
		}

		m.work()
		log.Printf("From repo %s/%s/%s to repo %s/%s mirror completed", repos[0].Host, repos[0].Registry, repos[0].Name, repos[1].Registry, repos[1].Name)
		wg.Done()
	}
}

func main() {
	ctx := context.Background()
	content, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalf("Could not read config file: %v", err)
	}

	if err = yaml.Unmarshal(content, &config); err != nil {
		log.Fatalf("Could not parse config file: %v", err)
	}

	log.Println("Creating Docker client")
	dc, err := docker.NewClientWithOpts(docker.FromEnv, docker.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Could not create Docker client: %v", err)
	}

	info, err := dc.Info(ctx)
	if err != nil {
		log.Fatalf("Could not get Docker info: %v", err)
	}
	log.Printf("Connected to Docker daemon: %s @ %s", info.Name, info.ServerVersion)

	log.Println("Creating AWS client")
	// Override the AWS region with the ecrPublicRegion for ECR authentication.
	cfg, err := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(ecrPublicRegion))
	if err != nil {
		log.Fatalf("Unable to load AWS SDK config: %v", err)
	}

	em := &ecrManager{client: ecrpublic.NewFromConfig(cfg)}
	backoffSettings = backoff.NewExponentialBackOff()
	backoffSettings.InitialInterval = 1 * time.Second
	backoffSettings.MaxElapsedTime = 10 * time.Second
	notifyError := func(err error, d time.Duration) {
		log.Fatalf("%v (%s)", err, d.String())
	}

	// Build ECR registries cache.
	if err = backoff.RetryNotify(em.buildCacheBackoff(ctx), backoffSettings, notifyError); err != nil {
		log.Fatalf("Could not build ECR cache: %v", err)
	}

	token, err := em.getAuthToken(ctx)
	if err != nil {
		log.Fatalf("Could not retrieve authentication token: %v", err)
	}
	credentials, err := getDockerCredentials(token)
	if err != nil {
		log.Fatalf("Could not parse authentication token: %v", err)
	}
	em.registryAuth = credentials

	workersNum := runtime.NumCPU()

	workerCh := make(chan []Repository, 5)
	var wg sync.WaitGroup

	for i := 0; i < workersNum; i++ {
		go worker(ctx, &wg, workerCh, dc, em)
	}

	// Add source repo and target repo pair to worker channel.
	for i := range config.SourceRepos {
		wg.Add(1)
		workerCh <- []Repository{config.SourceRepos[i], config.TargetRepos[i]}
	}

	wg.Wait()
	log.Println("Finished mirroring repositories")
}
