package scanner_test

import (
	"testing"

	"github.com/rockcutter/go_json_scannner/src/scanner"
)

func TestScanner_Get(t *testing.T) {
	testCase := []struct {
		name    string
		jsonStr []byte
		keys    []string
		want    string
	}{
		{
			name:    "simple_1",
			jsonStr: []byte(`{"name": "John", "age": 30}`),
			keys:    []string{"name"},
			want:    "John",
		},
		{
			name:    "simple_2",
			jsonStr: []byte(`{"name": "John", "age": 30}`),
			keys:    []string{"age"},
			want:    "30.000000",
		},
		{
			name:    "ref_nested_json",
			jsonStr: []byte(`{"name": "John", "age": 30, "address": {"city": "New York", "zip": 10001}}`),
			keys:    []string{"address"},
			want:    `{"city":"New York","zip":10001}`,
		},
		{
			name:    "ref_nested_json_element",
			jsonStr: []byte(`{"name": "John", "age": 30, "address": {"city": "New York", "zip": 10001}}`),
			keys:    []string{"address", "city"},
			want:    "New York",
		},
		{
			name:    "ref_nested_json_element",
			jsonStr: []byte(`{"name": "John", "age": 30, "address": {"city": "New York", "zip": 10001}}`),
			keys:    []string{"address", "zip"},
			want:    "10001.000000",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			s := scanner.NewJsonScanner(tc.jsonStr)
			for _, key := range tc.keys {
				s = s.Get(key)
			}
			got := *(s.ToString())
			if got != tc.want {
				t.Errorf("got: %s, want: %s", got, tc.want)
			}
		})
	}
}

func TestScanner_IsInvalid(t *testing.T) {
	testCase := []struct {
		name    string
		jsonStr []byte
		keys    []string
		want    bool
	}{
		{
			name:    "valid",
			jsonStr: []byte(`{"name": "John", "age": 30}`),
			keys:    []string{"name"},
			want:    false,
		},
		{
			name:    "invalid_key",
			jsonStr: []byte(`{"name": "John", "age": 30}`),
			keys:    []string{"this_field_is_not_exist"},
			want:    true,
		},
		{
			name:    "invalid_json_format",
			jsonStr: []byte(`{"name": "John",hoge=hoge "age": 30, "address": {"city": "New York", "zip": 10001}}`),
			keys:    []string{"name"},
			want:    true,
		},
		{
			name:    "invalid_key_nested_json",
			jsonStr: []byte(`{"name": "John", "age": 30, "address": {"city": "New York", "zip": 10001}}`),
			keys:    []string{"address", "this_field_is_not_exist"},
			want:    true,
		},
		{
			name:    "invalid_first_key_nested_json",
			jsonStr: []byte(`{"name": "John", "age": 30, "address": {"city": "New York", "zip": 10001}}`),
			keys:    []string{"this_field_is_not_exist", "city"},
			want:    true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			s := scanner.NewJsonScanner(tc.jsonStr)
			for _, key := range tc.keys {
				s = s.Get(key)
			}
			got := s.IsInvalid()
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func TestScanner_PanicForNonObjectJson(t *testing.T) {
	jsonStr := []byte(`[1, 2, 3]`)
	s := scanner.NewJsonScanner(jsonStr)
	result := s.Get("any")
	if !result.IsInvalid() {
		t.Errorf("配列のみのJSONに対しては動作しない。")
	}
}
