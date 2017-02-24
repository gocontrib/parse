package parse

import "testing"

func TestParseDuration(t *testing.T) {
	Duration("")
	MustDuration("11s")
	MustDuration("11sec")
	MustDuration("11min")
	MustDuration("11h")
	MustDuration("1hour")
	MustDuration("11hours")
	MustDuration("11d")
	MustDuration("1day")
	MustDuration("11days")
	MustDuration("2h45m")
	MustDuration("100")
	MustDuration("1month")
	MustDuration("11months")
	MustDuration("1week")
	MustDuration("11weeks")
	MustDuration("11w")
}
