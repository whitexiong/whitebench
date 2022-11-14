FROM php:8.0-fpm-buster

## 更换源
RUN sed -i "s/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/g" /etc/apt/sources.list

RUN apt-get update && apt-get install -y \
    libfreetype6-dev libjpeg62-turbo-dev libmcrypt-dev libpng-dev \
    && apt-get install -y libmemcached-dev zlib1g-dev \
    && apt-get install -y openssh-server \
    && apt-get install -y vim \
    && apt-get install -y git \
    && apt-get install -y wget \
    && apt-get install -y libzip-dev \
    && apt-get install -y unzip zip  \
    && docker-php-ext-install -j$(nproc) iconv \
    && docker-php-ext-configure gd --with-freetype --with-jpeg \
    && docker-php-ext-install -j$(nproc) gd \
    # PHP8必须装1.0.4  参考 https://bugs.php.net/bug.php?id=80459
    && pecl install mcrypt-1.0.4 \
    && docker-php-ext-enable mcrypt \
    && docker-php-ext-install mysqli \
    && docker-php-ext-install pdo pdo_mysql \
    && docker-php-ext-install zip \
    && pecl install redis-5.3.2 \
    && docker-php-ext-enable redis \
    && docker-php-ext-install -j$(nproc) bcmath

COPY id_rsa /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa \
 && ssh-keygen -y -f /root/.ssh/id_rsa > /root/.ssh/authorized_keys \
 && sed -i "s/root\:\!/root\:/g" /etc/shadow

# 安装 composer \
RUN php -r "copy('https://install.phpcomposer.com/installer', 'composer-setup.php');" \
    && php composer-setup.php \
    && php -r "unlink('composer-setup.php');" \
    && mv composer.phar /usr/local/bin/composer \
    && composer selfupdate \
    && composer config -g repo.packagist composer https://packagist.phpcomposer.com

# 配置 nodejs 可以自己换源
RUN wget https://github.com/whitexiong/whitebench/releases/download/nodejs/node-v16.18.0-linux-x64.tar.gz -O /home/node-v16.18.0-linux-x64.tar.gz\
    && tar zxvf /home/node-v16.18.0-linux-x64.tar.gz -C /home \
    && rm -rf /home/node-v16.18.0-linux-x64.tar.gz \
    && mv /home/node-v16.18.0-linux-x64 /usr/local/node16 \
    && chmod 755 /usr/local/node16 \
    && chmod -R 755 /usr/local/node16/* \
    && chown root:root /usr/local/node16 \
    && chown -R root:root /usr/local/node16/*

WORKDIR /whitebench