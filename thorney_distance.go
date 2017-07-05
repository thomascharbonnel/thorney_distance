package main

import (
  "math"
)

// Store a geolocation
type Location struct {
  lat float64
  long float64
}

// Helper function to take care of errors
func handleErrors(err error) {
  if err != nil {
    panic(err)
  }
}

// Converts degrees into radians
func degreesToRadians(degrees float64) float64 {
  for degrees >= 360.0 {
    degrees -= 360
  }

  return degrees * (math.Pi / 180)
}

// Haversine distance computation
// Takes two locations and returns the distance between the two
func haversineDistance(loc0 Location, loc1 Location) float64 {
  lat0 := degreesToRadians(loc0.lat)
  lat1 := degreesToRadians(loc1.lat)
  delta_lat := degreesToRadians(loc1.lat - loc0.lat)
  delta_long := degreesToRadians(loc1.long - loc0.long)
  r := 6365.079 // Earth radius under London

  return 2.0 * r * math.Asin(math.Sqrt(math.Pow(delta_lat / 2.0, 2.0) + math.Cos(lat1) * math.Cos(lat0) * math.Pow(delta_long / 2.0, 2.0)))
}

// Fetches location data from a string postcode
func getLocationFromPostcode(postcode string) Location {
  return Location{0, 0}
}

func main() {

}
