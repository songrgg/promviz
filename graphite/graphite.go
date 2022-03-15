package graphite

import (
	"context"
	client "github.com/lorehov/graphite-api-client"
	prommodel "github.com/prometheus/common/model"
	"strings"
	"time"
)

type Client struct {
	c *client.Client
}

func NewClient(addr string) (Client, error) {
	c, err := client.NewFromString(addr)
	return Client{
		c: c,
	}, err
}

func (c *Client) Query(ctx context.Context, query string, ts time.Time) (prommodel.Value, error) {
	// fetch the 1 minute duration metrics and use the last value.
	// By convention, the metric name must be <"source" "target" "status_code"> then it will generate
	// prometheus labels according to the values.
	series, err := c.c.QueryRender(client.RenderRequest{
		From:  ts.Add(-1 * time.Minute),
		Until: ts,
		Targets: []string{
			query,
		},
	})
	if err != nil {
		return nil, err
	}

	vec := prommodel.Vector{}
	for _, s := range series {
		parts := strings.Split(s.Target, " ")
		if len(parts) != 3 {
			continue
		}

		if len(s.Datapoints) < 1 {
			continue
		}
		final := len(s.Datapoints) - 1
		vec = append(vec, &prommodel.Sample{
			Metric: prommodel.Metric{
				"source": prommodel.LabelValue(parts[0]),
				"target": prommodel.LabelValue(parts[1]),
				"status": prommodel.LabelValue(parts[2]),
			},
			Value:     prommodel.SampleValue(s.Datapoints[final].Value),
			Timestamp: prommodel.Time(s.Datapoints[final].Timestamp.UnixMilli()),
		})
	}

	return vec, nil
}
