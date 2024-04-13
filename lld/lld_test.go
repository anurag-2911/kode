package lld

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	fmt.Print()
}

// Rate Limiter Implementation using token bucket Algorithm

type Bucket struct {
	MaxTokenCount     int
	CurrentTokenCount int
	Lock              sync.Mutex
	Ticks             *time.Ticker
}

func NewBucket(max int, refillInterval time.Duration) *Bucket {
	bucket := &Bucket{
		MaxTokenCount:     max,
		CurrentTokenCount: max/2,
		Ticks:             time.NewTicker(refillInterval),
	}
	go bucket.refillToken()
	return bucket
}
func (bucket *Bucket) refillToken() {
	for range bucket.Ticks.C {
		fmt.Println("adding a token")
		bucket.addToken(1)
	}
}
func (bucket *Bucket) addToken(tokens int) {
	bucket.Lock.Lock()
	defer bucket.Lock.Unlock()
	if (tokens + bucket.CurrentTokenCount) < bucket.MaxTokenCount {
		bucket.CurrentTokenCount += tokens
		fmt.Println("token added")
	} else {
		bucket.CurrentTokenCount = bucket.MaxTokenCount
		fmt.Println("token reset")
	}
}
func (bucket *Bucket) removeTokens(tokens int) bool {
	bucket.Lock.Lock()
	defer bucket.Lock.Unlock()
	if bucket.CurrentTokenCount >= tokens {
		bucket.CurrentTokenCount -= tokens
		fmt.Println("token removed")
		return true
	}
	return false
}
func (bucket *Bucket) getcurrentTokenCount() int {
	bucket.Lock.Lock()
	defer bucket.Lock.Unlock()
	return bucket.CurrentTokenCount
}
func TestRateLimiter(t *testing.T) {
	bucket := NewBucket(10, 1*time.Second)
	simulateRequest(bucket)
}

func simulateRequest(bucket *Bucket) {
	for i := 0; i < 20; i++ {
		fmt.Println("request counter ", i)
		if bucket.removeTokens(1) {
			fmt.Println("token consumed,remaining token ", bucket.getcurrentTokenCount())
		} else {
			fmt.Println("rate limit exceeded")
		}
		time.Sleep(time.Second)
	}
}
