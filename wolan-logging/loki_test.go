package main

import (
	"testing"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/publisher"
	"github.com/zx5435/wolan/wolan-logging/loki"
)

func TestConsoleOutput(t *testing.T) {
	client := &loki.LokiClient{}

	var events []publisher.Event
	events = append(events, publisher.Event{
		Content: beat.Event{
			Timestamp:  time.Now(),
			Meta:       nil,
			Fields:     nil,
			Private:    nil,
			TimeSeries: false,
		},
		Flags: 0,
		Cache: publisher.EventCache{},
	})
	events = append(events, publisher.Event{
		Content: beat.Event{
			Timestamp:  time.Now(),
			Meta:       nil,
			Fields:     nil,
			Private:    nil,
			TimeSeries: false,
		},
		Flags: 1,
		Cache: publisher.EventCache{},
	})
	client.PublicBatch(events)
}
