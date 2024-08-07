package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var _ RateLimitConfig = (*rateLimitConfig)(nil)

const (
	rtPeriod = "RATE_LIMIT_PERIOD"
	rtLimit  = "RATE_LIMIT_LIMIT"
)

type RateLimitConfig interface {
	Period() time.Duration
	Limit() int
}

type rateLimitConfig struct {
	period time.Duration
	limit  int
}

func NewRateLimitConfig() (*rateLimitConfig, error) {
	periodStr := os.Getenv(rtPeriod)
	if periodStr == "" {
		return nil, fmt.Errorf("invalid period: empty value")
	}

	period, err := time.ParseDuration(periodStr)
	if err != nil {
		return nil, fmt.Errorf("invalid period: %s", err)
	}

	limitStr := os.Getenv(rtLimit)
	if limitStr == "" {
		return nil, fmt.Errorf("invalid limit: empty value")
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, fmt.Errorf("invalid limit: %s", err)
	}

	fmt.Printf("RateLimitConfig: period=%v, limit=%d\n", period, limit)

	return &rateLimitConfig{period: period, limit: limit}, nil
}

func (c *rateLimitConfig) Period() time.Duration {
	return c.period
}

func (c *rateLimitConfig) Limit() int {
	return c.limit
}
