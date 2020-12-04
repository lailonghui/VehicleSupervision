package scalar

import (
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
	"time"
)

const DATE_FORMAT = "2006-01-02"

func MarshalDate(t time.Time) graphql.Marshaler {
	if t.IsZero() {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Format(DATE_FORMAT)))
	})
}

func UnmarshalDate(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		return time.Parse(DATE_FORMAT, tmpStr)
	}
	return time.Time{}, errors.New("time should be 2006-01-02 formatted string")
}
