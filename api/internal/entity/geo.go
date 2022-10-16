package entity

type GeoPoint struct {
	Lng float32
	Lat float32
}

func NewGeoPoint(lng float32, lat float32) *GeoPoint {
	return &GeoPoint{
		Lng: lng,
		Lat: lat,
	}
}

type NearestVerticy struct {
	Id       uint64
	Point    GeoPoint
	Distance float32
}

func NewNearestVerticy(id uint64, point GeoPoint, distance float32) *NearestVerticy {
	return &NearestVerticy{
		Id:       id,
		Point:    point,
		Distance: distance,
	}
}
