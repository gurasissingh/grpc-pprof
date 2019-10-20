package json

import (
	"encoding/json"
)

func JsonMarshal() {
	for i := 0; i < 8000; i++ {
		a := map[string]string{
			"key1": "value1",
			"key2":  "value2",
		}
		json.Marshal(a)
	}
}
