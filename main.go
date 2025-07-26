package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"os"
	"strings"
)

// checkPassword 检查密码是否匹配
func checkPassword(username, plainText, salt, cipherText string) bool {
	checkPassword := md5.Sum([]byte(username + plainText + salt))
	checkPasswordStr := fmt.Sprintf("%x", checkPassword)
	return cipherText == checkPasswordStr
}

// 生成密码
func generatePassword(username, plainText, salt string) string {
	generatePassword := md5.Sum([]byte(username + plainText + salt))
	generatePasswordStr := fmt.Sprintf("%x", generatePassword)
	return generatePasswordStr
}

func main() {
	var (
		passwordFile string
		cipherText   string
		salt         string
		username     string
		hashFile     string
	)

	flag.StringVar(&passwordFile, "f", "", "file path for passwords")
	flag.StringVar(&cipherText, "p", "", "md5(username+pass+salt),complex pass")
	flag.StringVar(&salt, "s", "", "the salt value for pass")
	flag.StringVar(&username, "u", "", "the username value for pass")
	flag.StringVar(&hashFile, "h", "", "username:salt:password  crack")
	flag.Parse()

	fmt.Println("用于爆破若依cms账号密码[MD5(username+password+salt)]")
	fmt.Println(`example:
	爆破md5: main.exe -f pass.txt -p d6ddbdeba60446cd1a732e8148eba29c -s 111 -u admin
	生成md5: main.exe -p 123456 -s 111 -u admin
	批量爆破(username:password:salt 格式) main.exe -f pass.txt -h hash.txt
	`)

	// 生成
	if passwordFile == "" {

		if username != "" {
			fmt.Println(generatePassword(username, cipherText, salt))
		}

		os.Exit(0)
	}

	passwords, err := os.ReadFile(passwordFile)
	if err != nil {
		fmt.Printf("无法打开密码文件: %s\n", err)
		os.Exit(0)
	}

	passwordList := strings.Split(string(passwords), "\n")

	// 批量爆破
	if hashFile != "" {
		crack(passwordList, hashFile)
		os.Exit(0)
	}

	var flag bool
	for _, password := range passwordList {
		password = strings.TrimSpace(password)
		if password == "" {
			continue
		}

		if checkPassword(username, password, salt, cipherText) {
			flag = true
			fmt.Printf("%s 爆破成功，密码：%s\n", username, password)
			break
		}

	}

	if !flag {
		fmt.Println("抱歉，找不到密码！")
	}
}

func crack(passwordList []string, hashFile string) {

	var userList []string
	if hashFile != "" {
		users, err := os.ReadFile(hashFile)
		if err != nil {
			fmt.Printf("读取 hash 文件失败：%v\n", err)
			return
		}
		lines := strings.Split(string(users), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" {
				userList = append(userList, line)
			}
		}
	}

	for _, password := range passwordList {
		password = strings.TrimSpace(password)
		if password == "" {
			continue
		}

		for _, usp := range userList {
			uspList := strings.Split(string(usp), ":")

			u := uspList[0]
			p := uspList[1]
			s := uspList[2]

			if checkPassword(u, password, s, p) {
				fmt.Printf("%s 爆破成功，密码：%s\n", u, password)
				break
			}
		}

	}

	fmt.Println("爆破结束")

}
