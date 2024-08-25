package enum

import (
	"encoding/json"
	"fmt"
)

type MessageType string

var (
	Text  MessageType = "text"
	Image MessageType = "image"
)

func (m MessageType) UnmarshalJSON(b []byte) error {
	//unmarshal json
	v := new(string)
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}

	//validate value
	val := MessageType(*v)
	if val != Text && val != Image {
		return fmt.Errorf("invalid sex type", val)
	}

	//set value
	m = val
	return nil
}
