package routes

import "testing"

func TestResponse(t *testing.T) {
	type args struct {
		Data  interface{}
		Error string
	}

	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "TestResponse with valid arguments",
			args: args{
				Data:  "Hello, World!",
				Error: "",
			},
			want: Response{
				Data:  "Hello, World!",
				Error: "",
			},
		},
		{
			name: "TestResponse with missing arguments",
			args: args{
				Data: "Hello, World!",
			},
			want: Response{
				Data: "Hello, World!",
			},
		},
		{
			name: "TestResponse ",
			args: args{
				Data:  "",
				Error: "Please complete missing fields",
			},
			want: Response{
				Data:  "",
				Error: "Please complete missing fields",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newResponse(tt.args.Data, tt.args.Error); got != tt.want {
				t.Errorf("NewResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewResponseUtil(t *testing.T) {
	resp1 := newResponse("Hello, World!", "")
	resp2 := newResponse("", "Please complete missing fields")

	if resp1.Data != "Hello, World!" || resp1.Error != "" {
		t.Errorf("NewResponseUtil() = %v, want %v", resp1, Response{Data: "Hello, World!", Error: ""})
	}

	if resp2.Data != "" || resp2.Error != "Please complete missing fields" {
		t.Errorf("NewResponseUtil() = %v, want %v", resp2, Response{Data: "", Error: "Please complete missing fields"})
	}
}
