package des

import (
	"crypto/cipher"
	"crypto/des"

	cipherutil "github.com/wind-coco/go-witls/cipher"
)

// TripleCipher //DES 和 3DES加密区别
// 前者 加密  密钥必须是8byte
// 后者加密 解密 再加密  密钥必须是24byte
type TripleCipher struct {
	key   string
	iv    string
	block cipher.Block
}

func New(key, iv string) (*TripleCipher, error) {
	block, err := des.NewTripleDESCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	return &TripleCipher{
		key:   key,
		iv:    iv,
		block: block,
	}, nil
}

func (this *TripleCipher) Encrypt(src []byte) []byte {
	// 填充明文
	src = cipherutil.PKCS5Padding(src, this.block.BlockSize())
	// 创建一个CBC模式的加密算法
	mode := cipher.NewCBCEncrypter(this.block, []byte(this.iv))
	// 加密
	dst := make([]byte, len(src))
	mode.CryptBlocks(dst, src)
	return dst
}

func (this *TripleCipher) Decrypt(dst []byte) ([]byte, error) {
	// 解密
	decrypter := cipher.NewCBCDecrypter(this.block, []byte(this.iv))
	decrypted := make([]byte, len(dst))
	decrypter.CryptBlocks(decrypted, dst)
	// 去掉填充
	decrypted, err := cipherutil.PKCS5Unpadding(decrypted)
	return decrypted, err
}
