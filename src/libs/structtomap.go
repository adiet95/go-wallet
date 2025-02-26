package libs

import "encoding/json"

func StructToMap(obj interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
