package scalar

import (
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
	"time"
)

const DATE_TIME_FORMAT = "2006-01-02 15:04:05"

func MarshalDateTime(t time.Time) graphql.Marshaler {
	if t.IsZero() {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Format(DATE_TIME_FORMAT)))
	})
}

func UnmarshalDateTime(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		return time.Parse(DATE_TIME_FORMAT, tmpStr)
	}
	return time.Time{}, errors.New("time should be 2006-01-02 15:04:05 formatted string")
}
