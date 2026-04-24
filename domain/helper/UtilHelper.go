package helper

import (
	"fmt"
	"reflect"
	"strconv"
)

func ParsePrimitive[T any](s string) (T, error) {
	var zero T
	switch any(zero).(type) {
	case string:
		return any(s).(T), nil
	case bool:
		v, err := strconv.ParseBool(s)
		return any(v).(T), err
	case int:
		v, err := strconv.Atoi(s)
		return any(v).(T), err
	case int8:
		v, err := strconv.ParseInt(s, 10, 8)
		return any(int8(v)).(T), err
	case int16:
		v, err := strconv.ParseInt(s, 10, 16)
		return any(int16(v)).(T), err
	case int32:
		v, err := strconv.ParseInt(s, 10, 32)
		return any(int32(v)).(T), err
	case int64:
		v, err := strconv.ParseInt(s, 10, 64)
		return any(v).(T), err
	case uint:
		v, err := strconv.ParseUint(s, 10, 0)
		return any(uint(v)).(T), err
	case uint8:
		v, err := strconv.ParseUint(s, 10, 8)
		return any(uint8(v)).(T), err
	case uint16:
		v, err := strconv.ParseUint(s, 10, 16)
		return any(uint16(v)).(T), err
	case uint32:
		v, err := strconv.ParseUint(s, 10, 32)
		return any(uint32(v)).(T), err
	case uint64:
		v, err := strconv.ParseUint(s, 10, 64)
		return any(v).(T), err

	case float32:
		v, err := strconv.ParseFloat(s, 32)
		return any(float32(v)).(T), err
	case float64:
		v, err := strconv.ParseFloat(s, 64)
		return any(v).(T), err
	}
	return zero, fmt.Errorf("unsupported type")
}

func GetTypeName[T any]() string {
	var zero T
	return reflect.TypeOf(zero).String()
}
