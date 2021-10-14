package kmeans

// Option updates a configuration.
type Option func(*Config)

// SetTrainRounds sets the number of training rounds.
func SetTrainRounds(trainRounds int) Option {
	return func(cfg *Config) { cfg.TrainRounds = trainRounds }
}

// SetInitMethod sets the initialization method.
func SetInitMethod(mthd InitMethod) Option {
	return func(cfg *Config) { cfg.Mthd = mthd }
}
