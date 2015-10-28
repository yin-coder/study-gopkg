/*func Getgroups() ([]int, error)

参数列表
无

返回值：
[]int 调用者所属用户组的所有id
error 返回错误信息对象

功能说明：
这个函数主要是返回调用者所属用户组的所有id*/

package main

import (
	"fmt"
	"os"
)

func main() {
	ids, err := os.Getgroups()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("%v\n", ids)
}
