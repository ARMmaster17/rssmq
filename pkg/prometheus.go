package pkg

import "github.com/prometheus/client_golang/prometheus"

var TotalChecks = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "rss_checks_total",
		Help: "Number of RSS feed checks",
	},
)

var NewItems = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "rss_new_items_total",
		Help: "Number of new RSS items found",
	},
	[]string{"url"},
)

func (a *App) registerPrometheusEndpoints() error {
	err := prometheus.Register(TotalChecks)
	if err != nil {
		return err
	}
	err = prometheus.Register(NewItems)
	if err != nil {
		return err
	}
	return nil
}
