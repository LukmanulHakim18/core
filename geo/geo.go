package geo

import (
	"math"
)

// Fungsi untuk mengubah derajat ke radian
func toRadians(degree float64) float64 {
	return degree * math.Pi / 180
}

// Fungsi untuk menghitung jarak antara dua koordinat
func GetDistanceInMeterFromCoordinate(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000 // Radius bumi dalam meter

	// Mengubah derajat ke radian
	lat1Rad := toRadians(lat1)
	lon1Rad := toRadians(lon1)
	lat2Rad := toRadians(lat2)
	lon2Rad := toRadians(lon2)

	// Menghitung perbedaan antara koordinat
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	// Menghitung jarak menggunakan rumus Haversine
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Menghitung jarak akhir dalam meter
	distance := R * c
	return distance
}
