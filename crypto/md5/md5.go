package md5

import (
	"crypto/md5"
	"fmt"
	//"hash"
	"io"
	"os"
)

//MD5的块大小（以字节为单位）。
const BlockSize = 64

//MD5校验和的大小（以字节为单位）。
const Size = 16

//New 返回一个新的哈希值，哈希值计算MD5校验和。哈希还实现了encoding.BinaryMarshaler和encoding.BinaryUnmarshaler来封送和解组哈希的内部状态。
//func New() hash.Hash

//Sum 返回数据的MD5校验和
//func Sum(data []byte) [Size]byte


func Md5Demo() {
	h := md5.New()
	io.WriteString(h, "i am jack")
	io.WriteString(h, " and she is milly")
	fmt.Printf("%x\n", h.Sum(nil))

	f,err := os.Open("./readme")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	s := md5.New()
	if _,err := io.Copy(s, f); err != nil {
		panic(err.Error())
	}
	fmt.Printf("%x\n", s.Sum(nil))
}


