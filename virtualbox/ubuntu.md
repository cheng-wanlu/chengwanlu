# 安装 Ubuntu

- Ubuntu 摘要
![安装Ubuntu 摘要](/images/virtualbox/ubuntu/step01.png)

- 选择 "Try or Install Ubuntu"
![安装Ubuntu 安装](/images/virtualbox/ubuntu/step02.png)

- 点击 "Install Ubuntu" 按钮
![安装Ubuntu 安装](/images/virtualbox/ubuntu/step03.png)

- 选择系统语言
![安装Ubuntu 语言](/images/virtualbox/ubuntu/step04.png)

- 不用做任何修改, 点击 "Continue"
![安装Ubuntu 安装](/images/virtualbox/ubuntu/step05.png)
![安装Ubuntu 安装](/images/virtualbox/ubuntu/step06.png)
![安装Ubuntu 安装](/images/virtualbox/ubuntu/step07.png)

- 选择时区: Shanghai
![安装Ubuntu 时区](/images/virtualbox/ubuntu/step08.png)

- 用户和密码
![安装Ubuntu 用户和密码](/images/virtualbox/ubuntu/step09.png)

- 开始安装
![安装Ubuntu 安装](/images/virtualbox/ubuntu/step10.png)

- 安装完成, 点击 "Restart Now" 重启 
![安装Ubuntu 重启](/images/virtualbox/ubuntu/step11.png)

- 重启成功
![安装Ubuntu 重启成功](/images/virtualbox/ubuntu/step12.png)

- 打开终端
```bash
# 安装openssh-server
sudo apt-get install openssh-server
```
好啦, 现在可以SSH到这台ubuntu啦

- 设置静态IP地址
点击右上方 **以太坊图标**, 并点击 "Settings"
![安装Ubuntu 静态IP](/images/virtualbox/ubuntu/step13.png)
点击 **Ethernet (enp0s8)** 右侧的 **设置** 按钮
![安装Ubuntu 静态IP](/images/virtualbox/ubuntu/step14.png)
在弹出的对话框中, 点击 "IPv4", 然后选择 "Manual", 并设置 **Addresses**
![安装Ubuntu 静态IP](/images/virtualbox/ubuntu/step15.png)
查看是否设置成功
![安装Ubuntu 静态IP](/images/virtualbox/ubuntu/step16.png)