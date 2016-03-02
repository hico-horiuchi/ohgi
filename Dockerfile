FROM ubuntu:14.04

MAINTAINER hico-horiuchi <12ff5b8@gmail.com>

RUN sed -i "s/archive\.ubuntu\.com/ftp\.jaist\.ac\.jp/g" /etc/apt/sources.list && \
    apt-get update && \
    apt-get -y upgrade && \
    apt-get -y install wget

RUN wget http://packages.erlang-solutions.com/erlang-solutions_1.0_all.deb && \
    dpkg -i erlang-solutions_1.0_all.deb && \
    apt-get update && \
    apt-get -y install erlang-nox=1:18.2

RUN wget http://www.rabbitmq.com/releases/rabbitmq-server/v3.6.0/rabbitmq-server_3.6.0-1_all.deb && \
    dpkg -i rabbitmq-server_3.6.0-1_all.deb

RUN apt-get -y install redis-server

RUN wget -q http://repositories.sensuapp.org/apt/pubkey.gpg -O- | apt-key add - && \
    echo "deb     http://repositories.sensuapp.org/apt sensu main" | tee /etc/apt/sources.list.d/sensu.list && \
    apt-get update && \
    apt-get -y install sensu && \
    wget -O /etc/sensu/config.json http://sensuapp.org/docs/latest/files/config.json && \
    wget -O /etc/sensu/conf.d/default_handler.json http://sensuapp.org/docs/latest/files/default_handler.json && \
    wget -O /etc/sensu/conf.d/client.json http://sensuapp.org/docs/latest/files/client.json && \
    echo "{\n  \"checks\": {\n    \"default\": {\n      \"command\": \"/etc/sensu/plugins/default.sh\",\n      \"subscribers\": [\n        \"test\"\n      ],\n      \"interval\": 10,\n      \"handler\": \"default\",\n      \"aggregate\": true\n    }\n  }\n}" | tee /etc/sensu/conf.d/default_check.json && \
    echo "#!/bin/sh\n\necho \"Default WARNING\"\nexit 1" | tee /etc/sensu/plugins/default.sh && \
    chmod +x /etc/sensu/plugins/default.sh && \
    sudo chown -R sensu:sensu /etc/sensu

RUN apt-get -y autoremove && \
    apt-get clean && \
    rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/* /tmp/* /var/tmp/*

EXPOSE 4567

CMD service rabbitmq-server restart && \
    rabbitmqctl add_vhost /sensu && \
    rabbitmqctl add_user sensu secret && \
    rabbitmqctl set_permissions -p /sensu sensu ".*" ".*" ".*" && \
    service redis-server restart && \
    service sensu-server restart && \
    service sensu-api restart && \
    service sensu-client restart && \
    tailf /var/log/sensu/*.log
