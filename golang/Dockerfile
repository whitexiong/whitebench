FROM golang:latest

COPY id_rsa /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa \
 && ssh-keygen -y -f /root/.ssh/id_rsa > /root/.ssh/authorized_keys \
 && sed -i "s/root\:\!/root\:/g" /etc/shadow

## 其他方式替换源
#RUN echo \
#    deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye main contrib non-free\
#    deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye main contrib non-free\
#    deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-updates main contrib non-free\
#    deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-updates main contrib non-free\
#    deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-backports main contrib non-free\
#    deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-backports main contrib non-free\
#    deb https://mirrors.tuna.tsinghua.edu.cn/debian-security bullseye-security main contrib non-free\
#    deb-src https://mirrors.tuna.tsinghua.edu.cn/debian-security bullseye-security main contrib non-free\
#    > /etc/apt/sources.list

RUN sed -i "s/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/g" /etc/apt/sources.list

# 配置 nodejs 可以自己换源
RUN wget https://github.com/whitexiong/whitebench/releases/download/nodejs/node-v16.18.0-linux-x64.tar.gz -O /home/node-v16.18.0-linux-x64.tar.gz\
    && tar zxvf /home/node-v16.18.0-linux-x64.tar.gz -C /home \
    && rm -rf /home/node-v16.18.0-linux-x64.tar.gz \
    && mv /home/node-v16.18.0-linux-x64 /usr/local/node16 \
    && chmod 755 /usr/local/node16 \
    && chmod -R 755 /usr/local/node16/* \
    && chown root:root /usr/local/node16 \
    && chown -R root:root /usr/local/node16/*

# 下载 1.8 版本
RUN wget https://github.com/whitexiong/whitebench/releases/download/nodejs/go1.18.8.linux-amd64.tar.gz -O /home/go1.18.8.linux-amd64.tar.gz\
    && tar zxvf /home/go1.18.8.linux-amd64.tar.gz -C /home \
    && rm -rf /home/go1.18.8.linux-amd64.tar.gz \
    && mv /home/go /usr/local/go1.8 \
    && chmod 755 /usr/local/go1.8 \
    && chmod -R 755 /usr/local/go1.8/* \
    && chown root:root /usr/local/go1.8 \
    && chown -R root:root /usr/local/go1.8/*

# 环境变量
RUN apt-get update \
    && apt-get -y upgrade \
    && apt-get install -y vim \
    && apt-get install -y locales \
    && go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct

RUN echo export GO111MODULE=on >> /etc/profile\
    && echo \
    export GOPROXY=https://goproxy.cn >> /etc/profile\
    && echo \
    export GOROOT=/usr/local/go1.8 >> /etc/profile\
    && echo \
    export NODE_HOME=/usr/local/node16 >> /etc/profile\
    && echo \
    export PATH=\$PATH:\$NODE_HOME/bin:\$PATH:\$GOROOT/bin:\$GOPATH/bin >> /etc/profile

#安装 node
RUN echo \
        source /etc/profile\
    >> ~/.bashrc\
    && /bin/bash -c "source ~/.bashrc"

# 配置 vimrc
RUN curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

WORKDIR /whitebench