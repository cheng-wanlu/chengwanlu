# 用户授权

## 把新用户添加到 Sudoers 列表中：

- Ubuntu 

```bash
# 以 root 或者其他 sudo 用户身份运行下面的命令.
usermod -aG sudo wanlu
# 用户 "wanlu" 无需输入密码
echo "wanlu ALL=(ALL) NOPASSWD:ALL" | tee /etc/sudoers.d/wanlu
```

- CentOS

```bash
# 以 root 或者其他 sudo 用户身份运行下面的命令.
groups wanlu
usermod -aG wheel wanlu
sudo whoami
# 用户 "wanlu" 无需输入密码
echo "wanlu ALL=(ALL) NOPASSWD:ALL" | tee /etc/sudoers.d/wanlu
```