ARG MYSQL_VERSION=${MYSQL_VERSION}

FROM mysql:${MYSQL_VERSION}

# Set debian default locale to ja_JP.UTF-8
RUN apt-get update && \
    apt-get install -y locales && \
    rm -rf /var/lib/apt/lists/* && \
    echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen && \
    locale-gen ja_JP.UTF-8
ENV LC_ALL ja_JP.UTF-8

# Set MySQL character
COPY my.cnf /etc/mysql/conf.d/charset.cnf

EXPOSE 3306
