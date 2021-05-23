package basic

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		q    []int
		l    int
		r    int
		want []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"case 1",
			args{
				[]int{5, 2, 3, 1},
				0,
				3,
				[]int{1, 2, 3, 5},
			},
		},
		{
			"case 2",
			args{
				[]int{5, 1, 1, 2, 0, 0},
				0,
				5,
				[]int{0, 0, 1, 1, 2, 5},
			},
		},
		{
			"case 3",
			args{
				[]int{5, 1, 1, 2, 0, 0},
				0,
				3,
				[]int{1, 1, 2, 5, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.q, tt.args.l, tt.args.r)
			if !reflect.DeepEqual(tt.args.want, tt.args.q) {
				t.Fatalf("want: %v, got: %v", tt.args.want, tt.args.q)
			}
		})
	}
}
