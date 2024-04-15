package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
)

// checkPassword 检查密码是否匹配
func checkPassword(username, plainText, salt, cipherText string) bool {
	checkPassword := md5.Sum([]byte(username + plainText + salt))
	checkPasswordStr := hex.EncodeToString(checkPassword[:])
	return cipherText == checkPasswordStr
}

func main() {
	var (
		passwordFile string
		cipherText   string
		salt         string
		username     string
	)

	flag.StringVar(&passwordFile, "p", "", "file path for passwords")
	flag.StringVar(&cipherText, "c", "", "md5(username+pass+salt),complex pass")
	flag.StringVar(&salt, "s", "", "the salt value for pass")
	flag.StringVar(&username, "u", "", "the username value for pass")
	flag.Parse()

	if passwordFile == "" || cipherText == "" || salt == "" || username == "" {
		fmt.Println("用于爆破若依cms账号密码[MD5(username+password+salt)]")
		flag.Usage()
		fmt.Println("example:   go run main.go -p pass.txt -c d6ddbdeba60446cd1a732e8148eba29c -s 111 -u admin")

		os.Exit(0)
	}

	passwords, err := os.ReadFile(passwordFile)
	if err != nil {
		fmt.Printf("无法打开密码文件: %s\n", err)
		os.Exit(0)
	}

	passwordList := strings.Split(string(passwords), "\n")

	var flag bool
	for _, password := range passwordList {
		password = strings.TrimSpace(password)
		if password == "" {
			continue
		}
		if checkPassword(username, password, salt, cipherText) {
			flag = true
			fmt.Printf("爆破成功，密码：%s\n", password)
			break
		}
	}

	if !flag {
		fmt.Println("抱歉，找不到密码！")
	}
}
