package service

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db/influx"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	PositionMeasurement  string = "position"
	PulseRateMeasurement string = "pr"
)

type TrackerService struct {
	cb *influx.InfluxDbClientBuilder
	influx.InfluxDBOptions
}

type DistanceCovered struct {
	positions []*models.Position
	mu        sync.RWMutex
}

func NewTrackerService(options *influx.InfluxDBOptions) *TrackerService {
	if options == nil {
		options = &influx.InfluxDBOptions{}
	}
	return &TrackerService{
		cb: influx.NewInfluxDbClientBuilder().WithOrg(options.Org).WithToken(options.Token).WithURL(options.Url).WithBucket(options.Bucket),
	}
}

func NewDistanceCovered() *DistanceCovered {
	return &DistanceCovered{
		mu:        sync.RWMutex{},
		positions: make([]*models.Position, 0),
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
		position.Longitude = rs.Record().ValueByKey("_value").(float64)
	}
	if rs.Next() && fieldType == "Latitude" {
		position.Latitude = rs.Record().ValueByKey("_value").(float64)
	}
	return nil
}

func (ts *TrackerService) UpdateLocation(parent context.Context, position *models.Position) error {
	if position == nil || position.Longitude == 0.0 || position.Latitude == 0.0 {
		return fmt.Errorf("position is not correct")
	}
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
	}, func() time.Time {
		if position.Time == nil {
			return time.Now()
		}
		return *position.Time
	}())

	return writeApi.WritePoint(ctx, p)
}

func (ts *TrackerService) Checkpoint(parent context.Context, position *models.Position) (uint64, error) {
	if position == nil {
		return 0, fmt.Errorf("sent position is nil")
	}
	if position.Longitude == 0.0 || position.Latitude == 0.0 {
		return 0, fmt.Errorf("longitude and latitude not set")
	}
	if position.CheckPointID == 0 {
		position.CheckPointID = rand.Uint64()
	}
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
	}, func() time.Time {
		if position.Time == nil {
			return time.Now()
		}
		return *position.Time
	}())

	return position.CheckPointID, writeApi.WritePoint(ctx, p)
}

func (ts *TrackerService) UpdatePulseRate(parent context.Context, pr *models.PulseRateWithUserID) (*models.UserID, error) {
	// improve this client frequent creation
	c := ts.cb.InfluxDbClient()
	defer c.Close()
	ctx, cancel := context.WithTimeout(parent, time.Second*10)
	defer cancel()
	writeApi := c.WriteAPIBlocking(ts.Org, ts.Bucket)
	p := influxdb2.NewPoint(PulseRateMeasurement, map[string]string{
		"UserID": fmt.Sprintf("%v", pr.UserID),
	}, map[string]interface{}{
		"PulseRate": pr.PulseRate,
	}, time.Now())
	err := writeApi.WritePoint(ctx, p)
	if err != nil {
		return nil, err
	}
	return &pr.UserID, err
}

func (ts *TrackerService) PulseRateAlert(parent context.Context, UserID *models.UserID) (*models.Alert, error) {
	c := ts.cb.InfluxDbClient()
	defer c.Close()
	ctx, cancel := context.WithTimeout(parent, time.Second*10)
	defer cancel()
	queryApi := c.QueryAPI(ts.Org)
	pr, err := queryApi.Query(ctx, fmt.Sprintf(`
		from(bucket: "%v")
		|> range(start: -10m)
		|> filter(fn: (r) => r.measurement == "%v"
		|> filter(fn: (r) => r.UserID == "%v"
		|> mean()
    `, ts.Bucket, PulseRateMeasurement, UserID))
	if err != nil {
		return nil, err
	}
	if pr == nil {
		return nil, nil
	}
	mean := pr.Record().ValueByKey("PulseRate")
	alert := &models.Alert{}
	if mean.(float32) > 1.5 {
		alert.Type = models.HighPulseRate
		// future plan, integrate AI to give suggestions
		alert.Message = "This message will be generated by AI"
	} else if mean.(float32) > 1.2 {
		alert.Type = models.Normal
		alert.Message = "This message will be generated by AI"
	} else {
		alert.Type = models.LowPulseRate
		alert.Message = "This message will be generated by AI"
	}
	return alert, nil
}

func (ts *TrackerService) AppendDistance(pos *models.Position, input chan<- *models.Position) {
	input <- pos
}

func (ts *TrackerService) GetRealTimeDistanceCovered(ctx context.Context, inputStream <-chan *models.Position) <-chan float64 {
	dc := NewDistanceCovered()
	outputStream := make(chan float64)
	go ts.getHaversineDistance(ctx, dc, inputStream, outputStream)
	return outputStream
}

func (ts *TrackerService) getHaversineDistance(ctx context.Context, dc *DistanceCovered, inputStream <-chan *models.Position, outputStream chan float64) {
	workStream := make(chan struct{}, 1)
	// there we can cache previous response obj
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-workStream:
				dc.mu.RLock()
				// this is safe I presume, because we only gonna append in the dc.positions. so shared access races
				distance := dc.positions
				dc.mu.RUnlock()
				outputStream <- totalDistance(distance)
			}
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case np := <-inputStream:
			dc.mu.Lock()
			dc.positions = append(dc.positions, np)
			dc.mu.Unlock()
			workStream <- struct{}{}
		}
	}
}

func (ts *TrackerService) GetTotalDistanceBetweenCheckpoint(parent context.Context, ck *models.CheckpointToAndFrom) (float64, error) {
	ctx, cancel := context.WithTimeout(parent, time.Second*10)
	defer cancel()
	c := ts.cb.InfluxDbClient()
	queryApi := c.QueryAPI(ts.Org)
	rs, err := queryApi.Query(ctx, fmt.Sprintf(`
		from(bucket: "%v")
		|> range(start: -1h)
		|> filter(fn (r) => r.measurement == "%v")
		|> filter(fn (r) => r.Checkpoint == "%v")
		|> last()
    `, ts.Bucket, PositionMeasurement, ck.To))
	if err != nil {
		return 0, err
	}
	if !rs.Next() {
		return 0, fmt.Errorf("no rows found")
	}
	t1 := rs.Record().ValueByKey("_time")
	rs, err = queryApi.Query(ctx, fmt.Sprintf(`
		from(bucket: "%v")
		|> range(start: -1h)
		|> filter(fn (r) => r.measurement == "%v")
		|> filter(fn (r) => r.Checkpoint == "%v")
		|> last()
    `, ts.Bucket, PositionMeasurement, ck.From))
	if err != nil {
		return 0, err
	}
	if !rs.Next() {
		return 0, fmt.Errorf("no rows found")
	}
	t2 := rs.Record().ValueByKey("_time")
	rs, err = queryApi.Query(ctx, fmt.Sprintf(`
		from(bucket: "%v")
		|> range(start: %v, stop: %v)
		|> filter(fn (r) => r.measurement == "%v")
    `, ts.Bucket, t1, t2, PositionMeasurement))
	dc := NewDistanceCovered()
	for rs.Next() {
		dc.positions = append(dc.positions, &models.Position{
			Latitude:  rs.Record().ValueByKey("Latitude").(float64),
			Longitude: rs.Record().ValueByKey("Longitude").(float64),
		})
	}
	return totalDistance(dc.positions), nil
}
