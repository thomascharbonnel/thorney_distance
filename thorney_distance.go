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
  return 42.0
}

// Fetches location data from a string postcode
func getLocationFromPostcode(postcode string) Location {
  return Location{0, 0}
}

func main() {

}
