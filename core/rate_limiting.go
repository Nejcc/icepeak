package core

import (
	"net/http"
	"sync"
	"time"
)

// RateLimiter stores information about rate limiting for clients.
type RateLimiter struct {
	visitors map[string]*Visitor
	mu       sync.Mutex
	rate     time.Duration
	burst    int
}

// Visitor tracks request count and last seen time.
type Visitor struct {
	limiter  *time.Ticker
	lastSeen time.Time
}

// NewRateLimiter creates a new RateLimiter.
func NewRateLimiter(rate time.Duration, burst int) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*Visitor),
		rate:     rate,
		burst:    burst,
	}
	go rl.cleanupVisitors()
	return rl
}

// AddVisitor adds or updates a visitor in the rate limiter.
func (rl *RateLimiter) AddVisitor(ip string) *time.Ticker {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[ip]
	if !exists {
		limiter := time.NewTicker(rl.rate)
		rl.visitors[ip] = &Visitor{limiter: limiter, lastSeen: time.Now()}
		return limiter
	}
	v.lastSeen = time.Now()
	return v.limiter
}

// cleanupVisitors removes old visitors from the rate limiter.
func (rl *RateLimiter) cleanupVisitors() {
	for {
		time.Sleep(1 * time.Minute)
		rl.mu.Lock()
		for ip, v := range rl.visitors {
			if time.Since(v.lastSeen) > rl.rate*2 {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}

// RateLimitingMiddleware limits the number of requests per client.
func RateLimitingMiddleware(rateLimiter *RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			limiter := rateLimiter.AddVisitor(ip)

			select {
			case <-limiter.C:
				next.ServeHTTP(w, r)
			default:
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			}
		})
	}
}
