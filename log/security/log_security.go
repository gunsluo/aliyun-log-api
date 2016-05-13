package security

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

func ContentMd5(buffer []byte) string {
	hasher := md5.New()
	hasher.Write(buffer)
	return strings.ToUpper(hex.EncodeToString(hasher.Sum(nil)))
}

func HmacSha1(key, message []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}

// 签名
func Signature(accessKeySecret, stringToSign string) string {
	mac := HmacSha1([]byte(accessKeySecret), []byte(stringToSign))
	return base64.StdEncoding.EncodeToString(mac)
}
