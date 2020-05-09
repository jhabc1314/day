package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

//包密码实现标准的分组密码模式，可以将其包装在低级分组密码实现中。参见https://csrc.nist.gov/groups/ST/toolkit/BCM/current_modes.html和NIST特殊出版物800-38A。

//AEAD 是一种密码模式，可对关联数据进行身份验证加密。
type AEAD interface {

	//NonceSize 返回必须传递给Seal and Open的随机数的大小。 nonce 随机数
	NonceSize() int

	//开销返回明文长度与密文长度之间的最大差值。
	Overhead() int

	//Seal 对明文进行加密和身份验证，对其他数据进行身份验证，然后将结果附加到dst，返回更新后的切片。 对于给定的密钥，随机数必须为NonceSize（）字节长，并且在所有时间内都是唯一的。
	//要将纯文本的存储重新用于加密输出，请使用纯文本[：0]作为dst。 否则，dst的剩余容量不得与明文重叠。
	// Seal encrypts and authenticates plaintext, authenticates the
    // additional data and appends the result to dst, returning the updated
    // slice. The nonce must be NonceSize() bytes long and unique for all
    // time, for a given key.
    //
    // To reuse plaintext's storage for the encrypted output, use plaintext[:0]
    // as dst. Otherwise, the remaining capacity of dst must not overlap plaintext.
	Seal(dst,nonce,plaintext,additionalData []byte) []byte

	/**
	Open会对密文进行解密和身份验证，对其他数据进行身份验证，如果成功，则将生成的明文附加到dst，返回更新后的切片。 随机数必须为NonceSize（）字节长，并且它和其他数据都必须与传递给Seal的值匹配。
	要将密文的存储重新用于解密的输出，请使用密文[：0]作为dst。 否则，dst的剩余容量不得与明文重叠。
	即使该函数失败，dst的内容（直至其容量）也可能会被覆盖。
	*/
	// Open decrypts and authenticates ciphertext, authenticates the
    // additional data and, if successful, appends the resulting plaintext
    // to dst, returning the updated slice. The nonce must be NonceSize()
    // bytes long and both it and the additional data must match the
    // value passed to Seal.
    //
    // To reuse ciphertext's storage for the decrypted output, use ciphertext[:0]
    // as dst. Otherwise, the remaining capacity of dst must not overlap plaintext.
    //
    // Even if the function fails, the contents of dst, up to its capacity,
	// may be overwritten.
	Open(dst, nonce, ciphertext,additionalData []byte) []byte
}

//NewGCM 返回以标准随机数长度在Galois计数器模式下包装的给定128位块密码。 Galois Counter Mode
//通常，此GCM实施执行的GHASH操作不是固定时间。 一个例外是当底层aes由aes.NewCipher在具有AES硬件支持的系统上创建时。 有关详细信息，请参见crypto / aes软件包文档。
//func NewGCM(cipher Block)(AEAD,error)

//Decrypt 解密demo
func Decrypt() {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	key,_ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")

	ciphertext,_ := hex.DecodeString("c3aaa29f002ca75870806e44086700f62ce4d43e902b3888e23ceff797a7a471")

	nonce,_ := hex.DecodeString("64a9433eae7ccceee2fc0eda")

	block,err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm,err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	plaintext,err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s\n", plaintext)
}

//Encrypt 加密demo
func Encrypt() {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).

	key,_ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	plaintext := []byte("exampleplaintext")

	block,err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	//随机数
	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12);
	if _,err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm,err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}
	//加密
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("%x\n", ciphertext)
}

//NewGCMWithNonceSize returns the given 128-bit, block cipher wrapped in Galois Counter Mode, which accepts nonces of the given length. The length must not be zero.
//仅在需要与使用非标准随机数长度的现有密码系统兼容时，才使用此功能。所有其他用户都应使用NewGCM，该速度更快且更易于使用。
//func NewGCMWithNonceSize(cipher Block, size int) (AEAD, error)

//NewGCMWithTagSize 返回以Galois计数器模式包装的给定的128位块密码，该密码将生成具有给定长度的标签。
//Tag sizes between 12 and 16 bytes are allowed.
//Only use this function if you require compatibility with an existing cryptosystem that uses non-standard tag lengths. All other users should use NewGCM, which is more resistant to misuse.
//func NewGCMWithTagSize(cipher Block, tagSize int) (AEAD, error)

