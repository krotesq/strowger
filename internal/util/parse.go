package util

import (
	"encoding/json"
	"io"
)

// this function takes any type as dst, but usually struct
// can be called like this: ParseBody(r.Body, &rb)
func ParseBody[T any](body io.Reader, dst *T) error {
	dec := json.NewDecoder(body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(dst); err != nil {
		return err
	}
	return nil
}