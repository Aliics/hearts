package data

import (
	"encoding/json"
	"errors"
)

func unmarshalEvent(bytes []byte) (string, []byte, error) {
	var eventMap map[string]any
	err := json.Unmarshal(bytes, &eventMap)
	if err != nil {
		return "", nil, err
	}

	t, hasType := eventMap["type"]
	if !hasType {
		return "", nil, errors.New("type field missing")
	}
	typeString, typeIsString := t.(string)
	if !typeIsString {
		return "", nil, errors.New("type field is not a string")
	}

	d, hasData := eventMap["data"]
	if !hasData {
		return "", nil, errors.New("data field missing")
	}
	dataMap, dataIsMap := d.(map[string]any)
	if !dataIsMap {
		return "", nil, errors.New("data field is not an object")
	}

	dataBytes, err := json.Marshal(dataMap)
	if err != nil {
		return "", nil, err
	}

	return typeString, dataBytes, nil
}
