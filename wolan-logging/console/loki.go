package console

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/libbeat/outputs"
	"github.com/elastic/beats/v7/libbeat/publisher"
)

var (
	logger = logp.NewLogger("output.loki")
)

func init() {
	outputs.RegisterType("loki", newLoki)
}

type LokiConfig struct {
}

func newLoki(_ outputs.IndexManager, beat beat.Info, observer outputs.Observer, cfg *common.Config) (outputs.Group, error) {
	config := LokiConfig{}
	if err := cfg.Unpack(&config); err != nil {
		return outputs.Fail(err)
	}

	hosts, err := outputs.ReadHostList(cfg)
	if err != nil {
		return outputs.Fail(err)
	}

	clients := make([]outputs.NetworkClient, len(hosts))
	for i, host := range hosts {
		clients[i] = &LokiClient{
			host:     host,
			observer: observer,
		}
	}

	return outputs.SuccessNet(false, 200, 3, clients)
}

type LokiClient struct {
	host     string
	client   *http.Client
	observer outputs.Observer
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
	var values [][]interface{}

	labels := common.MapStr{}

	for i := range events {
		event := &events[i].Content
		nano := event.Timestamp.UnixNano()
		fields := event.Fields
		msg, _ := fields.GetValue("message")
		values = append(values, []interface{}{strconv.FormatInt(nano, 10), msg})

		if i == 0 {
			fields.Delete("agent")
			fields.Delete("message")
			for k, v := range fields.Flatten() {
				labels[strings.ReplaceAll(strings.ReplaceAll(k, "-", "_"), ".", "_")] = v
			}
		}
	}

	s1 := common.MapStr{}
	s11 := common.MapStr{}
	s11["stream"] = labels
	s11["values"] = values
	s1["streams"] = []common.MapStr{s11}
	fmt.Println("----http =", s1.StringToPrint())

	req, _ := http.NewRequest("POST", "http://"+c.host+"/loki/api/v1/push", strings.NewReader(s1.StringToPrint()))
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("resp", string(content))
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
