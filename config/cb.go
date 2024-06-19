package config

import (
 "github.com/sony/gobreaker"
 "time"
)

func CBConfig() *gobreaker.CircuitBreaker {
 settings := gobreaker.Settings{
  Name:    "server-circuit-breaker",
  Timeout: 5 * time.Second,
  ReadyToTrip: func(counts gobreaker.Counts) bool {
   return counts.TotalFailures >= 3
  },
 }
 return gobreaker.NewCircuitBreaker(settings)
}
