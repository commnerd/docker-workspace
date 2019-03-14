FROM commnerd/workspace

RUN sudo wget -O /etc/apt/trusted.gpg.d/php.gpg https://packages.sury.org/php/apt.gpg && \
    sudo sh -c 'echo "deb https://packages.sury.org/php/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/php.list' && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y \
	php7.2 \
        php7.2-common \
        php7.2-cli \
        php7.2-fpm \
        php7.2-sqlite \
        php7.2-mysql \
        php7.2-mbstring \
        php7.2-xml \
        php7.2-curl \
        php7.2-zip \
        phpunit

RUN curl -sL https://deb.nodesource.com/setup_11.x | bash -

RUN apt-get install -y nodejs

RUN npm i -g yarn n && n latest

RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" && \
    sudo php composer-setup.php --filename=composer --install-dir=/usr/bin && \
    php -r "unlink('composer-setup.php');"

EXPOSE 80 9000

CMD ["supervisord", "-nc", "/etc/supervisor/supervisord.conf"]
