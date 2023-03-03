package helpers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/constraints"
)

func MapIntFromString[V constraints.Integer](s, sep, sepKV string) (result map[string]V, err error) {
	parts := strings.Split(s, sep)
	result = make(map[string]V, len(parts))

	for _, v := range parts {
		kv := strings.Split(v, sepKV)
		if len(kv) != 2 {
			return nil, errors.New("invalid string format, should be 'key" + sepKV + "val" + sep + "key" + sepKV + "val'")
		}
		i, err := strconv.ParseInt(kv[1], 10, 64)
		if err != nil {
			return nil, err
		}
		result[kv[0]] = V(i)
	}
	return
}

func MapFloatFromString[V constraints.Float](s, sep, sepKV string) (result map[string]V, err error) {
	parts := strings.Split(s, sep)
	result = make(map[string]V, len(parts))

	for _, v := range parts {
		kv := strings.Split(v, sepKV)
		if len(kv) != 2 {
			return nil, errors.New("invalid string format, should be 'key" + sepKV + "val" + sep + "key" + sepKV + "val'")
		}
		i, err := strconv.ParseFloat(kv[1], 64)
		if err != nil {
			return nil, err
		}
		result[kv[0]] = V(i)
	}
	return
}

func MapStringFromString[V ~string](s, sep, sepKV string) (result map[string]V, err error) {
	parts := strings.Split(s, sep)
	result = make(map[string]V, len(parts))
	for _, v := range parts {
		kv := strings.Split(v, sepKV)
		if len(kv) != 2 {
			return nil, errors.New("invalid string format, should be 'key" + sepKV + "val" + sep + "key" + sepKV + "val'")
		}
		result[kv[0]] = V(kv[1])
	}
	return
}

func MapTimeFromString[V time.Time](s, sep, sepKV, layout string) (result map[string]V, err error) {
	parts := strings.Split(s, sep)
	result = make(map[string]V, len(parts))
	for _, v := range parts {
		kv := strings.Split(v, sepKV)
		if len(kv) != 2 {
			return nil, errors.New("invalid string format, should be 'key" + sepKV + "val" + sep + "key" + sepKV + "val'")
		}
		t, err := time.Parse(layout, kv[1])
		if err != nil {
			return nil, fmt.Errorf("invalid date format for key %s: %w", kv[0], err)
		}
		result[kv[0]] = V(t)
	}
	return
}

func MapDurationFromString[V time.Duration](s, sep, sepKV string) (result map[string]V, err error) {
	parts := strings.Split(s, sep)
	result = make(map[string]V, len(parts))
	for _, v := range parts {
		kv := strings.Split(v, sepKV)
		if len(kv) != 2 {
			return nil, errors.New("invalid string format, should be 'key" + sepKV + "val" + sep + "key" + sepKV + "val'")
		}
		t, err := time.ParseDuration(kv[1])
		if err != nil {
			return nil, fmt.Errorf("invalid duration format for key %s: %w", kv[0], err)
		}
		result[kv[0]] = V(t)
	}
	return
}

func MapUUIDFromString[V uuid.UUID](s, sep, sepKV string) (result map[string]V, err error) {
	parts := strings.Split(s, sep)
	result = make(map[string]V, len(parts))
	for _, v := range parts {
		kv := strings.Split(v, sepKV)
		if len(kv) != 2 {
			return nil, errors.New("invalid string format, should be 'key" + sepKV + "val" + sep + "key" + sepKV + "val'")
		}
		t, err := uuid.Parse(kv[1])
		if err != nil {
			return nil, fmt.Errorf("invalid duration format for key %s: %w", kv[0], err)
		}
		result[kv[0]] = V(t)
	}
	return
}
