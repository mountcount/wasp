package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/iotaledger/wasp/packages/isc"
)

type IChainStateManagerMetrics interface {
	SetCacheSize(int)
	IncBlocksFetching()
	DecBlocksFetching()
	IncBlocksPending()
	DecBlocksPending()
	BlockIndexCommitted(uint32)
	IncRequestsWaiting()
	SubRequestsWaiting(int)
	SetRequestsWaiting(int)
	ConsensusStateProposalHandled(time.Duration)
	ConsensusDecidedStateHandled(time.Duration)
	ConsensusBlockProducedHandled(time.Duration)
	ChainFetchStateDiffHandled(time.Duration)
	StateManagerTimerTickHandled(time.Duration)
	StateManagerBlockFetched(time.Duration)
	StatePruned(time.Duration, uint32)
	PruningCompleted(time.Duration, int)
}

var (
	_ IChainStateManagerMetrics = &emptyChainStateManagerMetric{}
	_ IChainStateManagerMetrics = &chainStateManagerMetric{}
)

type emptyChainStateManagerMetric struct{}

func NewEmptyChainStateManagerMetric() IChainStateManagerMetrics {
	return &emptyChainStateManagerMetric{}
}
func (m *emptyChainStateManagerMetric) SetCacheSize(int)                            {}
func (m *emptyChainStateManagerMetric) IncBlocksFetching()                          {}
func (m *emptyChainStateManagerMetric) DecBlocksFetching()                          {}
func (m *emptyChainStateManagerMetric) IncBlocksPending()                           {}
func (m *emptyChainStateManagerMetric) DecBlocksPending()                           {}
func (m *emptyChainStateManagerMetric) BlockIndexCommitted(uint32)                  {}
func (m *emptyChainStateManagerMetric) IncRequestsWaiting()                         {}
func (m *emptyChainStateManagerMetric) SubRequestsWaiting(int)                      {}
func (m *emptyChainStateManagerMetric) SetRequestsWaiting(int)                      {}
func (m *emptyChainStateManagerMetric) ConsensusStateProposalHandled(time.Duration) {}
func (m *emptyChainStateManagerMetric) ConsensusDecidedStateHandled(time.Duration)  {}
func (m *emptyChainStateManagerMetric) ConsensusBlockProducedHandled(time.Duration) {}
func (m *emptyChainStateManagerMetric) ChainFetchStateDiffHandled(time.Duration)    {}
func (m *emptyChainStateManagerMetric) StateManagerTimerTickHandled(time.Duration)  {}
func (m *emptyChainStateManagerMetric) StateManagerBlockFetched(time.Duration)      {}
func (m *emptyChainStateManagerMetric) StatePruned(time.Duration, uint32)           {}
func (m *emptyChainStateManagerMetric) PruningCompleted(time.Duration, int)         {}

type chainStateManagerMetric struct {
	provider      *ChainMetricsProvider
	metricsLabels prometheus.Labels
}

func newChainStateManagerMetric(provider *ChainMetricsProvider, chainID isc.ChainID) *chainStateManagerMetric {
	metricsLabels := getChainLabels(chainID)

	// init values so they appear in prometheus
	provider.smCacheSize.With(metricsLabels)
	provider.smBlocksFetching.With(metricsLabels)
	provider.smBlocksPending.With(metricsLabels)
	provider.smBlocksCommitted.with(metricsLabels)
	provider.smRequestsWaiting.With(metricsLabels)
	provider.smCSPHandlingDuration.With(metricsLabels)
	provider.smCDSHandlingDuration.With(metricsLabels)
	provider.smCBPHandlingDuration.With(metricsLabels)
	provider.smFSDHandlingDuration.With(metricsLabels)
	provider.smTTHandlingDuration.With(metricsLabels)
	provider.smBlockFetchDuration.With(metricsLabels)
	provider.smPruningRunDuration.With(metricsLabels)
	provider.smPruningSingleStateDuration.With(metricsLabels)
	provider.smPruningStatesInRun.With(metricsLabels)
	provider.smStatesPruned.with(metricsLabels)

	return &chainStateManagerMetric{
		provider:      provider,
		metricsLabels: metricsLabels,
	}
}

func (m *chainStateManagerMetric) SetCacheSize(size int) {
	m.provider.smCacheSize.With(m.metricsLabels).Set(float64(size))
}

func (m *chainStateManagerMetric) IncBlocksFetching() {
	m.provider.smBlocksFetching.With(m.metricsLabels).Inc()
}

func (m *chainStateManagerMetric) DecBlocksFetching() {
	m.provider.smBlocksFetching.With(m.metricsLabels).Dec()
}

func (m *chainStateManagerMetric) IncBlocksPending() {
	m.provider.smBlocksPending.With(m.metricsLabels).Inc()
}

func (m *chainStateManagerMetric) DecBlocksPending() {
	m.provider.smBlocksPending.With(m.metricsLabels).Dec()
}

func (m *chainStateManagerMetric) BlockIndexCommitted(blockIndex uint32) {
	m.provider.smBlocksCommitted.countValue(m.metricsLabels, float64(blockIndex))
}

func (m *chainStateManagerMetric) IncRequestsWaiting() {
	m.provider.smRequestsWaiting.With(m.metricsLabels).Inc()
}

func (m *chainStateManagerMetric) SubRequestsWaiting(count int) {
	m.provider.smRequestsWaiting.With(m.metricsLabels).Sub(float64(count))
}

func (m *chainStateManagerMetric) SetRequestsWaiting(count int) {
	m.provider.smRequestsWaiting.With(m.metricsLabels).Set(float64(count))
}

func (m *chainStateManagerMetric) ConsensusStateProposalHandled(duration time.Duration) {
	m.provider.smCSPHandlingDuration.With(m.metricsLabels).Observe(duration.Seconds())
}

func (m *chainStateManagerMetric) ConsensusDecidedStateHandled(duration time.Duration) {
	m.provider.smCDSHandlingDuration.With(m.metricsLabels).Observe(duration.Seconds())
}

func (m *chainStateManagerMetric) ConsensusBlockProducedHandled(duration time.Duration) {
	m.provider.smCBPHandlingDuration.With(m.metricsLabels).Observe(duration.Seconds())
}

func (m *chainStateManagerMetric) ChainFetchStateDiffHandled(duration time.Duration) {
	m.provider.smFSDHandlingDuration.With(m.metricsLabels).Observe(duration.Seconds())
}

func (m *chainStateManagerMetric) StateManagerTimerTickHandled(duration time.Duration) {
	m.provider.smTTHandlingDuration.With(m.metricsLabels).Observe(duration.Seconds())
}

func (m *chainStateManagerMetric) StateManagerBlockFetched(duration time.Duration) {
	m.provider.smBlockFetchDuration.With(m.metricsLabels).Observe(duration.Seconds())
}

func (m *chainStateManagerMetric) StatePruned(duration time.Duration, stateIndex uint32) {
	m.provider.smPruningSingleStateDuration.With(m.metricsLabels).Observe(duration.Seconds())
	m.provider.smStatesPruned.countValue(m.metricsLabels, float64(stateIndex))
}

func (m *chainStateManagerMetric) PruningCompleted(duration time.Duration, statesPruned int) {
	m.provider.smPruningRunDuration.With(m.metricsLabels).Observe(duration.Seconds())
	m.provider.smPruningStatesInRun.With(m.metricsLabels).Observe(float64(statesPruned))
}
