package scalar

import (
	"VehicleSupervision/pkg/model"
	"encoding/json"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"io"
)

func MarshalJsonbArray(t model.JsonbArray) graphql.Marshaler {
	if t == nil {
		return graphql.Null
	}
	v, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, string(v))
	})
}

func UnmarshalJsonbArray(v interface{}) ([]map[string]interface{}, error) {
	if tmpStr, ok := v.(string); ok {
		var r = model.JsonbArray{}
		err := json.Unmarshal([]byte(tmpStr), &r)
		if err != nil {
			panic(err)
		}
		return r, err
	}
	return nil, errors.New("unexcept jsonb array format")
}
