# 安装 Docker

按照 docker 官方文档一步步安装

```bash
# 更新包管理器
sudo apt update
# 卸载旧版 docker
sudo apt-get remove docker docker-engine docker.io containerd runc
# 安装所需的依赖包
sudo apt-get install ca-certificates curl gnupg
# 创建目录用于存放密钥
sudo install -m 0755 -d /etc/apt/keyrings
# 下载 docker 的 GPG 密钥并保存到指定路径
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
# 设置密钥文件的权限
sudo chmod a+r /etc/apt/keyrings/docker.gpg
# 添加 docker 的软件源到 apt 源列表
echo \
  "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
# 更新包列表
sudo apt-get update
# 安装 docker
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```