package service

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db/influx"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	PositionMeasurement      string = "position"
	BloodPressureMeasurement string = "bp"
)

type TrackerService struct {
	cb *influx.InfluxDbClientBuilder
	influx.InfluxDBOptions
}

func NewTrackerService(options *influx.InfluxDBOptions) TrackerService {
	if options == nil {
		options = &influx.InfluxDBOptions{}
	}
	return TrackerService{
		cb: influx.NewInfluxDbClientBuilder().WithOrg(options.Org).WithToken(options.Token).WithURL(options.Url).WithBucket(options.Bucket),
	}
}

func (ts *TrackerService) GetLocation(_ context.Context, userID *models.UserID) (*models.Position, error) {
	c := ts.cb.InfluxDbClient()
	defer c.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	position := &models.Position{}
	queryApi := c.QueryAPI(ts.Org)
	err := ts.GetLastQuery(ctx, userID, queryApi, "Longitude", position)
	if err != nil {
		return nil, err
	}
	err = ts.GetLastQuery(ctx, userID, queryApi, "Latitude", position)
	if err != nil {
		return nil, err
	}
	position.UID = *userID
	return position, err
}

func (ts *TrackerService) GetLastQuery(ctx context.Context, userID *models.UserID, queryApi api.QueryAPI, fieldType string, position *models.Position) error {
	rs, err := queryApi.Query(ctx, fmt.Sprintf(`
	from(bucket: "%v")
	|> range (start: -1h)
	|> filter (fn: (r) => r._measurement == "%v")
	|> filter (fn: (r) => r._field == "%v")
	|> filter (fn: (r) => r.UserID == "%v")
    |> last()
	`, ts.Bucket, PositionMeasurement, fieldType, userID))
	if err != nil {
		log.Println(err)
		return err
	}
	if rs.Next() && fieldType == "Longitude" {
		position.Longitude = rs.Record().ValueByKey("_value").(float32)
	}
	if rs.Next() && fieldType == "Latitude" {
		position.Latitude = rs.Record().ValueByKey("_value").(float32)
	}
	return nil
}

func (ts *TrackerService) UpdateLocation(parent context.Context, position *models.Position) error {
	c := ts.cb.InfluxDbClient()
	defer c.Close()
	ctx, cancel := context.WithTimeout(parent, time.Second*5)
	defer cancel()
	writeApi := c.WriteAPIBlocking(ts.Org, ts.Bucket)
	p := influxdb2.NewPoint(PositionMeasurement, map[string]string{
		"UserID": fmt.Sprintf("%v", position.UID),
	}, map[string]interface{}{
		"Latitude":  position.Latitude,
		"Longitude": position.Longitude,
	}, time.Now())

	return writeApi.WritePoint(ctx, p)
}

func (ts *TrackerService) Checkpoint(parent context.Context, position *models.Position) error {
	c := ts.cb.InfluxDbClient()
	defer c.Close()
	ctx, cancel := context.WithTimeout(parent, time.Second*5)
	defer cancel()
	writeApi := c.WriteAPIBlocking(ts.Org, ts.Bucket)
	p := influxdb2.NewPoint(PositionMeasurement, map[string]string{
		"UserID":     fmt.Sprintf("%v", position.UID),
		"Checkpoint": fmt.Sprintf("%v", position.CheckPointID),
	}, map[string]interface{}{
		"Latitude":  position.Latitude,
		"Longitude": position.Longitude,
	}, time.Now())

	return writeApi.WritePoint(ctx, p)
}

func (ts *TrackerService) UpdateBloodPressure(parent context.Context, bp *models.BloodPressureWithUserID) (*models.Alert, error) {
	// improve this client frequent creation
	c := ts.cb.InfluxDbClient()
	defer c.Close()
	ctx, cancel := context.WithTimeout(parent, time.Second*10)
	defer cancel()
	writeApi := c.WriteAPIBlocking(ts.Org, ts.Bucket)
	queryApi := c.QueryAPI(ts.Org)
	p := influxdb2.NewPoint(BloodPressureMeasurement, map[string]string{
		"UserID": fmt.Sprintf("%v", bp.UserID),
	}, map[string]interface{}{
		"Systolic":  bp.BP.Systolic,
		"Diastolic": bp.BP.Diastolic,
	}, time.Now())
	err := writeApi.WritePoint(ctx, p)
	if err != nil {
		return nil, err
	}

	url, err := url.Parse("https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent")
	if err != nil {
		return nil, err
	}
	rc := Request{
		Contents: []Content{
			{
				Parts: []Part{
					{
						Text: "",
					},
				},
			},
		},
	}
	hc := http.Client{}
	req := http.Request{
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
		URL: url,
		Body:
	}
	hc.Do()
}

type Request struct {
	Contents []Content `json:"contents,omitempty"`
}

type Content struct {
	Parts []Part `json:"parts,omitempty"`
}
type Part struct {
	Text string `json:"text,omitempty"`
}
/*
{
  "contents": [{
    "parts":[
		{
			"text": "Explain how AI works"
		}
		]
    }]
   }
 */