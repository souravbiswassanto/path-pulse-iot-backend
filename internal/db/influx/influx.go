package influx

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

type InfluxDbClientBuilder struct {
	url    string
	org    string
	token  string
	bucket string
}

type InfluxDBOptions struct {
	Url    string `yaml:"url,omitempty"`
	Org    string `yaml:"org,omitempty"`
	Token  string `yaml:"token,omitempty"`
	Bucket string `yaml:"bucket,omitempty"`
}

func NewInfluxDbClientBuilder() *InfluxDbClientBuilder {
	return &InfluxDbClientBuilder{}
}

func (i *InfluxDbClientBuilder) WithOrg(org string) *InfluxDbClientBuilder {
	i.org = org
	return i
}

func (i *InfluxDbClientBuilder) WithBucket(bk string) *InfluxDbClientBuilder {
	i.bucket = bk
	return i
}

func (i *InfluxDbClientBuilder) WithURL(url string) *InfluxDbClientBuilder {
	i.url = url
	return i
}

func (i *InfluxDbClientBuilder) WithToken(token string) *InfluxDbClientBuilder {
	i.token = token
	return i
}

func (i *InfluxDbClientBuilder) InfluxDbClient() influxdb2.Client {
	return influxdb2.NewClient(i.url, i.token)

}

func InfluxDBClient() {
	bucket := "example-bucket"
	org := "example-org"
	token := "MyInitialAdminToken0=="
	// Store the URL of your InfluxDB instance
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, token)
	writeAPI := client.WriteAPIBlocking(org, bucket)
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now())
	writeAPI.WritePoint(context.Background(), p)
	client.Close()
}

// this is just a demo query, created this while practicing
func query() {
	bucket := "example-bucket"
	org := "example-org"
	token := "MyInitialAdminToken0=="
	// Store the URL of your InfluxDB instance
	url := "http://localhost:8086"
	// Create client
	client := influxdb2.NewClient(url, token)
	// Get query client
	queryAPI := client.QueryAPI(org)
	// Get QueryTableResult
	result, err := queryAPI.Query(context.Background(), fmt.Sprintf(`from(bucket:"%s")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`, bucket))
	if err == nil {
		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
			fmt.Printf("value: %v\n", result.Record().Value())
		}
		// Check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
	// Ensures background processes finishes
	client.Close()

}
