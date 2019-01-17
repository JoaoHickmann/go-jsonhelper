package main

import (
	"fmt"

	"github.com/JoaoHickmann/go-jsonhelper/jsonhelper"
)

func main() {
	json := `
		{
			"string": "string",
			"number": 2.5,
			"bool": true,
			"array": [
				1, 
				2,
				3
			],
			"object": {
				"key": "value"
			}
		}
	`

	jsonHelper, err := jsonhelper.NewJSONHelper([]byte(json))
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonHelper.Map()["string"].String())
	fmt.Println(jsonHelper.Map()["number"].Number())
	fmt.Println(jsonHelper.Map()["bool"].Bool())
	fmt.Println(jsonHelper.Map()["array"].Array()[0].Number())
	fmt.Println(jsonHelper.Map()["object"].Map()["key"].String())
}
