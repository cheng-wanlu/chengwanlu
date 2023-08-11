# nginx 镜像

[import, lang:"text"](Dockerfile)

[import, lang:"text"](nginx.conf)

```bash
# 创建镜像
sudo docker build -t nginx-custom docker/nginx
# 运行容器
sudo docker run -d -p 80:80 -v "$PWD:/www" --name=chengwanlu-com-serve nginx-custom
```
