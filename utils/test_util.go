package util

import (
	"math"
	"testing"
	"time"
)

func throwError(t *testing.T, given interface{}, expected interface{}) {
	t.Errorf("Difference captured:\n\tgiven= %v\n\texpected= %v", given, expected)
}

// AssertTimePrecision Compares time with 10 second precision.
func AssertTimePrecision(t *testing.T, given time.Time, expected time.Time) {
	// 10 second precision
	isEquals := math.Abs(float64(given.Second()-expected.Second())) < 10
	if !isEquals {
		throwError(t, given, expected)
	}
}

// AssertEquals Compare anything comparable via '=' operator.
func AssertEquals(t *testing.T, given interface{}, expected interface{}) {
	if given != expected {
		throwError(t, given, expected)
	}
}
