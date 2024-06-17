package ringarr

import (
	"reflect"
	"testing"
)

func Test_max(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{a: []int{1, 2, 3, 4, 5}},
			want: []int{2, 3, 4, 5, -1},
		},
		{
			name: "test case 2",
			args: args{a: []int{2, 1, 2, 4, 3}},
			want: []int{4, 2, 4, -1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}
