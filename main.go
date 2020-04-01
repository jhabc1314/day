package main

import (
	"github.com/jhabc1314/day/archive_go"
	"github.com/jhabc1314/day/bufio"
	"github.com/jhabc1314/day/builtin"
	"github.com/jhabc1314/day/bytes"
	"github.com/jhabc1314/day/container"
)
//单引号则用于表示Golang的一个特殊类型：rune，类似其他语言的byte但又不完全一样，是指：码点字面量（Unicode code point），不做任何转义的原始内容
func main() {
    archive_go.Tar()
    archive_go.Untar("./cache/output.tar")

    archive_go.Zip()
    archive_go.Unzip("./cache/output.zip")

    bufio.Read("./tar_file1.txt")
    bufio.Write()

    builtin.Builtin()

    bytes.Byte()

    h := container.IntHeap{9,8,11,2,3,5,7,7,15}
    container.HeapFunc(&h) //有序的队列，插入进去自动会排序

    container.ListFunc()
}
