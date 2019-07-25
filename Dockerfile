FROM centos:latest 
MAINTAINER Ashmeet Singh <ashmeet.s.arora@gmail.com>

RUN yum install -y epel-release

RUN yum install -y gcc  wget git vim  which && \
 mkdir -p /logs && mkdir -p /etc/geolocation/settings && \
 yum clean all

 RUN yum install -y supervisor make

 RUN curl -O https://storage.googleapis.com/golang/go1.11.2.linux-amd64.tar.gz && \
    tar -xvf go1.11.2.linux-amd64.tar.gz -C /usr/local

COPY dockerconfig/supervisord /etc/rc.d/init.d/

COPY ./dockerconfig/process.conf /etc/supervisord.d/
COPY ./settings/config.json /etc/geolocation/settings

ENV  GOROOT /usr/local/go

ENV  PATH ${PATH}:/usr/local/go/bin

RUN chmod 755 /etc/rc.d/init.d/supervisord