远程配置
VSCode 的 Remote 功能由三个插件组成, 分别实现三种不同场景的远程开发.

- Remote - SSH: 利用 SSH 连接远程主机进行开发.
- Remote - Container: 连接当前机器上的容器进行开发.
- Remote - WSL: 在Windows 10上, 连接子系统(Windows Subsystem for Linux)进行开发.

首先我们这里在本地环境 Mac 上安装上 VSCode, 远程开发的机器 IP 为 192.168.1.123, 配置该节点可以本地通过 SSH 远程连接. 然后在 VSCode 中安装 Remote SSH 插件:
![VSCode 安装扩展](/images/devtools/vscode/extension.png)
安装了 Remote - SSH 扩展后, 你会在最左边看到一个新的状态栏图标.
![VSCode 安装扩展](/images/devtools/vscode/statusbar.png)
远程状态栏图标可以快速显示 VS Code 在哪个上下文中运行(本地或远程), 点击该图标或者点击 F1 按键然后输入remote-ssh 便会弹出 Remote-SSH 的相关命令.
![VSCode 安装扩展](/images/devtools/vscode/connect.png)
选择 "Remote-SSH: Connect to Host" 命令, 然后按以下格式输入远程主机的连接信息, 连接到主机: user@hostname, 然后根据提示输入登录的密码.
VSCode 将打开一个新窗口, 然后你会看到 "VSCode 服务器" 正在 SSH 主机上初始化的通知, 一旦 VSCode 服务器安装在远程主机上, 它就可以运行扩展并与你的本地 VSCode 实例通信了. 通过查看状态栏中的指示器, 可以知道已连接到虚拟机了, 它显示的是你的虚拟机的主机名.
Remote-SSH 扩展还在你的活动栏上添加了一个新的图标, 点击它将打开远程浏览器. 从下拉菜单中, 可以选择 SSH 目标, 在这里你可以配置你的 SSH 连接.
一旦你连接到你的 SSH 主机, 你就可以与远程机器上的文件进行交互了, 如果你打开集成终端(⌃`), 你会发现现在我们是在远程的 Linux 下面了.
现在我们可以使用 bash shell 浏览远程主机上的文件系统, 还可以使用 "文件">"打开文件夹" 浏览和打开远程主目录上的文件夹.

[使用 VSCode 远程开发调试](https://www.qikqiak.com/post/use-vscode-remote-dev-debug/)