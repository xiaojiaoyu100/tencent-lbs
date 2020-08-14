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

// WithSecretKey changes a config SecretKey.
func WithSecretKey(secretKey string) Setter {
	return func(config *Config) error {
		config.SecretKey = secretKey
		return nil
	}
}
