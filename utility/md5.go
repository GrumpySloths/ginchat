package utility

/*用于实现密码加密*/
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func Md5Encode(data string) string {
	h := md5.New()

	return hex.EncodeToString(h.Sum([]byte(data)))
}

func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 加密
func MakeEncode(data, salt string) string {
	return MD5Encode(data + salt)
}

// 验证加密
// @plainwd为未加密的用户输入密码，用于用户注册后的密码验证
func ValidEncode(plainwd, salt, passwd string) bool {
	fmt.Println(plainwd + "    " + salt + "   " + passwd + "   " + MD5Encode(plainwd+salt))
	return MD5Encode(plainwd+salt) == passwd
}
