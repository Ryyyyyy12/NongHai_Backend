package enum

import (
	"encoding/json"
	"fmt"
)

type AnimalType string

var (
	Dog AnimalType = "dog"
	Cat AnimalType = "cat"
)

func (a AnimalType) UnmarshalJSON(b []byte) error {
	//unmarshal json
	v := new(string)
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}

	//validate value
	val := AnimalType(*v)
	if val != Dog && val != Cat {
		return fmt.Errorf("invalid animal type", val)
	}

	//set value
	a = val
	return nil
}
