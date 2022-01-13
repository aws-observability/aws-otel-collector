window.BENCHMARK_DATA = {
  "lastUpdate": 1642097435596,
  "repoUrl": "https://github.com/aws-observability/aws-otel-collector",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "name": "sethAmazon",
            "username": "sethAmazon",
            "email": "81644108+sethAmazon@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "f57ae1bc811622cc0ec86ae0dd096fcf9e7e1b4b",
          "message": "Remove version filter and trigger performance tests on all passing runs (#842)",
          "timestamp": "2022-01-03T20:56:28Z",
          "url": "https://github.com/aws-observability/aws-otel-collector/commit/f57ae1bc811622cc0ec86ae0dd096fcf9e7e1b4b"
        },
        "date": 1641326707403,
        "tool": "customSmallerIsBetter",
        "benches": [
          {
            "name": "logzio_exporter_trace_mock",
            "value": "3.70",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "89.18",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "6.29",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "72.07",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_mock",
            "value": "3.88",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_mock",
            "value": "71.02",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_trace",
            "value": "4.06",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_trace",
            "value": "71.80",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "3.18",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "69.51",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "xrayreceiver",
            "value": "4.71",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "xrayreceiver",
            "value": "141.62",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "3.77",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "129.18",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "2.36",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "75.18",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "5.57",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "77.92",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "3.61",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "84.03",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "25.33",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "85.90",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "28.93",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "458.96",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "33.18",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "74.23",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "xrayreceiver",
            "value": "34.16",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "xrayreceiver",
            "value": "579.24",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "28.50",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "661.41",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_trace",
            "value": "35.86",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_trace",
            "value": "77.83",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "28.39",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "101.91",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_mock",
            "value": "32.76",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_mock",
            "value": "76.31",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "27.65",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "71.60",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "16.82",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "152.34",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "28.55",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "495.05",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_trace",
            "value": "167.07",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_trace",
            "value": "12132.93",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "105.16",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "123.67",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "16.97",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "169.14",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "xrayreceiver",
            "value": "54.02",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "xrayreceiver",
            "value": "1044.25",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "124.05",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "3032.34",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "104.55",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "91.97",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "121.70",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "76.69",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_mock",
            "value": "140.71",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_mock",
            "value": "13436.02",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "131.67",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "82.44",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "7.87",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "256.43",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "statsd_mock",
            "value": "0.02",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "statsd_mock",
            "value": "59.30",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "59.78",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "statsd",
            "value": "25.83",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "statsd",
            "value": "68.83",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "59.82",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "prometheus_static",
            "value": "8.08",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "prometheus_static",
            "value": "258.53",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_metric",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_metric",
            "value": "58.81",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "60.19",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "60.01",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "60.27",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "prometheus_static",
            "value": "0.13",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "prometheus_static",
            "value": "72.42",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "0.13",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "72.59",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "60.79",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "59.62",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "61.09",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "59.95",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "statsd_mock",
            "value": "0.01",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "statsd_mock",
            "value": "59.91",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "58.32",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "statsd",
            "value": "0.56",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "statsd",
            "value": "67.37",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_metric",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_metric",
            "value": "58.77",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "60.00",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "statsd_mock",
            "value": "0.02",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "statsd_mock",
            "value": "58.89",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "60.98",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "58.67",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "62.48",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "prometheus_static",
            "value": "1.30",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "prometheus_static",
            "value": "116.67",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_metric",
            "value": "0.08",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_metric",
            "value": "59.66",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "statsd",
            "value": "5.17",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "statsd",
            "value": "68.96",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "59.89",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "1.31",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "115.16",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Khanh Nguyen",
            "username": "khanhntd",
            "email": "91758108+khanhntd@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "94d47a077e16882a635d15d9a7e296ac4e0e374d",
          "message": "Add all dev branches to CI (#861)",
          "timestamp": "2022-01-06T22:19:19Z",
          "url": "https://github.com/aws-observability/aws-otel-collector/commit/94d47a077e16882a635d15d9a7e296ac4e0e374d"
        },
        "date": 1641530286489,
        "tool": "customSmallerIsBetter",
        "benches": [
          {
            "name": "logzio_exporter_trace_mock",
            "value": "3.82",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "90.59",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "5.01",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "72.55",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_mock",
            "value": "4.22",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_mock",
            "value": "70.48",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_trace",
            "value": "4.84",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_trace",
            "value": "72.42",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "3.71",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "69.68",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "xrayreceiver",
            "value": "4.93",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "xrayreceiver",
            "value": "140.78",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "3.13",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "131.69",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "2.34",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "74.11",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "5.56",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "77.71",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "2.96",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "83.71",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "28.25",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "86.10",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "30.22",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "467.75",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "33.14",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "73.77",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "xrayreceiver",
            "value": "33.61",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "xrayreceiver",
            "value": "606.89",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "27.67",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "656.84",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_trace",
            "value": "35.84",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_trace",
            "value": "78.32",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "30.05",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "101.30",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_mock",
            "value": "34.01",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_mock",
            "value": "75.17",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "26.90",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "72.32",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "15.88",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "150.54",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "28.44",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "503.91",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_trace",
            "value": "174.77",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_trace",
            "value": "12741.63",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "116.64",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "123.15",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "16.19",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "165.45",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "xrayreceiver",
            "value": "54.14",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "xrayreceiver",
            "value": "974.94",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "113.63",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "3205.97",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "118.33",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "91.98",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "119.30",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "77.31",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_mock",
            "value": "140.25",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_mock",
            "value": "14075.07",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "128.56",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "82.28",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "8.26",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "259.11",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "statsd_mock",
            "value": "0.01",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "statsd_mock",
            "value": "58.11",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "58.16",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "statsd",
            "value": "26.77",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "statsd",
            "value": "68.71",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "58.73",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "prometheus_static",
            "value": "7.76",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "prometheus_static",
            "value": "263.53",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_metric",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_metric",
            "value": "59.37",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "59.03",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "59.22",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "60.50",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "prometheus_static",
            "value": "0.13",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "prometheus_static",
            "value": "72.68",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "0.13",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "73.12",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "58.03",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "62.72",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "58.43",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "60.05",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "statsd_mock",
            "value": "0.01",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "statsd_mock",
            "value": "58.83",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "58.28",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "statsd",
            "value": "0.73",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "statsd",
            "value": "68.70",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_metric",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_metric",
            "value": "58.35",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "60.48",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "statsd_mock",
            "value": "0.02",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "statsd_mock",
            "value": "60.34",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "58.54",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "59.90",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "59.80",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "prometheus_static",
            "value": "1.29",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "prometheus_static",
            "value": "116.69",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_metric",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_metric",
            "value": "58.46",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "statsd",
            "value": "5.27",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "statsd",
            "value": "68.27",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "59.96",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "1.34",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "115.50",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Khanh Nguyen",
            "username": "khanhntd",
            "email": "91758108+khanhntd@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "e5751cddec558235ab0eeda1f53e99d6c99d7b66",
          "message": "Add buildtool kit for ADOT collector image (#860)",
          "timestamp": "2022-01-11T17:30:51Z",
          "url": "https://github.com/aws-observability/aws-otel-collector/commit/e5751cddec558235ab0eeda1f53e99d6c99d7b66"
        },
        "date": 1641934409898,
        "tool": "customSmallerIsBetter",
        "benches": [
          {
            "name": "logzio_exporter_trace_mock",
            "value": "3.30",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "90.05",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "5.39",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "72.49",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_mock",
            "value": "4.73",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_mock",
            "value": "71.88",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "3.11",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "70.26",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "3.89",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "129.95",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "2.49",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "75.90",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "5.75",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "78.45",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "3.21",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "84.36",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "xrayreceiver_mock",
            "value": "4.55",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "xrayreceiver_mock",
            "value": "142.17",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "25.79",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "86.72",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "29.63",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "486.71",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "32.76",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "75.01",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "xrayreceiver_mock",
            "value": "25.83",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "xrayreceiver_mock",
            "value": "529.63",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "30.14",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "698.77",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "28.61",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "101.94",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_mock",
            "value": "36.44",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_mock",
            "value": "76.65",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "26.67",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "72.05",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "17.63",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "148.34",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "30.92",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "517.57",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "120.74",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "125.26",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "17.31",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "171.06",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "xrayreceiver_mock",
            "value": "38.39",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "xrayreceiver_mock",
            "value": "700.72",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "126.56",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "3137.77",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "101.48",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "90.46",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "117.80",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "76.41",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_mock",
            "value": "143.72",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_mock",
            "value": "15403.34",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "127.10",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "82.51",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "7.56",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "259.73",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "statsd_mock",
            "value": "0.02",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "statsd_mock",
            "value": "58.64",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_metric_mock",
            "value": "59.35",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "63.66",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "58.79",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "59.72",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "59.81",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "61.28",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "61.72",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "statsd_mock",
            "value": "0.02",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "statsd_mock",
            "value": "59.92",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "0.03",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "59.00",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "57.51",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "61.63",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "59.42",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "1.41",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "106.34",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_metric_mock",
            "value": "58.92",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "0.13",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "72.42",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "59.00",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "61.49",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "58.89",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_metric_mock",
            "value": "58.91",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "0.06",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "61.30",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "statsd_mock",
            "value": "0.01",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "statsd_mock",
            "value": "58.25",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "58.66",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Khanh Nguyen",
            "username": "khanhntd",
            "email": "91758108+khanhntd@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "543db00eb9f7f469d512df6f9ab643a4697df1c4",
          "message": "Fix condition for ssm (#874)\n\n* fix condition for ssm package\r\n\r\n* fix output dir in ssm manifest file\r\n\r\n* change name to clean-ssm-package",
          "timestamp": "2022-01-11T19:32:01Z",
          "url": "https://github.com/aws-observability/aws-otel-collector/commit/543db00eb9f7f469d512df6f9ab643a4697df1c4"
        },
        "date": 1642097435043,
        "tool": "customSmallerIsBetter",
        "benches": [
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "118.25",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "76.46",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "123.33",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "3143.67",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "30.55",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "507.55",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "109.82",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "124.79",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_mock",
            "value": "149.18",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_mock",
            "value": "16972.31",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "18.26",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "169.45",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "111.98",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "91.29",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "119.58",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "82.81",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "xrayreceiver_mock",
            "value": "37.15",
            "unit": "%",
            "extra": "Trace (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "xrayreceiver_mock",
            "value": "785.78",
            "unit": "MB",
            "extra": "Trace (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "3.32",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "70.36",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_mock",
            "value": "4.78",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_mock",
            "value": "71.65",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "3.01",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "91.62",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "4.08",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "85.45",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "2.97",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "134.46",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "4.84",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "72.20",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "2.34",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "76.31",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "xrayreceiver_mock",
            "value": "5.22",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "xrayreceiver_mock",
            "value": "155.48",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "5.68",
            "unit": "%",
            "extra": "Trace (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "79.54",
            "unit": "MB",
            "extra": "Trace (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "27.66",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "logzio_exporter_trace_mock",
            "value": "104.16",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "xrayreceiver_mock",
            "value": "25.71",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "xrayreceiver_mock",
            "value": "546.60",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_mock",
            "value": "32.97",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_mock",
            "value": "76.52",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "27.74",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_trace_mock",
            "value": "71.05",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "26.96",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "sapm_exporter_trace_mock",
            "value": "86.26",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "16.58",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "jaeger_mock",
            "value": "151.25",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "28.45",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_trace_mock",
            "value": "714.09",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "29.97",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "zipkin_mock",
            "value": "495.85",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "31.76",
            "unit": "%",
            "extra": "Trace (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_trace_mock",
            "value": "74.76",
            "unit": "MB",
            "extra": "Trace (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_metric_mock",
            "value": "59.25",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "statsd_mock",
            "value": "0.02",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "statsd_mock",
            "value": "58.96",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "60.25",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "0.14",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "72.03",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "58.88",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "60.16",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "62.55",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 100) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "62.01",
            "unit": "MB",
            "extra": "Metric (TPS: 100) - Average Memory Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "61.44",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "7.57",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "259.22",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "58.44",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "60.32",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_metric_mock",
            "value": "59.29",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "61.45",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "statsd_mock",
            "value": "0.01",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "statsd_mock",
            "value": "58.50",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 5000) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "58.92",
            "unit": "MB",
            "extra": "Metric (TPS: 5000) - Average Memory Usage"
          },
          {
            "name": "statsd_mock",
            "value": "0.01",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "statsd_mock",
            "value": "57.97",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "dynatrace_exporter_metric_mock",
            "value": "59.00",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "0.05",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_grpc_exporter_metric_mock",
            "value": "60.88",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "signalfx_exporter_metric_mock",
            "value": "60.32",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "1.48",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "prometheus_mock",
            "value": "105.92",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_http_exporter_metric_mock",
            "value": "60.45",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "otlp_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "otlp_metric_mock",
            "value": "60.38",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "0.04",
            "unit": "%",
            "extra": "Metric (TPS: 1000) - Average CPU Usage"
          },
          {
            "name": "datadog_exporter_metric_mock",
            "value": "61.86",
            "unit": "MB",
            "extra": "Metric (TPS: 1000) - Average Memory Usage"
          }
        ]
      }
    ]
  }
}