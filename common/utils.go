package common

import (
	"encoding/json"
)

func JsonMarshal(data any) []byte {
	jsonData, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		panic(err)
	}
	return jsonData
}
