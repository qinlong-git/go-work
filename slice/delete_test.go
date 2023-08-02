package slice

import (
	"fmt"
	"testing"
)

func TestDelete(t *testing.T) {
	s := []string{"1", "2", "3", "4", "5"}
	sli, err := delete(s, 2)
	t.Log(sli)
	t.Logf("%v", sli)
	fmt.Printf("%v,错误是：%v", sli, err)
	//output:[1 2 4 5]
	//[1 2 3 4 5]"}
}
