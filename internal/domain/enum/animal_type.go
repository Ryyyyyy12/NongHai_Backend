package enum

import (
	"encoding/json"
	"fmt"
)

type AnimalType string

var (
	Dog AnimalType = "Dog" // Use capital "D"
	Cat AnimalType = "Cat" // Use capital "C"
)

func (a *AnimalType) UnmarshalJSON(b []byte) error {
	// Unmarshal json
	v := new(string)
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}

	// Validate value
	val := AnimalType(*v)
	if val != Dog && val != Cat {
		return fmt.Errorf("invalid animal type: %s", val) // Use correct format string
	}

	// Set value
	*a = val // Update the original value using pointer dereference
	return nil
}
