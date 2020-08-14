package txlbs

import (
	"time"

	"github.com/xiaojiaoyu100/cast"
)

// Config represents a client configuration.
type Config struct {
	// Key for auth
	Key       string
	SecretKey string
}

// Client encapsulates a set of apis.
type Client struct {
	Config
	ca *cast.Cast
}

// New generates a client.
func New(setters ...Setter) (*Client, error) {
	var err error
	c := &Client{}
	c.ca, err = cast.New(
		cast.WithBaseURL(base),
		cast.WithRetry(2),
		cast.WithLinearBackoffStrategy(10*time.Millisecond),
	)
	if err != nil {
		return nil, err
	}
	for _, setter := range setters {
		if err := setter(&c.Config); err != nil {
			return nil, err
		}
	}
	return c, nil
}
