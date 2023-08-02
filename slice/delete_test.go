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

func TestDeleteV1(t *testing.T) {
	s := []int32{1, 2, 3, 4, 5}
	sli, err := deleteV1(s, 2)
	t.Log(sli)
	t.Logf("%v", sli)
	fmt.Printf("%v,错误是：%v", sli, err)
	s1 := []string{"1", "2", "3", "4", "5"}
	sli1, err1 := deleteV1(s1, 2)
	t.Log(sli1)
	t.Logf("%v", sli1)
	fmt.Printf("%v,错误是：%v", sli1, err1)

}
