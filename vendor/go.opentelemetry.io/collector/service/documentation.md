[comment]: <> (Code generated by mdatagen. DO NOT EDIT.)

# service

## Internal Telemetry

The following telemetry is emitted by this component.

### otelcol_process_cpu_seconds

Total CPU user and system time in seconds

| Unit | Metric Type | Value Type | Monotonic |
| ---- | ----------- | ---------- | --------- |
| s | Sum | Double | true |

### otelcol_process_memory_rss

Total physical memory (resident set size)

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### otelcol_process_runtime_heap_alloc_bytes

Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### otelcol_process_runtime_total_alloc_bytes

Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')

| Unit | Metric Type | Value Type | Monotonic |
| ---- | ----------- | ---------- | --------- |
| By | Sum | Int | true |

### otelcol_process_runtime_total_sys_memory_bytes

Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### otelcol_process_uptime

Uptime of the process

| Unit | Metric Type | Value Type | Monotonic |
| ---- | ----------- | ---------- | --------- |
| s | Sum | Double | true |