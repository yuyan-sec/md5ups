使用AI 把 perl 脚本修改成 go ，方便在 windows 下使用。

```
爆破md5: go run main.go -f pass.txt -p d6ddbdeba60446cd1a732e8148eba29c -s 111 -u admin
生成md5: go run main.go -p 123456 -s 111 -u admin
```
> 添加生成md5 方便可以连接数据库但是解不开管理员密码，可以生成 md5 来替换登录后台

![1](https://github.com/yuyan-sec/md5ups/assets/43353917/257a329a-8fc5-456f-8c81-cbb82956615b)



# md5ups 若依密码爆破
爆破md5(用户名+密码+salt)的脚本
最近一直遇见若依cms后台总能获取所有账号密码手机号salt等字段。
![image](https://user-images.githubusercontent.com/46959313/141808155-4c8da579-6b0f-4d6a-b5db-ead5708521a3.png)

可是cmd5解不开这种加salt的md5
![image](https://user-images.githubusercontent.com/46959313/141808289-840c53b3-6ec4-4746-8d50-a04784287540.png)


百度得知new Md5Hash(username + password + salt).toHex();

密码为：用户名+密码+salt 的md5

本来打算用burp添加payload规则

可是处理的很慢，后尝试hashcat，看说明好像不支持前后掩码

然后github找轮子，也没有相关脚本，但是发现了一个类似的脚本

https://github.com/wdsjxh/md5_salt_brute_force

根据大佬的脚本修改后完成，速度快的令人不敢相信
![image](https://user-images.githubusercontent.com/46959313/141808814-9176bb2e-fc31-48ea-b53e-85a284bc2285.png)

不过测试几遍 确实几万字典也是秒出

后续如果遇见其他md5加盐的，大佬们可以根据脚本修改自行修改食用
