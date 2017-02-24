package parse

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func MustDuration(s string) time.Duration {
	d, err := Duration(s)
	if err != nil {
		panic(err)
	}
	return d
}

func Duration(s string) (time.Duration, error) {
	n := len(s)
	if n == 0 {
		return time.Duration(0), errors.New("empty input")
	}
	if strings.HasSuffix(s, "months") {
		v, err := Int(s[:n-6])
		return time.Duration(30*24*v) * time.Hour, err
	}
	if strings.HasSuffix(s, "month") {
		v, err := Int(s[:n-5])
		return time.Duration(30*24*v) * time.Hour, err
	}
	if strings.HasSuffix(s, "weeks") {
		v, err := Int(s[:n-5])
		return time.Duration(7*24*v) * time.Hour, err
	}
	if strings.HasSuffix(s, "week") {
		v, err := Int(s[:n-4])
		return time.Duration(7*24*v) * time.Hour, err
	}
	if strings.HasSuffix(s, "w") {
		v, err := Int(s[0 : n-1])
		return time.Duration(7*24*v) * time.Hour, err
	}
	if strings.HasSuffix(s, "hours") {
		v, err := Int(s[:n-5])
		return time.Duration(v) * time.Hour, err
	}
	if strings.HasSuffix(s, "hour") {
		v, err := Int(s[:n-4])
		return time.Duration(v) * time.Hour, err
	}
	if strings.HasSuffix(s, "h") {
		v, err := Int(s[:n-1])
		return time.Duration(v) * time.Hour, err
	}
	if strings.HasSuffix(s, "days") {
		v, err := Int(s[:n-4])
		return 24 * time.Duration(v) * time.Hour, err
	}
	if strings.HasSuffix(s, "day") {
		v, err := Int(s[:n-3])
		return 24 * time.Duration(v) * time.Hour, err
	}
	if strings.HasSuffix(s, "d") {
		v, err := Int(s[:n-1])
		return 24 * time.Duration(v) * time.Hour, err
	}
	if strings.HasSuffix(s, "s") {
		v, err := Int(s[:n-1])
		return time.Duration(v) * time.Second, err
	}
	if strings.HasSuffix(s, "sec") {
		v, err := Int(s[:n-3])
		return time.Duration(v) * time.Second, err
	}
	if strings.HasSuffix(s, "min") {
		v, err := Int(s[:n-3])
		return time.Duration(v) * time.Minute, err
	}
	d, err := time.ParseDuration(s)
	if err == nil {
		return d, nil
	}
	v, err := strconv.ParseInt(s, 10, 32)
	if err == nil {
		return time.Duration(v) * time.Minute, nil
	}
	return time.Duration(0), err
}

func Int(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 32)
}
