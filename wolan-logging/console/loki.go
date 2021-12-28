package console

import (
	"context"
	"net/http"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/libbeat/outputs"
	"github.com/elastic/beats/v7/libbeat/publisher"
)

func init() {
	outputs.RegisterType("loki", newLoki)
}

type clientConfig struct {
	Endpoint string `config:"endpoint"`
}

func newLoki(_ outputs.IndexManager, _ beat.Info, observer outputs.Observer, cfg *common.Config) (outputs.Group, error) {
	config := clientConfig{}
	if err := cfg.Unpack(&config); err != nil {
		return outputs.Fail(err)
	}

	clients := make([]outputs.NetworkClient, 1)
	client := &LokiClient{
		observer: observer,
	}
	clients[0] = client

	return outputs.SuccessNet(false, 200, 3, clients)
}

type LokiClient struct {
	observer outputs.Observer
	client   *http.Client
}

func (c *LokiClient) String() string {
	return "loki"
}

func (c *LokiClient) Connect() error {
	c.client = &http.Client{
		Timeout: 2 * time.Second,
	}

	return nil
}

func (c *LokiClient) Close() error {
	c.client = nil
	return nil
}

func (c *LokiClient) Publish(_ context.Context, batch publisher.Batch) error {
	st := c.observer
	events := batch.Events()
	st.NewBatch(len(events))

	for _, event := range events {
		logp.L().Info(event.Content)
	}

	batch.ACK()
	st.Dropped(0)
	st.Acked(len(events))

	return nil
}
