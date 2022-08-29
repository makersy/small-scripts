package main

import "small-scripts/mystring"

const dir = "/Users/yhl/develop/markdown/blog" // 根目录

//main 按需更改dir
func main() {
	mystring.ReplaceTextForFilesIn(dir)
}
