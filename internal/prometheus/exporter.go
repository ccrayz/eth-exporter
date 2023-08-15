package prometheus

import (
	"ccrayz/eth-exporter/config"
	"ccrayz/eth-exporter/internal/ethereum"
	"math/big"
	"sync"

	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type Exporter struct {
	metrics          *ethereum.MetricsCollector
	accountAddresses []config.EthAccount
}

func NewExporter(metrics *ethereum.MetricsCollector, accountAddresses []config.EthAccount) *Exporter {
	return &Exporter{
		metrics:          metrics,
		accountAddresses: accountAddresses,
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- prometheus.NewDesc("eth_test",
		"Balance of the Ethereum account",
		nil, nil,
	)
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.collectMetricsPeriodically(ch) // Start the periodic metrics collection in a goroutine
}

func (e *Exporter) collectMetricsPeriodically(ch chan<- prometheus.Metric) {
	var wg sync.WaitGroup
	wg.Add(len(e.accountAddresses)) // 고루틴의 개수만큼 대기 그룹을 설정

	for _, ac := range e.accountAddresses {
		go func(address string, purpose string) {
			defer wg.Done() // 고루틴 종료 시 대기 그룹 해제
			balance, err := e.metrics.GetAccountBalance(address)
			if err != nil {
				// Handle error
				fmt.Println(err)
				balance = new(big.Float)
			}
			balanceFloat64, _ := balance.Float64()
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					"eth_account_balance",
					"Balance of the Ethereum account",
					[]string{"purpose"},
					nil,
				),
				prometheus.GaugeValue,
				balanceFloat64,
				purpose,
			)
		}(ac.Address, ac.Purpose)
	}

	wg.Wait() // 모든 고루틴이 종료될 때까지 대기
}

func RegisterExporter(exporter *Exporter) {
	prometheus.MustRegister(exporter)
}
