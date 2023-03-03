package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"golang.org/x/exp/constraints"
)

func SliceIntFromString[V constraints.Integer](s, sep string) (result []V, err error) {
	parts := strings.Split(s, sep)
	result = make([]V, len(parts))
	for idx, v := range parts {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		result[idx] = V(i)
	}
	return
}

func SliceFloatFromString[V constraints.Float](s, sep string) (result []V, err error) {
	parts := strings.Split(s, sep)
	result = make([]V, len(parts))
	for idx, v := range parts {
		i, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		result[idx] = V(i)
	}
	return
}

func SliceTimeFromString[V time.Time](s, sep, sepKV, layout string) (result []V, err error) {
	parts := strings.Split(s, sep)
	result = make([]V, len(parts))
	for idx, v := range parts {
		t, err := time.Parse(layout, v)
		if err != nil {
			return nil, fmt.Errorf("invalid date format for index %d: %w", idx, err)
		}
		result[idx] = V(t)
	}
	return
}

func SliceDurationFromString[V time.Duration](s, sep string) (result []V, err error) {
	parts := strings.Split(s, sep)
	result = make([]V, len(parts))
	for idx, v := range parts {
		t, err := time.ParseDuration(v)
		if err != nil {
			return nil, fmt.Errorf("invalid duration format for index %d: %w", idx, err)
		}
		result[idx] = V(t)
	}
	return
}

func SliceUUIDFromString[V uuid.UUID](s, sep string) (result []V, err error) {
	parts := strings.Split(s, sep)
	result = make([]V, len(parts))
	for idx, v := range parts {
		t, err := uuid.Parse(v)
		if err != nil {
			return nil, fmt.Errorf("invalid duration format for key %d: %w", idx, err)
		}
		result[idx] = V(t)
	}
	return
}
