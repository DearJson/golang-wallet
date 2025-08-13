package library

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

//加密过程：
//  1、处理数据，对数据进行填充，采用PKCS7（当密钥长度不够时，缺几位补几个几）的方式。
//  2、对数据进行加密，采用AES加密方法中CBC加密模式
//  3、对得到的加密数据，进行base64加密，得到字符串
// 解密过程相反

// 16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法
// key不能泄露
var PwdKey = []byte("Mp8HqAggL2sBrUJIV0SqUEpm1mlxJFkI")

// pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("数据长度为0")
	}

	unPadding := int(data[length-1])
	if unPadding > aes.BlockSize || unPadding == 0 {
		return nil, errors.New("无效的PKCS7填充")
	}

	if length < unPadding {
		return nil, errors.New("数据长度小于填充长度")
	}

	for i := length - unPadding; i < length; i++ {
		if data[i] != byte(unPadding) {
			return nil, errors.New("PKCS7填充验证失败")
		}
	}

	return data[:length-unPadding], nil
}

// AesEncrypt 加密
func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("加密数据不能为空")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("创建加密实例失败: " + err.Error())
	}

	// 创建随机IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, errors.New("生成IV失败: " + err.Error())
	}

	// 填充数据
	paddedData := pkcs7Padding(data, aes.BlockSize)

	// 加密数据
	encrypted := make([]byte, len(paddedData)+aes.BlockSize) // IV + 加密数据
	copy(encrypted[:aes.BlockSize], iv)

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted[aes.BlockSize:], paddedData)

	return encrypted, nil
}

// AesDecrypt 解密
func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	if len(data) < aes.BlockSize {
		return nil, errors.New("加密数据长度不足")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("创建解密实例失败: " + err.Error())
	}

	// 提取IV
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("加密数据长度不是块大小的整数倍")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(data))
	mode.CryptBlocks(decrypted, data)

	return pkcs7UnPadding(decrypted)
}

// EncryptByAes Aes加密后base64编码
func EncryptByAes(data []byte) (string, error) {
	if len(data) == 0 {
		return "", errors.New("加密数据不能为空")
	}

	encrypted, err := AesEncrypt(data, PwdKey)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// DecryptByAes Aes解密
func DecryptByAes(data string) ([]byte, error) {
	if data == "" {
		return nil, errors.New("解密数据不能为空")
	}

	// Base64解码
	encrypted, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, errors.New("Base64解码失败: " + err.Error())
	}

	if len(encrypted) < aes.BlockSize {
		return nil, errors.New("解密数据长度不足")
	}

	decrypted, err := AesDecrypt(encrypted, PwdKey)
	if err != nil {
		return nil, err
	}

	// 去除末尾的空字符和%字符
	decrypted = bytes.TrimRight(decrypted, "\x00%")
	return decrypted, nil
}
