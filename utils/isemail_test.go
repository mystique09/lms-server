package utils

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	type args struct {
		email string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid email",
			args: args{
				email: "testemail@gotmail.com",
			},
			want: true,
		},
		{
			name: "Valid email",
			args: args{
				email: "mystique09@gmail.com",
			},
			want: true,
		},
		{
			name: "Invalid email",
			args: args{
				email: "awdwd@",
			},
			want: false,
		},
		{
			name: "Invalid email",
			args: args{
				email: "testinginvalid@@email",
			},
			want: false,
		},
		{
			name: "Valid email",
			args: args{
				email: "testemail@gotmail",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmail(tt.args.email); got != tt.want {
				t.Errorf("IsEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
