# 密钥登录

- 以 root 用户身份运行下面的命令.

```bash
# 创建目录
mkdir -p /root/.ssh 
# 更改目录权限
chmod 700 /root/.ssh 
# 写入公钥
vim /root/.ssh/authorized_keys 
# 修改公钥文件权限
chmod 600 /root/.ssh/authorized_keys
```

- 以 wanlu 用户身份运行下面的命令.

```bash
# 创建文件夹
mkdir -p .ssh    
# 修改文件夹权限为拥有者可'读,写,执行'
chmod 700 .ssh    
# 复制公钥
sudo cat /root/.ssh/authorized_keys > .ssh/authorized_keys
# 修改公钥文件权限
chmod 600 .ssh/authorized_keys
```


- 修改配置文件: /etc/ssh/sshd_config

```text
RSAAuthentication yes # 开启密钥登入的认证方式
PubkeyAuthentication yes # 开启密钥登入的认证方式
PermitRootLogin yes #此处请留意 root 用户能否通过 SSH 登录，默认为yes：
PasswordAuthentication yes #当我们完成全部设置并以密钥方式登录成功后，可以禁用密码登录。这里我们先不禁用，
```

- 重启 SSH 服务

```
$ sudo service sshd restart
$ sudo systemctl restart sshd
```