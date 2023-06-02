# 安装 CentOS

- CentOS 摘要
![安装CentOS 摘要](/images/virtualbox/centos/step01.png)

- 选择 "Install CentOS 7"
![安装CentOS 安装](/images/virtualbox/centos/step02.png)

- 选择系统语言
![安装CentOS 语言](/images/virtualbox/centos/step03.png)

- 选择默认分区
这里 "INSTALLATION DESTINATION" 出现有感叹号, 要求选择分区方式, 点击进入选择默认就好.
![安装CentOS 分区](/images/virtualbox/centos/step04.png)
![安装CentOS 分区](/images/virtualbox/centos/step05.png)
不用做任何修改, 直接点击Done, 退出.

现在可以看到感叹号消失了.

- 选择安装的系统模式
![安装CentOS 软件](/images/virtualbox/centos/step06.png)
这里说的模式其实是系统的不同内容, 比如说无图形界面的, 有图形界面的.点击"SOFTWARE SELECTION" 进入: 
![安装CentOS 软件](/images/virtualbox/centos/step07.png)
这里可以看到不同模式, 第一个最小安装就是无界面的, 对图形界面有兴趣的可以选择GNOME Desktop. 
选择好后点击Done退出.

- 选择网络连接
![安装CentOS 网络](/images/virtualbox/centos/step08.png)
如果不设置, 默认是无网络连接, 后期要重新设置. 点击 "NETWORK & HOSTNAME" 进入: 
![安装CentOS 网络](/images/virtualbox/centos/step09.png)
刚进入, Ethernet后面的按钮是OFF, 点击变成ON.
![安装CentOS 网络](/images/virtualbox/centos/step10.png)
然后点击Done退出.
可以看到 "IPv4 Address" 变了
![安装CentOS 网络](/images/virtualbox/centos/step11.png)

- 安装摘要
![安装CentOS 摘要](/images/virtualbox/centos/step12.png)
点击 "Begin Installation" 开始安装.

- 用户和密码
![安装CentOS 网络](/images/virtualbox/centos/step13.png)
![安装CentOS 网络](/images/virtualbox/centos/step14.png)
![安装CentOS 网络](/images/virtualbox/centos/step15.png)
![安装CentOS 网络](/images/virtualbox/centos/step16.png)

- 安装完成, 重启
![安装CentOS 网络](/images/virtualbox/centos/step17.png)
![安装CentOS 网络](/images/virtualbox/centos/step18.png)

- 许可协议
![安装CentOS 网络](/images/virtualbox/centos/step19.png)
![安装CentOS 网络](/images/virtualbox/centos/step20.png)

点击 "FINISH CONFIGURATION"

好啦, 现在可以SSH到这台CentOS啦
