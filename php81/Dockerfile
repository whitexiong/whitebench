FROM php:8.1.0-fpm-buster

RUN sed -i "s/deb.debian.org/mirrors.aliyun.com/g" /etc/apt/sources.list

RUN apt-get update && apt-get install -y \
    libfreetype6-dev \
    libjpeg62-turbo-dev \
    libpng-dev \
    libxml2-dev \
    openssh-server \
    vim \
    git \
    wget \
    && apt-get install -y unzip zip \
    && docker-php-ext-install -j$(nproc) iconv \
    && docker-php-ext-configure gd --with-freetype --with-jpeg \
    && docker-php-ext-install -j$(nproc) gd \
    && docker-php-ext-install pdo_mysql \
    && docker-php-ext-install soap \
    && docker-php-ext-install xml \
    && pecl install swoole \
    && pecl install mongodb \
    && docker-php-ext-enable mongodb

COPY id_rsa /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa \
 && ssh-keygen -y -f /root/.ssh/id_rsa > /root/.ssh/authorized_keys \
 && sed -i "s/root\:\!/root\:/g" /etc/shadow

# 安装 composer \
RUN wget https://getcomposer.org/installer
RUN php installer
RUN mv composer.phar /usr/local/bin/composer
RUN composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/

# 配置 nodejs 可以自己换源
RUN wget https://github.com/whitexiong/whitebench/releases/download/nodejs/node-v18.14.0-linux-x64.tar.gz -O /home/node-v18.14.0-linux-x64.tar.gz\
    && tar zxvf /home/node-v18.14.0-linux-x64.tar.gz -C /home \
    && rm -rf /home/node-v18.14.0-linux-x64.tar.gz \
    && mv /home/node-v18.14.0-linux-x64 /usr/local/node18 \
    && chmod 755 /usr/local/node18 \
    && chmod -R 755 /usr/local/node18/* \
    && chown root:root /usr/local/node18 \
    && chown -R root:root /usr/local/node18/*

WORKDIR /whitebench