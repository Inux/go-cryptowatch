package common

import (
	"bytes"
	"io"
	"time"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetBytes(reader io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	return buf.Bytes()
}

// GetCorrectedTime returns the middle of two time values
func GetCorrectedTime(start, end time.Time) time.Time {
	duration := end.Sub(start)
	return start.Add(duration / 2)
}