//A Block represents an implementation of block cipher using a given key. It provides the capability to encrypt or decrypt individual blocks. The mode implementations extend that capability to streams of blocks.
//块表示使用给定密钥的块密码的实现。它提供了加密或解密单个块的功能。模式实现将该功能扩展到块流。
type Block interface {
	//BlockSize returns the cipher's block size.
	BlockSize() int

	// Encrypt encrypts the first block in src into dst.
    // Dst and src must overlap entirely or not at all.
	Encrypt(dst, src []byte)
	//Dst和src必须完全重叠或完全不重叠
	Decrypt(dst, src []byte)
}

//BlockMode 表示以基于块的模式（CBC，ECB等）运行的块密码。
type BlockMode interface {

	BlockSize() int

	// CryptBlocks encrypts or decrypts a number of blocks. The length of
    // src must be a multiple of the block size. Dst and src must overlap
    // entirely or not at all.
    //
    // If len(dst) < len(src), CryptBlocks should panic. It is acceptable
    // to pass a dst bigger than src, and in that case, CryptBlocks will
    // only update dst[:len(src)] and will not touch the rest of dst.
	//多次调用CryptBlocks的行为就像一次运行中传递了src缓冲区的串联一样。 也就是说，BlockMode保持状态，并且不会在每次CryptBlocks调用时重置。
	CryptBlocks(dst,src []byte)
}

//NewCB​​CDecrypter 返回一个BlockMode，它使用给定的Block以密码块链接模式解密。 iv的长度必须与块的块大小相同，并且必须与用于加密数据的iv匹配。
//func NewCBCDecrypter(b Block, iv []byte) BlockMode

//func NewCBCEncrypter(b Block, iv []byte) BlockMode

//A Stream represents a stream cipher. 表示流密码
type Stream interface {
	/*
	XORKeyStream对给定片中的每个字节与密码密钥流中的一个字节进行XOR。 Dst和src必须完全重叠或完全不重叠。

	  如果len（dst）<len（src），则XORKeyStream应该会惊慌。 传递大于src的dst是可以接受的，在这种情况下，XORKeyStream只会更新dst [：len（src）]而不会接触dst的其余部分。

	多次调用XORKeyStream的行为就像一次运行中传递了src缓冲区的串联一样。 也就是说，Stream保持状态，并且不会在每次XORKeyStream调用时重置。
	*/
	XORKeyStream(dst, src []byte)
}


//NewCFBDecrypter 返回一个Stream，该Stream使用给定的Block以密码反馈模式解密。 iv的长度必须与块的块大小相同。
//func NewCFBDecrypter(blcok Block, iv []byte) Stream

//NewCFBEncrypter 返回一个Stream，该Stream使用给定的Block以密码反馈模式进行加密。 iv的长度必须与块的块大小相同。
//func NewCFBEncrypter(block Block, iv []byte) Stream

//加密需要的几个参数 盐key 块对象block iv向量 待加密的数据 plaintext

//NewCTR 返回一个Stream，该Stream在计数器模式下使用给定的Block进行加密/解密。 iv的长度必须与块的块大小相同。
//func NewCTR(block Block, iv []byte) Stream

//NewOFB 返回在输出反馈模式下使用块密码b加密或解密的Stream。初始化向量iv的长度必须等于b的块大小。
//func NewOFB(block Block, iv []byte) Stream

//StreamReader wraps a Stream into an io.Reader. It calls XORKeyStream to process each slice of data which passes through.
//StreamReader将Stream包装到io.Reader中。它调用XORKeyStream处理通过的每个数据片。
type StreamReader struct {
	S Stream
	R io.Reader
}

//func (r StreamReader) Read(dst []byte) (n int, err error)

//StreamWriter 将Stream包装到io.Writer中。 它调用XORKeyStream处理通过的每个数据片。 如果任何Write调用返回的时间很短，则StreamWriter将不同步，必须将其丢弃。 StreamWriter没有内部缓冲。 不需要调用Close来刷新写入数据。
type StreamWriter struct {
    S Stream
    R io.Reader
}

//func (w StreamWriter) Close() error
//如果Writer也是io.Closer，则Close将关闭基础Writer并返回其Close返回值。否则返回nil。

//func (w StreamWriter) Write(src []byte) (n int, err error)
