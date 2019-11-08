package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func (e *examples) emitCounters() {
	rand.Seed(time.Now().UnixNano())
	endpoints := []string{"one", "two", "three", "four", "five"}
	zones := []string{"abc1", "abc2", "xyz1", "xyz2"}
	var hosts [85]map[string]string
	for i := 0; i < len(hosts); i++ {
		random := rand.Int()
		name := fmt.Sprintf("host%02d", i+1)
		zone := zones[random%len(zones)]
		hosts[i] = map[string]string{
			"name": name,
			"zone": zone,
		}
	}
	// Generating fake requests from all hosts.
	log.Printf("Faking requests for %d hosts and %d endpoints", len(hosts), len(endpoints))
	for _, host := range hosts {
		for i, endpoint := range endpoints {
			waitModulo := i*5 + 1
			go func(h map[string]string, ep string, m int) {
				for {
					random := rand.Int()
					labels := []string{"/endpoints/" + ep, h["name"], h["zone"]}
					e.metrics.requestsTotal.WithLabelValues(labels...).Inc()
					ms := time.Duration(random%m + 1)
					time.Sleep(ms * time.Millisecond)
				}
			}(host, endpoint, waitModulo)
		}
	}
	log.Printf("%d goroutines launched", len(hosts)*len(endpoints))
}
