package requests

import (
	"sync/atomic"

	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/golib/v3/sfxclient"
)

// InternalMetrics returns datapoints that describe the current state of the
// dimension update client
func (rs *ReqSender) InternalMetrics() []*datapoint.Datapoint {
	return []*datapoint.Datapoint{
		sfxclient.CumulativeP("sfxagent.dim_updates_started", map[string]string{"client": rs.clientName}, &rs.TotalRequestsStarted),
		sfxclient.CumulativeP("sfxagent.dim_updates_completed", map[string]string{"client": rs.clientName}, &rs.TotalRequestsCompleted),
		sfxclient.CumulativeP("sfxagent.dim_updates_failed", map[string]string{"client": rs.clientName}, &rs.TotalRequestsFailed),
		sfxclient.Gauge("sfxagent.dim_request_senders", map[string]string{"client": rs.clientName}, atomic.LoadInt64(&rs.RunningWorkers)),
	}
}
