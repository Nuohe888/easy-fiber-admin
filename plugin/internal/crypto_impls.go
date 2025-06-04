package internal

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

// CryptoMD5Impl MD5加密实现
type CryptoMD5Impl struct{}

func (c *CryptoMD5Impl) EncryptPassword(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("密码不能为空")
	}
	
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:]), nil
}

func (c *CryptoMD5Impl) VerifyPassword(password, hashedPassword string) bool {
	encrypted, err := c.EncryptPassword(password)
	if err != nil {
		return false
	}
	return encrypted == hashedPassword
}

func (c *CryptoMD5Impl) GetAlgorithm() string {
	return "md5"
}

// CryptoSHA1Impl SHA1加密实现
type CryptoSHA1Impl struct{}

func (c *CryptoSHA1Impl) EncryptPassword(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("密码不能为空")
	}
	
	hash := sha1.Sum([]byte(password))
	return hex.EncodeToString(hash[:]), nil
}

func (c *CryptoSHA1Impl) VerifyPassword(password, hashedPassword string) bool {
	encrypted, err := c.EncryptPassword(password)
	if err != nil {
		return false
	}
	return encrypted == hashedPassword
}

func (c *CryptoSHA1Impl) GetAlgorithm() string {
	return "sha1"
}

// CryptoSHA256Impl SHA256加密实现
type CryptoSHA256Impl struct{}

func (c *CryptoSHA256Impl) EncryptPassword(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("密码不能为空")
	}
	
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:]), nil
}

func (c *CryptoSHA256Impl) VerifyPassword(password, hashedPassword string) bool {
	encrypted, err := c.EncryptPassword(password)
	if err != nil {
		return false
	}
	return encrypted == hashedPassword
}

func (c *CryptoSHA256Impl) GetAlgorithm() string {
	return "sha256"
}

// CryptoSHA512Impl SHA512加密实现
type CryptoSHA512Impl struct{}

func (c *CryptoSHA512Impl) EncryptPassword(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("密码不能为空")
	}
	
	hash := sha512.Sum512([]byte(password))
	return hex.EncodeToString(hash[:]), nil
}

func (c *CryptoSHA512Impl) VerifyPassword(password, hashedPassword string) bool {
	encrypted, err := c.EncryptPassword(password)
	if err != nil {
		return false
	}
	return encrypted == hashedPassword
}

func (c *CryptoSHA512Impl) GetAlgorithm() string {
	return "sha512"
}
