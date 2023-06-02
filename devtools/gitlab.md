# GitLab

- 参考 https://about.gitlab.com/install/#centos-7

```bash
sudo yum update -y
sudo yum upgrade -y

sudo yum install -y curl policycoreutils-python openssh-server perl

sudo systemctl enable sshd
sudo systemctl start sshd

curl https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.rpm.sh | sudo bash

sudo EXTERNAL_URL="http://gitlab.chengwanlu.com" yum install -y gitlab-ce
```

- 至少需要 4 GB 内存

- 修改 /etc/gitlab/gitlab.rb 文件

```text
# nginx['listen_port'] = nil

nginx['listen_port'] = 8000
```

- 配置更新

```bash
gitlab-ctl reconfigure
```

- 修改账户密码

```bash
# 进入gitlab串口环境
gitlab-rails console -e production 
```

输入: 

```
> user = User.where(id: 1).first 定位到gitlab 数据库中Users表中的一个用户，通常就是管理员用户admin@local.host
> user.password=12345678 重置管理员密码为12345678
> user.password_confirmation=12345678 确认管理员密码为12345678
> user.save! 保存更改信息,需要使用后面的感叹号!
```

- 修改 clone 地址
修改文件 /opt/gitlab/embedded/service/gitlab-rails/config/gitlab.yml

```text
## GitLab settings
gitlab:
    ## Web server settings (note: host is the FQDN, do not include http://)
    host: gitlab.chengwanlu.com
    port: 80
    https: false
```

- 重启

```bash
gitlab-ctl restart
```