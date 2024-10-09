package enum

import (
	"encoding/json"
	"fmt"
)

type Status string

var (
	Lost    Status = "Lost"
	Safe       Status = "Safe"
)

func (s *Status) UnmarshalJSON(b []byte) error {
	// Unmarshal json
	v := new(string)
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}

	// Validate value
	val := Status(*v)
	if val != Lost && val != Safe{
		return fmt.Errorf("invalid status type: %s", val)
	}

	// Set value
	*s = val
	return nil
}
