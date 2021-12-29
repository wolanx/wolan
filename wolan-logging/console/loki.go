package console

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/libbeat/outputs"
	"github.com/elastic/beats/v7/libbeat/outputs/codec"
	jsonc "github.com/elastic/beats/v7/libbeat/outputs/codec/json"
	"github.com/elastic/beats/v7/libbeat/publisher"
)

var (
	logger = logp.NewLogger("output.loki")
)

func init() {
	outputs.RegisterType("loki", newLoki)
}

type LokiConfig struct {
	Endpoint string `config:"endpoint"`
}

func newLoki(_ outputs.IndexManager, beat beat.Info, observer outputs.Observer, cfg *common.Config) (outputs.Group, error) {
	config := LokiConfig{}
	if err := cfg.Unpack(&config); err != nil {
		return outputs.Fail(err)
	}

	clients := make([]outputs.NetworkClient, 1)
	client := &LokiClient{
		observer: observer,
		codec: jsonc.New(beat.Version, jsonc.Config{
			Pretty:     true,
			EscapeHTML: false,
		}),
	}
	clients[0] = client

	return outputs.SuccessNet(false, 200, 3, clients)
}

type LokiClient struct {
	client   *http.Client
	observer outputs.Observer
	codec    codec.Codec
	index    string
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

	c.PublicBatch(events)

	batch.ACK()
	st.Dropped(0)
	st.Acked(len(events))

	return nil
}

func (c *LokiClient) PublicBatch(events []publisher.Event) {
	b := make(map[string]interface{})
	var v [][]interface{}

	for _, event := range events {
		vv := &event.Content
		flatten := vv.Fields.Flatten()
		flatten.Delete("message")
		fmt.Printf("----fields = %+v\n", flatten)
		msg, _ := vv.Fields.GetValue("message")
		v = append(v, []interface{}{strconv.FormatInt(vv.Timestamp.UnixNano(), 10), msg})
	}

	sss := make(map[string]interface{})
	sss["stream"] = nil
	sss["values"] = v
	b["streams"] = []map[string]interface{}{sss}
	indent, _ := json.MarshalIndent(b, "", "  ")
	fmt.Println("----http =", string(indent))
}

type event struct {
	Timestamp time.Time     `json:"@timestamp"`
	Fields    common.MapStr `json:"-"`
}

func makeEvent(v *beat.Event) map[string]json.RawMessage {
	e := event{Timestamp: v.Timestamp.UTC(), Fields: v.Fields}
	b, err := json.Marshal(e)
	if err != nil {
		logger.Warn("Error encoding event to JSON: %v", err)
	}

	var eventMap map[string]json.RawMessage
	err = json.Unmarshal(b, &eventMap)
	if err != nil {
		logger.Warn("Error decoding JSON to map: %v", err)
	}
	// Add the individual fields to the map, flatten "Fields"
	for j, k := range e.Fields {
		b, err = json.Marshal(k)
		if err != nil {
			logger.Warn("Error encoding map to JSON: %v", err)
		}
		eventMap[j] = b
	}
	indent, _ := json.MarshalIndent(eventMap, "", "  ")
	fmt.Println(string(indent))

	return eventMap
}
