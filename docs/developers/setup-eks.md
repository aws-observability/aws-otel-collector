# Prerequisite

* [Install eksctl CLI](https://docs.aws.amazon.com/eks/latest/userguide/getting-started-eksctl.html)
* [Install kubectl CLI](https://docs.aws.amazon.com/eks/latest/userguide/install-kubectl.html)
* [Install AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html)

# Set UP EKS Cluster

## Create an EKS Cluster with Managed Nodes

1. Create SSH key pair (Optional)

```bash
aws ec2 create-key-pair --region {{region}} --key-name {{keypair}}
```

2. Create `cluster.yaml` file

```yaml
apiVersion: eksctl.io/v1alpha5
kind: ClusterConfig
metadata:
  name: {{cluster_name}}
  region: {{region}}
iam:
  withOIDC: true
nodeGroups:
  - name: mng-m5large
    instanceType: m5.large
    desiredCapacity: 2
    volumeSize: 100
    iam:
      attachPolicyARNs:
        - arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy
        - arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy
        - arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly
        - arn:aws:iam::aws:policy/CloudWatchAgentServerPolicy # CWAgent
        - arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore # SSM
        - arn:aws:iam::aws:policy/AWSXrayWriteOnlyAccess # xray
        - arn:aws:iam::aws:policy/AWSXRayDaemonWriteAccess # xray
        - arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess # s3
        - arn:aws:iam::aws:policy/AmazonPrometheusFullAccess
    ssh:
      allow: true
      publicKeyName: {{keypair}}
cloudWatch:
  clusterLogging:
    enableTypes: [ "*" ]
```

You can customize vpc to use an existing one.

```yaml
apiVersion: eksctl.io/v1alpha5
kind: ClusterConfig
metadata:
  name: {{cluster_name}}
  region: {{region}}
iam:
  withOIDC: true
vpc:
  subnets:
    public:
      us-west-2a: { id: {{subnet-id1}} }
      us-west-2b: { id: {{subnet-id2}} }
      us-west-2c: { id: {{subnet-id3}} }
nodeGroups:
  - name: mng-m5large
    instanceType: m5.large
    desiredCapacity: 2
    volumeSize: 100
    iam:
      attachPolicyARNs:
        - arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy
        - arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy
        - arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly
        - arn:aws:iam::aws:policy/CloudWatchAgentServerPolicy
        - arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore
        - arn:aws:iam::aws:policy/AWSXrayWriteOnlyAccess
        - arn:aws:iam::aws:policy/AWSXRayDaemonWriteAccess
        - arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess
        - arn:aws:iam::aws:policy/AmazonPrometheusFullAccess
    ssh:
      allow: true # will use ~/.ssh/id_rsa.pub as the default ssh key
      publicKeyName: {{keypair}}
cloudWatch:
  clusterLogging:
    enableTypes: [ "*" ]
```

3. Create EKS cluster

```bash
eksctl create cluster -f cluster.yaml
```

To create a cluster with fargate nodes, or look for other customization to your cluster,
see [Getting started with Amazon EKS – eksctl](https://docs.aws.amazon.com/eks/latest/userguide/getting-started-eksctl.html)
for more details.

## Use an existing EKS Cluster

1. Open the [Amazon EC2 console](https://console.aws.amazon.com/ec2/).
2. Select one of the worker node instances and choose the IAM role in the description.
3. On the IAM role page, choose **Attach policies**.
4. Attach the following policies to node instance role.
```bash
AmazonEKSWorkerNodePolicy
AmazonEKS_CNI_Policy
AmazonEC2ContainerRegistryReadOnly
CloudWatchAgentServerPolicy
AmazonSSMManagedInstanceCore
AWSXrayWriteOnlyAccess
AWSXRayDaemonWriteAccess
AmazonS3ReadOnlyAccess
AmazonPrometheusFullAccess
```