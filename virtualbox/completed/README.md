# 查看系统版本
```bash
$ cat /etc/os-release
```

# 设置系统sudo不需要密码
```bash
$ sudo gedit /etc/sudoers
```

修改这2个位置, 增加NOPASSWD:, 注意NOPASSWD:的冒号, 且前后有空格, 保存, sudo就不需要密码了:
![安装完成 sudoers](/images/virtualbox/completed/sudoers.png)

```bash
# 为用户username添加sudo权限
$ sudo usermod -a -G sudo username
 
# 去除用户username的sudo权限
$ sudo usermod -G usergroup username

# 安装 git、vim、curl 等常用工具
$ sudo apt update
$ sudo apt install -y git vim curl jq

# 安装 openssh-server 来支持远程登录
$ sudo apt install -y openssh-server

# 安装Docker
# 查看系统中是否已经安装Docker
$ docker --version
# 使用如下命令安装Docker的最新版本
$ sudo apt update
$ sudo apt install -y docker.io
# 重启docker
$ sudo systemctl start docker
# 设置系统启动时docker启动, 可选
$ sudo systemctl enable docker

# 安装docker-compose
# 确定系统中是否已安装docker-compose工具
$ docker-compose --version
# 如果系统提示未安装, 则使用如下命令安装docker-compose工具
$ sudo apt install -y docker-compose
```
配置镜像源, 国内可用的几个镜像源:
- Docker 官方中国区: registry.docker-cn.com
- 网易: hub-mirror.c.163.com
- 中国科技大学: docker.mirrors.ustc.edu.cn
- 阿里云: y0qd3iq.mirror.aliyuncs.com
修改 Docker 的镜像源配置文件 /etc/docker/daemon.json, 如果没有配置过镜像该文件默认是不存的, 在其中增加如下内容:
```bash
$ sudo vim /etc/docker/daemon.json
#把以下代码加进去
{
"registry-mirrors":["https://registry.docker-cn.com"]
}
```

重启Docker服务
```bash
# 重启
$ sudo systemctl restart docker
# 命令查看配置是否生效：
$ docker info|grep Mirrors -A 1
```