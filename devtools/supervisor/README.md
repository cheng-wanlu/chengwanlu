# Supervisor

## 安装
```bash
# 安装 python
$ sudo yum install python3 python3-devel
# 安装 supervisor
$ sudo pip3 install supervisor
```

## 配置

[import, title:"/data/supervisord/supervisord.conf", lang:"text"](supervisord.conf)

## 启动

```bash
# 启动 supervisor
$ supervisord -c /data/supervisord/supervisord.conf
```

### 代理服务器

[import, title:"/data/supervisord/confs/daze.conf", lang:"text"](daze.conf)

```bash
# 代理服务器
$ supervisorctl update
$ supervisorctl start daze
```

### etcd

[import, title:"/data/supervisord/confs/etcd.conf", lang:"text"](etcd.conf)

```bash
# etcd
$ supervisorctl update
$ supervisorctl start etcd
```