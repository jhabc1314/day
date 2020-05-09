package rand

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	//"math/big"
)

//Package rand实现了加密安全的随机数生成器。

/**
阅读器是密码安全的随机数生成器的全局共享实例。

在Linux和FreeBSD上，Reader使用getrandom（2）（如果可用），否则使用/ dev / urandom。 在OpenBSD上，Reader使用getentropy（2）。 在其他类似Unix的系统上，Reader从/ dev / urandom中读取。 在Windows系统上，Reader使用CryptGenRandom API。 在Wasm上，Reader使用Web Crypto API。
*/
var Reader io.Reader
//Int 以[0，最大值）返回统一的随机值。如果max <= 0，则表示恐慌
//func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)

//素数返回给定大小的数字p，因此p是极有可能的素数。 如果rand.Read返回的任何错误或bit <2，Prime将返回错误。
//func Prime(rand io.Reader, bits int)(p *big.Int, err error)

//Read 是一个辅助函数，它使用io.ReadFull调用Reader.Read。返回时，当且仅当err == nil时，n == len（b）。
//func Read(b []byte) (n int, err error)


func RandDemo() {
	c := 10
	b := make([]byte, c)

	_,err := rand.Read(b)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(bytes.Equal(b, make([]byte, c)))

	fmt.Println(b)
}

