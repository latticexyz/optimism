package metrics

import (
	"github.com/ethereum-optimism/optimism/op-service/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type AltDAMetricer interface {
	RecordActiveChallenge(commBlock uint64, startBlock uint64, hash []byte)
	RecordResolvedChallenge(hash []byte)
	RecordExpiredChallenge(hash []byte)
	RecordStorageError()
}

type AltDAMetrics struct {
	ChallengesStatus *prometheus.GaugeVec

	StorageErrors *metrics.Event
}

var _ AltDAMetricer = (*AltDAMetrics)(nil)

func MakeAltDAMetrics(ns string, factory metrics.Factory) *AltDAMetrics {
	return &AltDAMetrics{
		ChallengesStatus: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: ns,
			Name:      "challenges_status",
			Help:      "Gauge representing the status of challenges synced",
		}, []string{"status"}),
		StorageErrors: metrics.NewEvent(factory, ns, "", "storage_errors", "errors when fetching or uploading to storage service"),
	}
}

func (m *AltDAMetrics) RecordChallenge(status string) {
	m.ChallengesStatus.WithLabelValues(status).Inc()
}

// RecordActiveChallenge records when a commitment is challenged including the block where the commitment
// is included, the block where the commitment was challenged and the commitment hash.
func (m *AltDAMetrics) RecordActiveChallenge(commBlock uint64, startBlock uint64, hash []byte) {
	m.RecordChallenge("active")
}

func (m *AltDAMetrics) RecordResolvedChallenge(hash []byte) {
	m.RecordChallenge("resolved")
}

func (m *AltDAMetrics) RecordExpiredChallenge(hash []byte) {
	m.RecordChallenge("expired")
}

func (m *AltDAMetrics) RecordStorageError() {
	m.StorageErrors.Record()
}

type NoopAltDAMetrics struct{}

func (m *NoopAltDAMetrics) RecordActiveChallenge(commBlock uint64, startBlock uint64, hash []byte) {}
func (m *NoopAltDAMetrics) RecordResolvedChallenge(hash []byte)                                    {}
func (m *NoopAltDAMetrics) RecordExpiredChallenge(hash []byte)                                     {}
func (m *NoopAltDAMetrics) RecordStorageError()                                                    {}
