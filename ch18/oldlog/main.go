// Oldlog reports the progress of a batch of jobs using the log package.
package main

import (
	"errors"
	"log"
	"time"
)

type job struct {
	id    int
	user  string
	bytes int
	err   error
}

var jobs = []job{
	{id: 1, user: "alice", bytes: 1024},
	{id: 2, user: "bob", err: errors.New("connection reset by peer")},
	{id: 3, user: "carol dean", bytes: 33},
}

func main() {
	log.SetPrefix("batch: ")
	log.SetFlags(log.Ltime)
	for _, j := range jobs {
		start := time.Now()
		if j.err != nil {
			log.Printf("job %d for %s failed: %v", j.id, j.user, j.err)
			continue
		}
		log.Printf("job %d for %s copied %d bytes in %v", j.id, j.user,
			j.bytes, time.Since(start).Round(time.Millisecond))
	}
}
