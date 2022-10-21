FROM golang:latest

COPY id_rsa /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa \
 && ssh-keygen -y -f /root/.ssh/id_rsa > /root/.ssh/authorized_keys \
 && sed -i "s/root\:\!/root\:/g" /etc/shadow

RUN echo \
    deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye main contrib non-free\
    deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye main contrib non-free\
    deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-updates main contrib non-free\
    deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-updates main contrib non-free\
    deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-backports main contrib non-free\
    deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-backports main contrib non-free\
    deb https://mirrors.tuna.tsinghua.edu.cn/debian-security bullseye-security main contrib non-free\
    deb-src https://mirrors.tuna.tsinghua.edu.cn/debian-security bullseye-security main contrib non-free\
    > /etc/apt/sources.list


RUN apt-get update
RUN apt-get -y upgrade

RUN apt-get install -y vim
RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

WORKDIR /whitebench
