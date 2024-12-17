# ECUT 东华理工大学校园网登录

> 只做http请求，不做任何处理，不存储任何信息

## 使用方法(设置开机自启动登录，需要设备先自动连接校园网)

### Windows

1. 下载[ecut.exe](https://github.com/FXAZfung/ECUT-Campus-Network/releases/download/v0.1.0/ecut.exe)
2. 右键点击`ecut.exe`，选择`发送到`->`桌面快捷方式`
3. 右键点击桌面快捷方式，选择`属性`，在`快捷方式`标签页中的`目标`后面加上`-u=账号 -p=密码 -s=运营商`，如
   `C:\Users\username\Desktop\ecut.exe -u="123456" -p="123456" -s="移动"`
4. 按下`CTRL+R`，输入`shell:startup`，将桌面快捷方式拖到打开的文件夹中