# whats-wrong-with-nginx

## 行为

- 获取Root用户权限，若没能获取失败，退出程序
- 确认系统是否安装systemd，若没有，使用service代替
- 生成随机数种子
- 循环执行关闭和开启指定程序