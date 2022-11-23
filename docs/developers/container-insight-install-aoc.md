# Overview

This document describes how to install ADOT Collector into your existing EKS cluster.

## Prerequisites

* [Set Up EKS Cluster](setup-eks.md)

## Install ADOT Collector

1. Set up variables to export metrics of your EKS cluster to the region where the logs should be published to.

```
export CLUSTER_NAME=<eks-cluster-name>
export AWS_REGION=<aws-region>
```

2. Deploy a standalone aws-otel-collector application. An example config template can be
   found [here](../../deployment-template/eks/otel-container-insights-prometheus.yaml).
    * Replace `{{region}}` with the name of the region where the logs are published (e.g. `us-west-2`).
    * Replace `{{cluster_name}}` with the actual eks cluster name.

```bash
cat otel-container-insights-prometheus.yaml |
sed "s/{{region}}/$AWS_REGION/g" | 
sed "s/{{cluster_name}}/$CLUSTER_NAME/g" |
kubectl apply -f - 
```

3. View the resources in the `aws-otel-eks` namespace.

```bash
kubectl get all -n aws-otel-eks
```
