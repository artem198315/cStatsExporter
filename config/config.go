package config

// Configurations exported
type Configurations struct {
	ScrapeIntervalSeconds int64
	Port                  int
	MetricsPath           string
	Kubernetes            bool			 
}
