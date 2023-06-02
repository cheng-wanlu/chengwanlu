# 防火墙

## Ubuntu

```bash
# 安装防火墙
sudo apt-get install ufw
# 启用防火墙
sudo ufw enable
# 默认的incoming策略更改为deny
sudo ufw default deny
# 允许所有的外部IP访问本机的22/tcp(ssh)端口
sudo ufw allow 22/tcp
# 查看状态
sudo ufw status
```