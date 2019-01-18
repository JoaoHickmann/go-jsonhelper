package jsonhelper

import (
	"testing"
)

const (
	jsonString string = `
		{
			"string": "string",
			"number": 5.7,
			"bool": true,
			"array": [
				1, 2, 3
			],
			"map": {
				"key": "value"
			}
		}
	`
)

func TestString(t *testing.T) {
	helper, err := NewJSONHelper([]byte(jsonString))
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		data JSONHelper
		want string
	}{
		{helper.Map()["string"], "string"},
		{helper.Map()["number"], ""},
	}

	for _, test := range tests {
		got := test.data.String()
		if got != test.want {
			t.Errorf("JSONHelper.String() was incorrect, got: %s, want: %s.", got, test.want)
		}
	}
}

func TestData(t *testing.T) {
	helper, err := NewJSONHelper([]byte(jsonString))
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		data JSONHelper
		want string
	}{
		{helper.Map()["string"], "string"},
		{helper.Map()["number"], ""},
	}

	for _, test := range tests {
		got := test.data.String()
		if got != test.want {
			t.Errorf("JSONHelper.String() was incorrect, got: %s, want: %s.", got, test.want)
		}
	}
}

func TestNumber(t *testing.T) {
	helper, err := NewJSONHelper([]byte(jsonString))
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		data JSONHelper
		want float64
	}{
		{helper.Map()["number"], 5.7},
		{helper.Map()["string"], 0},
	}

	for _, test := range tests {
		got := test.data.Number()
		if got != test.want {
			t.Errorf("JSONHelper.String() was incorrect, got: %f, want: %f.", got, test.want)
		}
	}
}

func TestBool(t *testing.T) {
	helper, err := NewJSONHelper([]byte(jsonString))
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		data JSONHelper
		want bool
	}{
		{helper.Map()["bool"], true},
		{helper.Map()["string"], false},
	}

	for _, test := range tests {
		got := test.data.Bool()
		if got != test.want {
			t.Errorf("JSONHelper.String() was incorrect, got: %t, want: %t.", got, test.want)
		}
	}
}
