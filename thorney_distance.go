package main

import (
  "os"
  "math"
  "errors"
  "fmt"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
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
  db, err := sql.Open("sqlite3", "./postcode.sl3")
  handleErrors(err)
  defer db.Close()

  var location Location
  rows, err := db.Query(fmt.Sprintf("select lat, long from postcode where code='%s' limit 1", postcode))
  rows.Next()
  handleErrors(rows.Scan(&location.lat, &location.long))

  return location
}

// Returns the distance between two postcodes
func distanceBetweenPostcodes(postcode0, postcode1 string) float64 {
  loc0 := getLocationFromPostcode(postcode0)
  loc1 := getLocationFromPostcode(postcode1)

  return haversineDistance(loc0, loc1)
}

func main() {
  if len(os.Args) != 3 {
    handleErrors(errors.New("Program takes two arguments, the two postcodes."))
  }

  postcode0 := os.Args[1]
  postcode1 := os.Args[2]

  distance := distanceBetweenPostcodes(postcode0, postcode1)
  distance_in_miles := distance / 1.609344

  fmt.Printf("The distance between postcodes %s and %s is %f km (%f mi).", postcode0, postcode1, distance, distance_in_miles)
}
