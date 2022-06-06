package storage

import (
	"io/ioutil"
	"testing"
)

func TestImageDownload(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"https://freshman.tech/images/dp-illustration.png", "image/png"},
		{"https://images.unsplash.com/photo-1519681393784-d120267933ba?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1470&q=80", "image/jpeg"},
	}

	for _, test := range tests {
		stream, err := getIcon(test.input)
		if err != nil {
			t.Error(err)
		}

		//read the stream
		bytes, err := ioutil.ReadAll(stream)
		if err != nil {
			t.Error(err)
		}
		stream.Close()

		mimeType, err := getIconType(bytes)
		if err != nil {
			t.Error(err)
		}
		if mimeType != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, mimeType)
		}
	}
}
