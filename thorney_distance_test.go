package main

import (
  "testing"
)

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
