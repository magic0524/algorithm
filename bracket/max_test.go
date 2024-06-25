package bracket

import "testing"

func Test_bracketMax(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{s: "(()"},
			want: 2,
		},
		{
			name: "test case 2",
			args: args{s: ")()())"},
			want: 4,
		},
		{
			name: "test case 3",
			args: args{s: ")(()"},
			want: 2,
		},
		{
			name: "test case 4",
			args: args{s: ""},
			want: 0,
		},
		{
			name: "test case 5",
			args: args{s: "(())"},
			want: 4,
		},
		{
			name: "test case 6",
			args: args{s: ")(()))()()(()))"},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bracketMax(tt.args.s); got != tt.want {
				t.Errorf("bracketMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
