package hpp

import (
	"encoding/base64"
	"encoding/json"
)

type JSONString string

func (s *JSONString) MarshalJSON() ([]byte, error) {
	es := base64.StdEncoding.EncodeToString(s.Bytes())
	return json.Marshal(es)
}

func (s JSONString) Bytes() []byte {
	return []byte(s)
}

func (s JSONString) String() string {
	return string(s)
}
