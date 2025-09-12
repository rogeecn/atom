package jsonx

import "encoding/json"

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal[T any](data []byte) (*T, error) {
	var v T
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func UnmarshalTo(data []byte, a any) error {
	return json.Unmarshal(data, a)
}

func MustUnmarshal[T any](data []byte) *T {
	v, err := Unmarshal[T](data)
	if err != nil {
		panic(err)
	}
	return v
}
