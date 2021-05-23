package basic

import (
	"fmt"
	"testing"
)

func Test_findFirstPostition(t *testing.T) {
	got := findFirstPostition([]int{1, 2, 3, 4, 4, 4, 5}, 4)
	fmt.Println(got)
}

func Test_findLastPostition(t *testing.T) {
	got := findLastPostition([]int{1, 2, 3, 4, 4, 4, 5}, 4)
	fmt.Println(got)
}
