package lps

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{s: "babadbb"},
			want: "aba",
		},
		{
			name: "Test Case 2",
			args: args{s: "sababad"},
			want: "ababa",
		},
		{
			name: "Test Case 3",
			args: args{s: "abbcba"},
			want: "bcb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
