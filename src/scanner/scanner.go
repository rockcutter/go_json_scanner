package scanner

import (
	"encoding/json"
	"fmt"
)

type JsonScanner struct {
	rawJson []byte
}

func NewJsonScanner(jsonStr []byte) *JsonScanner {
	return &JsonScanner{rawJson: jsonStr}
}

func (s *JsonScanner) Get(key string) *JsonScanner {
	var payload interface{}

	err := json.Unmarshal(s.rawJson, &payload)
	if err != nil {
		return nil
	}

	m := payload.(map[string]interface{})
	value, ok := m[key]
	if !ok {
		return nil
	}

	rs := &JsonScanner{[]byte(fmt.Sprintf("%v", value))}

	// json.Unmarshal を調整するよりもここで分岐するほうが楽
	// 必要になったら調整する
	switch vt := value.(type) {
	case map[string]interface{}:
		b, err := json.Marshal(vt)
		if err != nil {
			return nil
		}
		rs = &JsonScanner{[]byte(b)}
	case string:
		rs = &JsonScanner{[]byte(vt)}
	case float64:
		rs = &JsonScanner{[]byte(fmt.Sprintf("%f", vt))}
	default:
		rs = &JsonScanner{[]byte("Unsupported type")}
	}

	return rs
}

func (s *JsonScanner) ToString() *string {
	str := string(s.rawJson)
	return &str
}
