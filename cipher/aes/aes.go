package aes

import (
	"crypto/aes"
	"crypto/cipher"

	"errors"

	cipherutil "github.com/wind-coco/go-witls/cipher"
	"github.com/wind-coco/go-witls/encoding"
	"github.com/wind-coco/go-witls/encoding/noop"
)

type Cipher struct {
	keyBytes  []byte
	iv        []byte
	block     cipher.Block
	blockSize int

	Encoding encoding.EncoderDecoder
}

func NewInsecure(key string) (*Cipher, error) {
	//NewCipher该函数限制了输入key的长度必须为16, 24或者32
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	blockSize := block.BlockSize() // 获取秘钥块的长度
	if err != nil {
		return nil, err
	}

	iv := keyBytes[:blockSize]
	c := &Cipher{
		keyBytes:  keyBytes,
		iv:        iv,
		block:     block,
		blockSize: blockSize,
	}
	c.Encoding = noop.New()
	return c, nil
}

// iv 长度必须为 blockSize
func New(key, iv string) (*Cipher, error) {
	c, err := NewInsecure(key)
	if err != nil {
		return nil, err
	}
	c.iv = []byte(iv)
	if len(c.iv) != c.blockSize {
		return nil, errors.New("len(iv) must be equal to blockSize")
	}
	return c, nil

}

// cbc模式
func (this *Cipher) Encrypt(src []byte) []byte {

	src = cipherutil.PKCS5Padding(src, this.blockSize)       // 补全码
	blockMode := cipher.NewCBCEncrypter(this.block, this.iv) // 加密模式
	encrypted := make([]byte, len(src))                      // 创建数组
	blockMode.CryptBlocks(encrypted, src)                    // 加密

	return this.Encoding.Encode(encrypted)
}
func (this *Cipher) Decrypt(dst []byte) (bytes []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("panic in CryptBlocks")
		}
	}()
	dst, err = this.Encoding.Decode(string(dst))
	if err != nil {
		return nil, err
	}
	//提前检查，不然CryptBlocks会panic
	/*if len(dst)%this.blockSize != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}*/

	// 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(this.block, this.iv) // 加密模式
	decrypted := make([]byte, len(dst))                      // 创建数组
	//CryptBlocks 内部会直接panic,需要注意
	blockMode.CryptBlocks(decrypted, dst)             // 解密
	bytes, err = cipherutil.PKCS5Unpadding(decrypted) // 去除补全码
	return bytes, err
}
