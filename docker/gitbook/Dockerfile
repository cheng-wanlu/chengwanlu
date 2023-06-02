# 使用Node.js镜像作为基础镜像
FROM node:10

# 设置工作目录
WORKDIR /www

# 安装依赖
RUN npm install gitbook-cli -g

# 安装插件然后生成静态网站
CMD ["bash", "-c", "gitbook install && gitbook build"]