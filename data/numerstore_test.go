package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNumberStore(t *testing.T) {
	store := NewNumberStore(nil)
	assert.NotNil(t, store, "NewNumberStore should return a non-nil store")
}

func TestLoadNumbers(t *testing.T) {
	store := NewNumberStore(nil)
	err := store.LoadNumbers("testdata/test_input.txt")
	assert.NoError(t, err, "LoadNumbers should not return an error for valid input")
	assert.NotEmpty(t, store.Numbers, "Numbers slice should not be empty after loading valid data")
}

func TestFindIndex(t *testing.T) {
	store := NewNumberStore(nil)
	store.Numbers = []int{10, 20, 30, 40, 50}

	tests := []struct {
		name     string
		value    int
		expected int
		found    bool
	}{
		{"ValuePresent", 30, 2, true},
		{"ValueAbsent", 25, -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index, found := store.FindIndex(tt.value)
			assert.Equal(t, tt.found, found, "FindIndex found flag mismatch")
			assert.Equal(t, tt.expected, index, "FindIndex index mismatch")
		})
	}
}

func TestFindClosestIndex(t *testing.T) {
	store := NewNumberStore(nil)
	store.Numbers = []int{10, 20, 30, 40, 50, 1100, 1200, 1500}

	tests := []struct {
		name      string
		value     int
		tolerance float64
		expected  int
		found     bool
	}{
		{"ExactMatch", 30, 0.10, 2, true},
		{"CloseMatch", 35, 0.20, 3, true},
		{"NoMatch", 35, 0.10, -1, false},
		{"MatchAtToleranceEdge", 33, 0.10, 2, true},
		{"MatchBelow", 1150, 0.10, 5, true},
		{"MatchAbove", 1400, 0.10, 7, true},
		{"OutOfRangeBelow", 5, 0.10, -1, false},
		{"OutOfRangeAbove", 56, 0.10, -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index, found := store.FindClosestIndex(tt.value, tt.tolerance)
			assert.Equal(t, tt.expected, index, "FindClosestIndex returned wrong index")
			assert.Equal(t, tt.found, found, "FindClosestIndex found flag mismatch")
		})
	}
}
