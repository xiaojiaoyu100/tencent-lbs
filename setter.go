package txlbs

// Setter changes config.
type Setter func(config *Config) error

// WithKey changes a config key.
func WithKey(key string) Setter {
	return func(config *Config) error {
		config.Key = key
		return nil
	}
}
