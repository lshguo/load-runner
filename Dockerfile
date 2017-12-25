FROM index.alauda.cn/library/centos:7

MAINTAINER "lshguo@github.com"

COPY load-runner /usr/local/bin/load-runner

EXPOSE 8080

CMD /usr/local/bin/load-runner
