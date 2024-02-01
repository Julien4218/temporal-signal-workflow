package util

import (
	"encoding/base64"
	"encoding/json"
	"log"
)

func GetBase64Decode(input string) (string, error) {
	if len(input) > 0 {
		rawInput, err := base64.URLEncoding.DecodeString(input)
		if err != nil {
			return "", err
		}
		return string(rawInput), nil
	}
	return "", nil
}

func getJsonDecode(input string) (interface{}, error) {
	var result interface{}
	if len(input) > 0 {
		err := json.Unmarshal([]byte(input), &result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func GetInputParam(input string) interface{} {
	log.Printf("Getting input")
	rawInput, err := GetBase64Decode(input)
	if err != nil {
		// input is not base64 encoded
		rawInput = input
	}
	param, err := getJsonDecode(rawInput)
	if err != nil {
		log.Fatalf("Invalid json input receieved:%s detail:%s\n", rawInput, err.Error())
	}
	log.Printf("Got jsonInput:%s\n", param)
	return param
}
