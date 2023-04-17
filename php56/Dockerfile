FROM php:7.4.33-fpm-buster

USER root
## 更换源
RUN sed -i "s/deb.debian.org/mirrors.aliyun.com/g" /etc/apt/sources.list
RUN sed -i 's/listen = 127.0.0.1:9000/listen = 0.0.0.0:9000/g' /usr/local/etc/php-fpm.d/www.conf

RUN apt-get update && apt-get install -y \
    libfreetype6-dev \
    libjpeg62-turbo-dev \
    libmcrypt-dev \
    libpng-dev \
    libmemcached-dev \
    zlib1g-dev \
    openssh-server \
    vim \
    git \
    wget \
    libzip-dev \
    unzip \
    zip \
    && docker-php-ext-install -j$(nproc) iconv \
    && docker-php-ext-configure gd --with-freetype --with-jpeg \
    && docker-php-ext-install -j$(nproc) gd \
    && pecl install mcrypt-1.0.4 \
    && docker-php-ext-enable mcrypt \
    && docker-php-ext-install mysqli \
    && docker-php-ext-install pdo pdo_mysql \
    && docker-php-ext-install zip \
    && pecl install redis-5.3.2 \
    && docker-php-ext-enable redis \
    && docker-php-ext-install -j$(nproc) bcmath \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY id_rsa /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa \
 && ssh-keygen -y -f /root/.ssh/id_rsa > /root/.ssh/authorized_keys \
 && sed -i "s/root\:\!/root\:/g" /etc/shadow

# 安装 composer
RUN wget -O installer https://getcomposer.org/installer \
 && php installer \
 && mv composer.phar /usr/local/bin/composer \
 && composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/ \
 && rm installer

# 配置 nodejs 可以自己换源
RUN wget https://github.com/whitexiong/whitebench/releases/download/nodejs/node-v18.14.0-linux-x64.tar.gz -O /home/node-v18.14.0-linux-x64.tar.gz\
    && tar zxvf /home/node-v18.14.0-linux-x64.tar.gz -C /home \
    && rm -rf /home/node-v18.14.0-linux-x64.tar.gz \
    && mv /home/node-v18.14.0-linux-x64 /usr/local/node18 \
    && chmod 755 /usr/local/node18 \
    && chmod -R 755 /usr/local/node18/* \
    && chown root:root /usr/local/node18 \
    && chown -R root:root /usr/local/node18/*

RUN echo \
    export NODE_HOME=/usr/local/node18 >> /etc/profile\
    && echo \
    export PATH=\$PATH:\$NODE_HOME/bin >> /etc/profile

# 安装 node
RUN echo \
        source /etc/profile\
    >> ~/.bashrc\
    && /bin/bash -c "source ~/.bashrc"

WORKDIR /whitebench