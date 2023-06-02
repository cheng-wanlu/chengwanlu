# 在 VirtualBox 中新建虚拟机实例

本文描述了如何使用VirtualBox在计算机上安装Ubuntu和CentOS虚拟机. Ubuntu和CentOS分别是流行的Linux发行版, 可用于各种应用场景. 

## 打开计算机上的 VirtualBox, 点击 "新建(M)" 按钮.
![新建虚拟机 新建](/images/virtualbox/create/step01.png)

## 虚拟电脑名称与操作系统.
根据以下图示进行设置: 
- ubuntu
![新建虚拟机 虚拟电脑名称与操作系统](/images/virtualbox/create/step02_ubuntu.png)
```
名称(N): Ubuntu_01
文件夹(F): C:\VirtualBox VMs
虚拟光盘(I): D:\ubuntu-22.04.2-desktop-amd64.iso
```
- centos
![新建虚拟机 虚拟电脑名称与操作系统](/images/virtualbox/create/step02_centos.png)
```
名称(N): CentOS_01
文件夹(F): C:\VirtualBox VMs
虚拟光盘(I): D:\CentOS-7-x86_64-DVD-2009.iso
```
注意: 手动安装时, 请确保勾选 “跳过自动安装(S)” 选项.
点击 "下一步" 按钮继续.

## 设置内存大小和处理器.
根据以下图示进行设置: 
![新建虚拟机 硬件](/images/virtualbox/create/step03.png)
```
内存大小(M): 4096 MB(4G)
处理器(P): 4
```
点击 "下一步" 按钮继续.

# 分配虚拟硬盘的大小. 
根据以下图示进行设置: 
![新建虚拟机 摘要](/images/virtualbox/create/step04.png)
```
磁盘空间(S): 100 GB
```
点击 "下一步" 按钮继续.

## 新虚拟电脑的配置摘要
![新建虚拟机 摘要](/images/virtualbox/create/step05_ubuntu.png)
![新建虚拟机 摘要](/images/virtualbox/create/step05_centos.png)
点击 "完成" 按钮完成设置.
完成上述步骤后, VirtualBox的左侧列表中将添加一个虚拟机.

## 设置

- 选中新建的虚拟机后, 点击 "设置" 按钮.
![新建虚拟机 设置](/images/virtualbox/create/step06_ubuntu.png)

以下是一步步指南:

- 设置 -> 存储 -> 控制器: IDE, 确保已经分配光驱
![新建虚拟机 控制器](/images/virtualbox/create/step07_ubuntu.png)
![新建虚拟机 控制器](/images/virtualbox/create/step07_centos.png)

- 设置 -> 系统 -> 主板(M), 设置 "启动顺序(B)"
![新建虚拟机 启动顺序](/images/virtualbox/create/step08.png)

- 设置 -> 系统 -> 处理器(P)
![新建虚拟机 处理器](/images/virtualbox/create/step09.png)

- 设置 -> 网络 -> 网卡1 -> 高级(A), 点击 "端口转发(P)" 按钮
![新建虚拟机 端口转发](/images/virtualbox/create/step10.png)
在弹出 **端口转发规则** 窗口中, 点击右上方的 "+" 按钮
![新建虚拟机 端口转发规则](/images/virtualbox/create/step11.png)
配置 SSH
![新建虚拟机 端口转发SSH](/images/virtualbox/create/step12.png)

| 名称 | 协议 | 主机IP | 主机端口 | 子系统IP | 子系统端口 |
| ---- | ---- | ---- | ---- | ---- | ---- |
| SSH | TCP |   | 2022 |  | 22 |

- 设置 -> 网络 -> 网卡2,  勾选 "启用网络连接(E)", 连接方式选桥接网卡
![新建虚拟机 桥接网卡](/images/virtualbox/create/step13.png)

现在已经配置好了虚拟机, 可以启动啦.