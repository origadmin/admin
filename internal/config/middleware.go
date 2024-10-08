package config

type Metrics struct {
	Enabled bool
	Name    string
}

type Traces struct {
	Enabled bool
	Name    string
}

type Logger struct {
	Enabled bool
	Name    string
}

type Middleware struct {
	Metrics Metrics
	Traces  Traces
	Logger  Logger
}
