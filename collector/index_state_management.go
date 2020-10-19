package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

type ismMetric struct {
	Type   prometheus.ValueType
	Desc   *prometheus.Desc
	Value  func(snapshotStats SnapshotStatDataResponse) float64
	Labels func(indexName string, snapshotStats SnapshotStatDataResponse) []string
}
