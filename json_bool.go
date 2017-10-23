package hpp

import (
	"encoding/json"
	"fmt"
)

// JSONBool is a boolean that represents "1" and "0" as true / false
type JSONBool bool

// MarshalJSON converts bools to "1" / "0"
func (b *JSONBool) MarshalJSON() ([]byte, error) {
	result := "0"
	if *b {
		result = "1"
	}

	return json.Marshal(result)
}

// UnmarshalJSON converts "1" / "0" to bool
func (b *JSONBool) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "1" || s == "true" {
		*b = true
	} else if s == "0" || s == "false" {
		*b = false
	} else {
		return fmt.Errorf("Boolean unmarshal error: invalid input %s", s)
	}
	return nil
}
