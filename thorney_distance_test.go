package main

import (
  "testing"
)

// Test degreesToRadians
func TestDegreesToRadians(t *testing.T) {
  values := []struct {
    input float64 // value in degrees
    output float64 // expected output
  }{
    {0, 0.0},
    {90, 1.5708},
    {180, 3.141593},
    {360, 0.0},
  }

  for i, v := range values {
    actual := degreesToRadians(v.input)

    // We allow 1‰ as margin of error through rounding
    if actual > 1.001 * v.output || actual < 0.999 * v.output {
      t.Fatalf("Expected %f in iteration %d, got %f.", v.output, i, actual)
    }
  }
}

// Test the haversineDistance between London and Paris
// The result should be around 343 km.
func TestHaversineDistance(t *testing.T) {
  london := Location{51.5073, -0.12755}
  paris := Location{48.853363, 2.348749}

  result := haversineDistance(london, paris)

  // We allow 1% as margin of error
  if result > 346.43 || result < 339.57 {
    t.Fatalf("Expected %d but got %d.", 343, result)
  }
}

func TestGetLocationFromPostcode(t *testing.T) {
  result := getLocationFromPostcode("LL53")

  if result.lat != 52.88418 {
    t.Fatalf("Expected %f but got %f.", 52.88418, result)
  }
}

func TestDistanceBetweenPostcodes(t *testing.T) {
  result := distanceBetweenPostcodes("AB10", "ZE3")
  expected_distance := 309.3

  if result > expected_distance * 1.01 || result < expected_distance * 0.99 {
    t.Fatalf("Expected %f but got %f.", expected_distance, result)
  }
}
