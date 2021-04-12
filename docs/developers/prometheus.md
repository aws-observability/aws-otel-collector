# Prometheus support in AWS-OTel-Collector

## Overview

AWS-OTel-Collector ships the following prometheus related components

- [prometheusreceiver](https://github.com/open-telemetry/opentelemetry-collector/tree/main/receiver/prometheusreceiver)
  scrape prometheus metrics
- [ecsobserver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/observer/ecsobserver)
  discover scrape targets from ecs tasks.
- [awsemfexporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/awsemfexporter)
  sends metric to cloudwatch and supports auto dashboard in [container insight](container-insight-install-aoc.md)
- [awsprometheusremotewriteexporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/awsprometheusremotewriteexporter)
  send metrics to [Amazon Managed Service for Grafana](https://aws.amazon.com/grafana/)

## Known issues

<!-- TODO(bjrara): mention job, instance, __name__ relabel -->