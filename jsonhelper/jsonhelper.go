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

func (jsonfield jsonField) Bool() (result bool) {
	result, _ = jsonfield.data.(bool)
	return
}

func (jsonfield jsonField) String() (result string) {
	result, _ = jsonfield.data.(string)
	return
}

func (jsonfield jsonField) Number() (result float64) {
	result, _ = jsonfield.data.(float64)
	return
}

func (jsonfield jsonField) Array() (result []JSONHelper) {
	result = []JSONHelper{}

	data, _ := jsonfield.data.([]interface{})
	for _, value := range data {
		result = append(result, jsonField{value})
	}
	return result
}

func (jsonfield jsonField) Map() (result map[string]JSONHelper) {
	result = map[string]JSONHelper{}

	data, _ := jsonfield.data.(map[string]interface{})
	for key, value := range data {
		result[key] = jsonField{value}
	}
	return result
}
