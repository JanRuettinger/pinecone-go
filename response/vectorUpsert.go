package response

import (
	"encoding/json"
	"io"
)

type VectorUpsert struct {
	Response
	UpsertedCount int64 `json:"upsertedCount"`
}

func (r *VectorUpsert) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *VectorUpsert) SetBody(body io.Reader) error {
	return nil
}

func (r *VectorUpsert) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *VectorUpsert) AcceptContentType() string {
	return "application/json"
}
