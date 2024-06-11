package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers http.Header
		expect  string
	}{
		"no authorization header": {
			headers: http.Header{
				"Content-Type": []string{"application/json"},
			},
			expect: "",
		},
		"wrong prefix": {
			headers: http.Header{
				"Authorization": []string{"Bearer 123456"},
			},
			expect: "",
		},
		"only prefix": {
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expect: "",
		},
		"only prefix trailing spaces": {
			headers: http.Header{
				"Authorization": []string{"ApiKey      "},
			},
			expect: "",
		},
		"valid key": {
			headers: http.Header{
				"Authorization": []string{"ApiKey 123456"},
			},
			expect: "123456",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := GetAPIKey(tc.headers)
			diff := cmp.Diff(got, tc.expect)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
