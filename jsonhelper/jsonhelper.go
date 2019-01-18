package jsonhelper

import (
	"encoding/json"
	"errors"
)

type jsonField struct {
	data interface{}
}

type JSONHelper interface {
	Data() interface{}

	Bool() bool
	String() string
	Number() float64
	Map() map[string]JSONHelper
	Array() []JSONHelper
}

func NewJSONHelper(data []byte) (JSONHelper, error) {
	var jsonResponse map[string]interface{}
	err := json.Unmarshal(data, &jsonResponse)
	if err != nil {
		var jsonResponse2 []interface{}
		err := json.Unmarshal(data, &jsonResponse2)
		if err != nil {
			return nil, errors.New("invalid JSON")
		}
		return jsonField{jsonResponse2}, nil
	}
	return jsonField{jsonResponse}, nil
}

func (jsonfield jsonField) Data() interface{} {
	return jsonfield.data
}

func (jsonfield jsonField) Bool() bool {
	result, _ := jsonfield.data.(bool)
	return result
}

func (jsonfield jsonField) String() string {
	result, _ := jsonfield.data.(string)
	return result
}

func (jsonfield jsonField) Number() float64 {
	result, _ := jsonfield.data.(float64)
	return result
}

func (jsonfield jsonField) Array() []JSONHelper {
	result := []JSONHelper{}

	data, ok := jsonfield.data.([]interface{})
	if !ok {
		return result
	}

	for _, value := range data {
		result = append(result, jsonField{value})
	}
	return result
}

func (jsonfield jsonField) Map() map[string]JSONHelper {
	result := map[string]JSONHelper{}

	data, ok := jsonfield.data.(map[string]interface{})
	if !ok {
		return result
	}

	for key, value := range data {
		result[key] = jsonField{value}
	}
	return result
}
