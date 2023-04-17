# 使用官方 Nginx 镜像作为基础镜像
FROM nginx:latest

# 设置工作目录
WORKDIR /usr/share/nginx/html

# 移除默认的 Nginx 配置文件
RUN rm /etc/nginx/conf.d/default.conf

# 复制我们自定义的 Nginx 配置文件
COPY nginx.conf /etc/nginx/conf.d