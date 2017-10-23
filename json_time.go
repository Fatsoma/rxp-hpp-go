package hpp

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// TimeLayout is the format required for realex HPP (in ISO8601 YYYYMMDDHHMMSS)
const TimeLayout = "20060102150405"

// JSONTime is a wrapper around time.Time to override the marshal / unmarshal json functions
type JSONTime time.Time

func (jt JSONTime) String() string {
	return fmt.Sprintf("%s", time.Time(jt).Format(TimeLayout))
}

// MarshalJSON converts the time to the TimeLayout
func (jt JSONTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(jt.String())
}

// UnmarshalJSON converts TimeLayout formatted strings
func (jt *JSONTime) UnmarshalJSON(b []byte) error {
	if jt == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}

	parsedTime, err := time.Parse(TimeLayout, strings.Trim(string(b), `"`))
	if err != nil {
		return err
	}

	*jt = JSONTime(parsedTime)

	return nil
}
