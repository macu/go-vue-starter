package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// StringToInt converts base 10 string to int.
func StringToInt(s string) (int, error) {
	r, err := strconv.ParseInt(s, 10, 64)
	return int(r), err
}

// StringToUint converts base 10 string to uint.
func StringToUint(s string) (uint, error) {
	r, err := strconv.ParseUint(s, 10, 64)
	return uint(r), err
}

// StringToInt64 converts base 10 string to int64.
func StringToInt64(s string) (int64, error) {
	r, err := strconv.ParseInt(s, 10, 64)
	return r, err
}

// StringToInt64NilIfEmpty converts base 10 string to int64,
// and returns nil on err or empty.
func StringToInt64NilIfEmpty(s string) (*int64, error) {
	r, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		e := err.(*strconv.NumError)
		if e.Num == "" {
			// Input was blank; return no error
			return nil, nil
		}
		return nil, err
	}
	return &r, nil
}

// StringToUintNilIfEmpty converts base 10 string to uint,
// and returns nil on err or empty.
func StringToUintNilIfEmpty(s string) (*uint, error) {
	r, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		e := err.(*strconv.NumError)
		if e.Num == "" {
			// Input was blank; return no error
			return nil, nil
		}
		return nil, err
	}
	var u = uint(r)
	return &u, nil
}

// StringToPointerNilIfEmpty returns a pointer to the given string, or nil if given an empty string.
func StringToPointerNilIfEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// StringToBool converts the given string to boolean.
func StringToBool(a string) bool {
	return a == "1" || a == "true"
}

// StringToTimeNilIfEmpty returns a pointer to a time value represented by the given string,
// or nil if the given string is empty.
func StringToTimeNilIfEmpty(s string) (*time.Time, error) {
	if s == "" {
		return nil, nil
	}
	// Parse in default format output by JSON encoding
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// StringToStringArray is for JSON string arrays.
func StringToStringArray(s string) ([]string, error) {
	var out []string
	if len(s) == 0 {
		return out, nil
	}
	err := json.Unmarshal([]byte(s), &out)
	if err != nil {
		return []string{}, fmt.Errorf("parsing string array: %w", err)
	}
	return out, nil
}

// StringToInt64Array is for parsing ID lists where the input is a JSON encoded array of ints.
func StringToInt64Array(s string) ([]int64, error) {
	var out []int64
	if len(s) == 0 {
		return out, nil
	}
	err := json.Unmarshal([]byte(s), &out)
	if err != nil {
		return []int64{}, fmt.Errorf("parsing int array: %w", err)
	}
	return out, nil
}

// ToString converts the given value to a string.
func ToString(v interface{}) string {
	switch vt := v.(type) {
	case int:
		return strconv.FormatInt(int64(vt), 10)
	case int64:
		return strconv.FormatInt(vt, 10)
	case uint:
		return strconv.FormatUint(uint64(vt), 10)
	case uint64:
		return strconv.FormatUint(vt, 10)
	}
	return fmt.Sprint(v)
}

// https://mangatmodi.medium.com/go-check-nil-interface-the-right-way-d142776edef1
func isNil(a interface{}) bool {
	if a == nil {
		return true
	}
	switch reflect.TypeOf(a).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(a).IsNil()
	}
	return false
}
