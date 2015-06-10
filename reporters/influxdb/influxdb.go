package influxdbReporter

import (
	"log"
	"net/url"
	"time"

	"github.com/arussellsaw/telemetry"
	"github.com/influxdb/influxdb/client"
)

//Config - configuration for influxdb reporter
type Config struct {
	Host     string
	Interval time.Duration
	Tel      *telemetry.Telemetry
	Database string
}

//Report - report metrics to influxdb
func Report(conf Config) error {
	url, err := url.Parse(conf.Host)
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

	go submitMetrics(conn, conf)
	return nil
}

func submitMetrics(conn *client.Client, conf Config) {
	for {
		i := 0
		points := make([]client.Point, len(conf.Tel.GetAll()))
		for name, metric := range conf.Tel.GetAll() {
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
			Points:          points,
			Database:        conf.Database,
			RetentionPolicy: "default",
		}
		_, err := conn.Write(batch)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(conf.Interval)
	}
}
