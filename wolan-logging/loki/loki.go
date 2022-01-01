package loki

import (
	"context"
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
	"github.com/grafana/loki/pkg/logproto"
	"github.com/prometheus/common/model"
	"google.golang.org/grpc"
)

var (
	logger = logp.NewLogger("output.loki")
)

func init() {
	outputs.RegisterType("loki", newLoki)
}

type lokiConfig struct {
}

func newLoki(_ outputs.IndexManager, beat beat.Info, observer outputs.Observer, cfg *common.Config) (outputs.Group, error) {
	config := lokiConfig{}
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
	grpc     logproto.PusherClient
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

	conn, err := grpc.Dial(c.host, grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.grpc = logproto.NewPusherClient(conn)

	return nil
}

func (c *LokiClient) Close() error {
	c.client = nil
	c.grpc = nil
	return nil
}

func (c *LokiClient) Publish(ctx context.Context, batch publisher.Batch) error {
	st := c.observer
	events := batch.Events()
	st.NewBatch(len(events))

	//c.PublicBatchHttp(events)

	// grpc
	redo, err := c.PublicBatchGrpc(ctx, events)
	dropped := len(redo)
	if dropped == 0 {
		batch.ACK()
	} else {
		batch.RetryEvents(redo)
	}

	st.Dropped(dropped)
	st.Acked(len(events) - dropped)

	return err
}

func f2l(fields common.MapStr) model.LabelSet {
	fields.Delete("agent")
	fields.Delete("message")
	//ret := common.MapStr{}
	ret := model.LabelSet{}

	for k, v := range fields.Flatten() {
		label := strings.ReplaceAll(strings.ReplaceAll(k, "-", "_"), ".", "_")
		ret[model.LabelName(label)] = model.LabelValue(fmt.Sprintf("%s", v))
	}
	return ret
}

func (c *LokiClient) PublicBatchGrpc(ctx context.Context, events []publisher.Event) ([]publisher.Event, error) {
	for _, event := range events {
		event := &event.Content
		fields := event.Fields
		msg, _ := fields.GetValue("message")

		logs := []logproto.Entry{
			{Timestamp: event.Timestamp, Line: msg.(string)},
		}
		req := &logproto.PushRequest{Streams: []logproto.Stream{
			{Labels: f2l(fields).String(), Entries: logs},
		}}
		_, err := c.grpc.Push(ctx, req)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil, nil
}

func (c *LokiClient) PublicBatchHttp(events []publisher.Event) {
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
