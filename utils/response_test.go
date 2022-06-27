package utils

import "testing"

func TestResponse(t *testing.T) {
	type args struct {
		Status int
		Data   interface{}
		Error  string
	}

	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "TestResponse with valid arguments",
			args: args{
				Status: 1,
				Data:   "Hello, World!",
				Error:  "",
			},
			want: Response{
				Status: 1,
				Data:   "Hello, World!",
				Error:  "",
			},
		},
		{
			name: "TestResponse with missing arguments",
			args: args{
				Status: 1,
				Data:   "Hello, World!",
			},
			want: Response{
				Status: 1,
				Data:   "Hello, World!",
			},
		},
		{
			name: "TestResponse ",
			args: args{
				Status: 0,
				Data:   "",
				Error:  "Please complete missing fields",
			},
			want: Response{
				Status: 0,
				Data:   "",
				Error:  "Please complete missing fields",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResponse(tt.args.Status, tt.args.Data, tt.args.Error); got != tt.want {
				t.Errorf("NewResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewResponseUtil(t *testing.T) {
	resp1 := NewResponse(1, "Hello, World!", "")
	resp2 := NewResponse(0, "", "Please complete missing fields")

	if resp1.Status != 1 || resp1.Data != "Hello, World!" || resp1.Error != "" {
		t.Errorf("NewResponseUtil() = %v, want %v", resp1, Response{Status: 1, Data: "Hello, World!", Error: ""})
	}

	if resp2.Status != 0 || resp2.Data != "" || resp2.Error != "Please complete missing fields" {
		t.Errorf("NewResponseUtil() = %v, want %v", resp2, Response{Status: 0, Data: "", Error: "Please complete missing fields"})
	}
}
