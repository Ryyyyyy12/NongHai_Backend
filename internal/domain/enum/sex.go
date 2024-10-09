package enum

import (
	"encoding/json"
	"fmt"
)

type Sex string

var (
	Male   Sex = "Male"   // Use capital "M"
	Female Sex = "Female" // Use capital "F"
)

func (s *Sex) UnmarshalJSON(b []byte) error {
	// Unmarshal json
	v := new(string)
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}

	// Validate value
	val := Sex(*v)
	if val != Male && val != Female {
		return fmt.Errorf("invalid sex type: %s", val) // Use correct format string
	}

	// Set value
	*s = val // Update the original value using pointer dereference
	return nil
}
