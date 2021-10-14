package kmeans

// Config exposes configuration options to the caller.
type Config struct {
	TrainRounds int
	Mthd        InitMethod
}

// NewConfig returns the default configuration updated with any
// options.
func NewConfig(opts ...Option) Config {
	cfg := Config{
		TrainRounds: 1,
		Mthd:        Random,
	}

	cfg.update(opts...)
	return cfg
}

// update a configuration.
func (cfg *Config) update(opts ...Option) {
	for i := 0; i < len(opts); i++ {
		opts[i](cfg)
	}
}
