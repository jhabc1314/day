package crypto

import (
	"hash"
	"io"
)

//包加密收集常见的加密常数。

//RegisterHash注册一个函数，该函数返回给定哈希函数的新实例。旨在从实现散列函数的程序包中的init函数中调用它。
//func RegisterHash(h Hash, f func() hash.Hash)

//Decrypter 是不透明私钥的接口，可用于非对称解密操作。一个示例是保存在硬件模块中的RSA密钥。
type Decrypter interface {
	Public() PublicKey // public 返回与不透明对应的公钥私钥

	//解密加密文本
	Decrypt(rand io.Reader, msg []byte, opts DecrypterOpts) (plaintext []byte, err error)
}

type DecrypterOpts interface {

}

type Hash uint

const (
	MD4 Hash = iota + 1 
	MD5
	SHA1
	SHA224
	SHA256
	SHA384
	SHA512
	MD5SHA1
	RIPEMD160
	SHA3_224
	SHA_256
	SHA3_384
	SHA3_512
	SHA512_224
	SHA512_256
	BLAKE2s_256
	BLAKE2b_256
	BLAKE2b_384
	BLAKE2b_512
	
)

//PublicKey 使用未指定的算法表示公共密钥。
type PublicKey interface {}

//PrivateKey 使用未指定的算法表示私钥。
type PrivateKey interface {}


//Available reports whether the given hash function is linked into the binary.。
func (h Hash) Available() bool

//HashFunc simply returns the value of h so that Hash implements SignerOpts.
func (h Hash) HashFunc() Hash

//New 返回一个新的hash.Hash计算给定的hash函数。如果哈希函数未链接到二进制文件中，则会引起新的恐慌。
func (h Hash) New() hash.Hash

//Size 返回由给定哈希函数得出的摘要的长度（以字节为单位）。不需要将相关的哈希函数链接到程序中。
func (h Hash) Size() int

//Signer 是不透明私钥的接口，可用于签名操作。例如，保存在硬件模块中的RSA密钥。
type Signer interface {
	//Public 返回与不透明对应的公钥，私钥
	Public() PublicKey

	// Sign signs digest with the private key, possibly using entropy from
    // rand. For an RSA key, the resulting signature should be either a
    // PKCS#1 v1.5 or PSS signature (as indicated by opts). For an (EC)DSA
    // key, it should be a DER-serialised, ASN.1 signature structure.
    //
    // Hash implements the SignerOpts interface and, in most cases, one can
    // simply pass in the hash function used as opts. Sign may also attempt
    // to type assert opts to other types in order to obtain algorithm
    // specific values. See the documentation in each package for details.
    //
    // Note that when a signature of a hash of a larger message is needed,
    // the caller is responsible for hashing the larger message and passing
	// the hash (as digest) and the hash function (as opts) to Sign.
	Sign(rand io.Reader, digest []byte, opts SignerOpts) (signature []byte, err error)

}

//SignerOpts 包含用于使用Signer签名的选项。
type SignerOpts interface {

	//HashFunc 返回用于生成传递给Signer.Sign的消息的哈希函数的标识符，或者返回零以指示未完成哈希。
	HashFunc() Hash
}
