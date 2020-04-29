package crypto

import "io"

//包加密收集常见的加密常数。

//RegisterHash注册一个函数，该函数返回给定哈希函数的新实例。旨在从实现散列函数的程序包中的init函数中调用它。
//func RegisterHash(h Hash, f func() hash.Hash)

//解密器是不透明私钥的接口，可用于非对称解密操作。一个示例是保存在硬件模块中的RSA密钥。

type Decrypter interface {
	Public() PublicKey // public 返回与不透明对应的公钥私钥

	//解密加密文本
	Decrypt(rand io.Reader, msg []byte, opts DecrypterOpts) (plaintext []byte, err error)
}

type DecrypterOpts interface {}

type Hash uint

const (
	MD4 Hash = iota + 1 
	MD5
	SHA1
	
)

type PublicKey interface {}

type PrivateKey interface {}

