package slice

import "go_work/errs"

// Remove 切片删除指定下标元素
func delete(slice []string, index int) ([]string, error) {
	//判断下标是否越界
	if index < 0 || index >= len(slice) {
		return slice, errs.NewErrIndexOutOfRange(len(slice), index)
	}
	return append(slice[:index], slice[index+1:]...), nil
}
