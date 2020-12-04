package scalar

import (
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
	"time"
)

const TIME_FORMAT = "15:04:05"

func MarshalTime(t time.Time) graphql.Marshaler {
	if t.IsZero() {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Format(TIME_FORMAT)))
	})
}

func UnmarshalTime(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		return time.Parse(TIME_FORMAT, tmpStr)
	}
	return time.Time{}, errors.New("time should be 15:04:05 formatted string")
}
