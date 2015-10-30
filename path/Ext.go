// func Ext(path string) string
// 参数列表
// path 表示路径字符串
// 返回值：
// 返回string
// 功能说明：
// 这个函数主要是返回path中文件的扩展名 -规则：
// 1.最后一个点开始的扩展名
// 2.如果没有点，返回空
// 官方文档：
// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot in the final slash-separated element of path;
// it is empty if there is no dot.

package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Ext("/a/b/c/bar.css")) // .css
	fmt.Println(path.Ext("/a/b/c/bar"))     // ""
}
