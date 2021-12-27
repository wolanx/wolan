package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/elastic/beats/v7/filebeat/cmd"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/outputs"
	"github.com/elastic/beats/v7/libbeat/publisher"
)

//var Bundle = plugin.Bundle(
//	outputs.Plugin("loki", newHTTPOutput),
//)

func init() {
	outputs.RegisterType("loki", newHTTPOutput)
}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

type clientConfig struct {
	// The endpoint our client should be POSTing to
	Endpoint string `config:"endpoint"`
}

func newHTTPOutput(_ outputs.IndexManager, _ beat.Info, stats outputs.Observer, cfg *common.Config) (outputs.Group, error) {
	config := clientConfig{}
	if err := cfg.Unpack(&config); err != nil {
		return outputs.Fail(err)
	}

	clients := make([]outputs.NetworkClient, 0)
	clients[0] = &LokiClient{
		stats: stats,
	}

	return outputs.SuccessNet(false, 200, 3, clients)
}

type LokiClient struct {
	stats  outputs.Observer
	client *http.Client
}

func (client *LokiClient) String() string {
	return "loki"
}

func (client *LokiClient) Connect() error {
	client.client = &http.Client{
		Timeout: 2 * time.Second,
	}

	return nil
}

func (client *LokiClient) Close() error {
	client.client = nil
	return nil
}

func (client *LokiClient) Publish(_ context.Context, batch publisher.Batch) error {
	events := batch.Events()

	for i, event := range events {
		fmt.Println(i, event)
	}

	client.stats.Acked(len(events))
	batch.ACK()

	return nil
}
