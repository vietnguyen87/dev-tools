FROM redis:latest

MAINTAINER Viet Nguyen Thanh <vietnt@sendo.vn>

RUN apt-get update
RUN apt-get install -y apt-utils
RUN apt-get install -y gettext-base

COPY redis.conf /etc/redis/redis.conf.template
COPY start.sh /start.sh

ENTRYPOINT [ "/bin/bash", "/start.sh" ]