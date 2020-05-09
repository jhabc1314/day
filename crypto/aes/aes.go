package aes

import "crypto/cipher"

/**
包AES实现了AES加密（以前称为Rijndael），如美国联邦信息处理标准出版物197中所定义。

此软件包中的AES操作未使用固定时间算法实现。 在启用了对AES的硬件支持的系统上运行时，这些操作将保持恒定，这是一个例外。 示例包括使用AES-NI扩展的amd64系统和使用Message-Security-Assist扩展的s390x系统。 在这样的系统上，当将NewCipher的结果传递给cipher.NewGCM时，GCM使用的GHASH操作也是恒定时间的。
*/

//AES块大小（以字节为单位）。
const BlockSize = 16

//NewCipher 创建并返回一个新的cipher.Block。 key参数应为16、24或32个字节的AES密钥，以选择AES-128，AES-192或AES-256。
func NewCipher(key []byte)(cipher.Block, error)

type KeySizeError int

func (k KeySizeError) Error() string
