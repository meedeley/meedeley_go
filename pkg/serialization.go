package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func ToArray(data any) ([]any, error) {
	if slice, ok := data.([]any); ok {
		return slice, nil
	}

	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		result := make([]any, val.Len())
		for i := 0; i < val.Len(); i++ {
			result[i] = val.Index(i).Interface()
		}
		return result, nil
	}

	if val.Kind() == reflect.Struct || val.Kind() == reflect.Map {
		return []any{data}, nil
	}

	return nil, fmt.Errorf("unsupported data type: %T", data)
}

func ToJson(data any) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error serializing to JSON: %v", err)
		return "", err
	}
	return string(jsonData), nil
}
