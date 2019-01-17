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
	return jsonfield.data.(bool)
}

func (jsonfield jsonField) String() string {
	return jsonfield.data.(string)
}

func (jsonfield jsonField) Number() float64 {
	return jsonfield.data.(float64)
}

func (jsonfield jsonField) Array() []JSONHelper {
	result := make([]JSONHelper, 0)
	for _, value := range jsonfield.data.([]interface{}) {
		result = append(result, jsonField{value})
	}
	return result
}

func (jsonfield jsonField) Map() map[string]JSONHelper {
	result := make(map[string]JSONHelper)
	for key, value := range jsonfield.data.(map[string]interface{}) {
		result[key] = jsonField{value}
	}
	return result
}
