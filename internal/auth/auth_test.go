package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name: "successful api key extraction",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-api-key"},
			},
			expectedKey:   "test-api-key",
			expectedError: nil,
		},
		{
			name:          "missing authorization header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)
			
			if err != tt.expectedError {
				t.Errorf("GetAPIKey() error = %v, expected error %v", err, tt.expectedError)
				return
			}
			
			if key != tt.expectedKey {
				t.Errorf("GetAPIKey() key = %v, expected key %v", key, tt.expectedKey)
			}
		})
	}
} 