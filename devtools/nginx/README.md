- CentOS

```bash
$ sudo yum install yum-utils
```

[import, title:"/etc/yum.repos.d/nginx.repo", lang:"text"](nginx.repo)

[import, title:"/etc/nginx/nginx.conf", lang:"text"](nginx.conf)

```bash
$ sudo yum-config-manager --enable nginx-stable

$ sudo yum install nginx

$ nginx -c /etc/nginx/nginx.conf
```

[import, title:"/data/nginx/conf/gitlab.conf", lang:"text"](gitlab.conf)

