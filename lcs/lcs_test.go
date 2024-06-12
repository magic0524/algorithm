package lcs

import (
	"testing"
)

func Test_longestCommonSubstring(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{"abcdbcaa", "abc"},
			want: 3,
		},
		{
			name: "Test case 2",
			args: args{"abcdbcaa", "abcb"},
			want: 3,
		},
		{
			name: "Test case 3",
			args: args{"abcdbcaa", "cdcaa"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubstring(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("longestCommonSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonSubsequnece(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{"abcdbcaa", "abc"},
			want: 3,
		},
		{
			name: "Test case 2",
			args: args{"abcdbcaa", "acba"},
			want: 4,
		},
		{
			name: "Test case 3",
			args: args{"abcdbcaa", "adsbca"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequnece(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("longestCommonSubsequnece() = %v, want %v", got, tt.want)
			}
		})
	}
}
