package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var _ BreakerConfig = (*breakerConfig)(nil)

const (
	brRequests = "BREAKER_REQUESTS"
	brInterval = "BREAKER_INTERVAL"
	brTimeout  = "BREAKER_TIMEOUT"
)

type BreakerConfig interface {
	Requests() int
	Interval() time.Duration
	Timeout() time.Duration
}

type breakerConfig struct {
	requests int
	interval time.Duration
	timeout  time.Duration
}

func NewBreakerConfig() (*breakerConfig, error) {
	requestsStr := os.Getenv(brRequests)
	if requestsStr == "" {
		return nil, fmt.Errorf("invalid requests number: empty value")
	}

	requests, err := strconv.Atoi(requestsStr)
	if err != nil {
		return nil, fmt.Errorf("invalid requests number: %s", err)
	}

	intervalStr := os.Getenv(brInterval)
	if intervalStr == "" {
		return nil, fmt.Errorf("invalid interval: empty value")
	}

	interval, err := time.ParseDuration(intervalStr)
	if err != nil {
		return nil, fmt.Errorf("invalid interval: %s", err)
	}

	timeoutStr := os.Getenv(brTimeout)
	if timeoutStr == "" {
		return nil, fmt.Errorf("invalid timeout: empty value")
	}

	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		return nil, fmt.Errorf("invalid timeout: %s", err)
	}

	fmt.Printf("BreakerConfig: requests=%d, interval=%v, timeout=%v\n", requests, interval, timeout)

	return &breakerConfig{requests: requests, interval: interval, timeout: timeout}, nil
}

func (c *breakerConfig) Requests() int {
	return c.requests
}

func (c *breakerConfig) Interval() time.Duration {
	return c.interval
}

func (c *breakerConfig) Timeout() time.Duration {
	return c.timeout
}
