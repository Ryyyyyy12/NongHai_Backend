package enum

import (
	"encoding/json"
	"fmt"
)

type Sex string

var (
	Male   Sex = "male"
	Female Sex = "female"
)

func (s Sex) UnmarshalJSON(b []byte) error {
	//unmarshal json
	v := new(string)
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}

	//validate value
	val := Sex(*v)
	if val != Male && val != Female {
		return fmt.Errorf("invalid sex type", val)
	}

	//set value
	s = val
	return nil
}
