package hpp

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
)

type JSONInt int

func (i *JSONInt) MarshalJSON() ([]byte, error) {
	js, err := json.Marshal(i.Int())
	if err != nil {
		return nil, err
	}
	es := base64.StdEncoding.EncodeToString(js)
	return json.Marshal(es)
}

func (b JSONInt) Int() int {
	return int(b)
}

func (b JSONInt) String() string {
	return strconv.Itoa(b.Int())
}
