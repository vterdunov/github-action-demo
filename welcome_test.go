package main

import "testing"

func Test_welcome(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		want     string
	}{
		{
			name:     "with simple name",
			userName: "John",
			want:     "Welcome, John",
		},
		{
			name:     "with empty name",
			userName: "",
			want:     "Welcome, ",
		},
		{
			name:     "with special characters",
			userName: "O'Connor",
			want:     "Welcome, O'Connor",
		},
		{
			name:     "with unicode name",
			userName: "Иван",
			want:     "Welcome, Иван",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := welcome(tt.userName); got != tt.want {
				t.Errorf("welcome() = %v, want %v", got, tt.want)
			}
		})
	}
}
