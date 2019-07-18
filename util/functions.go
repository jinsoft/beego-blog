package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func GUID() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5Signer(base64.URLEncoding.EncodeToString(b))
}

func UUID() string {
	str := NewString(GUID())
	builder := NewStringBuilder()
	builder.Append(str.Substr(0, 8).ToString())
	builder.Append("-").Append(str.Substr(8, 12).ToString())
	builder.Append("-").Append(str.Substr(12, 16).ToString())
	builder.Append("-").Append(str.Substr(16, 20).ToString())
	builder.Append("-").Append(str.SubstrBegin(20).ToString())
	return builder.ToString()
}

/**
md5 sing: "123456" => ""
*/
func Md5Signer(message string) string {
	data := []byte(message)
	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash)
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
