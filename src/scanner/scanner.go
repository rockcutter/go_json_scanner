package scanner

import (
	"encoding/json"
	"fmt"
)

type JsonScanner struct {
	rawJson   []byte
	isInvalid bool
}

func NewJsonScanner(jsonStr []byte) *JsonScanner {
	return &JsonScanner{rawJson: jsonStr, isInvalid: false}
}

func newInvalidJsonScanner(jsonStr []byte) *JsonScanner {
	return &JsonScanner{rawJson: jsonStr, isInvalid: true}
}

func (s *JsonScanner) IsInvalid() bool {
	return s.isInvalid
}

func (s *JsonScanner) SetInvalid() {
	s.isInvalid = true
}

func (s *JsonScanner) Get(key string) *JsonScanner {
	var payload interface{}

	err := json.Unmarshal(s.rawJson, &payload)
	if err != nil {
		return newInvalidJsonScanner([]byte(""))
	}

	m, ok := payload.(map[string]interface{})
	if !ok {
		return newInvalidJsonScanner([]byte(""))
	}
	value, ok := m[key]
	if !ok {
		return newInvalidJsonScanner([]byte(""))
	}

	rs := NewJsonScanner([]byte(fmt.Sprintf("%v", value)))

	// json.Unmarshal を調整するよりもここで分岐するほうが楽
	// 必要になったら調整する
	switch vt := value.(type) {
	case map[string]interface{}:
		b, err := json.Marshal(vt)
		if err != nil {
			return newInvalidJsonScanner([]byte(""))
		}
		rs = NewJsonScanner(b)
	case string:
		rs = NewJsonScanner([]byte(vt))
	case float64:
		rs = NewJsonScanner([]byte(fmt.Sprintf("%f", vt)))
	default:
		rs = NewJsonScanner([]byte(""))
	}

	return rs
}

func (s *JsonScanner) ToString() *string {
	str := string(s.rawJson)
	return &str
}
