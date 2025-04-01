package service

import (
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"math"
)

// haversine function to calculate distance between two lat/lon points
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371.0 * 1000 // Radius of Earth in Miter

	// Convert degrees to radians
	dLat := (lat2 - lat1) * (math.Pi / 180.0)
	dLon := (lon2 - lon1) * (math.Pi / 180.0)

	lat1 = lat1 * (math.Pi / 180.0)
	lat2 = lat2 * (math.Pi / 180.0)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c // Distance in Miter
}

// totalDistance calculates the total distance covered from index 0 to the last index
func totalDistance(positions []*models.Position) float64 {
	if len(positions) < 2 {
		return 0.0 // No movement if there is only one or zero positions
	}

	totalDist := 0.0
	for i := 1; i < len(positions); i++ {
		totalDist += haversine(
			positions[i-1].Latitude, positions[i-1].Longitude,
			positions[i].Latitude, positions[i].Longitude,
		)
	}
	return totalDist
}
