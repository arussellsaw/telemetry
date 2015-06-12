package reporters

import (
	"fmt"
	"net/url"
	"time"

	"github.com/arussellsaw/telemetry"
	"github.com/influxdb/influxdb/client"
)

//InfluxReporter influxdb reporter
type InfluxReporter struct {
	Host       string
	Interval   time.Duration
	Tel        *telemetry.Telemetry
	Database   string
	Connection *client.Client
}

//Report - report metrics to influxdb
func (r *InfluxReporter) Report() error {
	url, err := url.Parse(r.Host)
	if err != nil {
		return err
	}
	clientConf := client.Config{
		URL: *url,
	}
	conn, err := client.NewClient(clientConf)
	if err != nil {
		return err
	}
	r.Connection = conn

	go r.submitMetrics()
	return nil
}

func (r *InfluxReporter) submitMetrics() {
	for {
		i := 0
		points := make([]client.Point, len(r.Tel.GetAll()))
		for name, metric := range r.Tel.GetAll() {
			point := client.Point{
				Measurement: name,
				Time:        time.Now(),
				Fields: map[string]interface{}{
					"value": metric,
				},
				Precision: "s",
			}
			points[i] = point
			i++
		}

		batch := client.BatchPoints{
			Points:   points,
			Database: r.Database,
		}
		_, err := r.Connection.Write(batch)
		if err != nil {
			fmt.Println("failed write: ", err)
		}
		time.Sleep(r.Interval)
	}
}
