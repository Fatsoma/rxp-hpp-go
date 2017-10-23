package hpp

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type JSONTime time.Time

func (jt JSONTime) String() string {
	return fmt.Sprintf("%s", time.Time(jt).Format(TimeLayout))
}

func (jt JSONTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(jt.String())
}

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
