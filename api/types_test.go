package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResponse(t *testing.T) {
	tests := []struct {
		name     string
		respType string
		arg      string
		want     Response[string]
	}{
		{
			name:     "Success response",
			respType: "success",
			arg:      "Hello, world",
			want:     newResponse("Hello, world"),
		},
		{
			name:     "Error response",
			respType: "success",
			arg:      "Hello, world",
			want:     newResponse("Hello, world"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.respType == "error" {
				require.Equal(t, tt.arg, tt.want.Error)
			} else {
				require.Equal(t, tt.arg, tt.want.Data)
			}
		})
	}
}

func TestNewResponseUtil(t *testing.T) {
	resp1 := newResponse("Hello, World!")
	resp2 := newError("Please complete missing fields")

	require.Equal(t, "Hello, World!", resp1.Data)
	require.Empty(t, resp1.Error)

	require.Empty(t, resp2.Data)
	require.Equal(t, "Please complete missing fields", resp2.Error)
}
