FROM alpine:latest

WORKDIR /uif
RUN apk add curl tar bash nftables iproute2
COPY linux_install.sh linux_install.sh
RUN chmod +x linux_install.sh
RUN echo "true" > "./use_docker"

RUN bash linux_install.sh

# COPY docker_run.sh /usr/bin/uif/docker_run.sh
# RUN chmod +x /usr/bin/uif/docker_run.sh
# CMD ["bash", "/usr/bin/uif/docker_run.sh"]

# 设置容器启动时执行的命令
CMD ["/usr/bin/uif/uif"]
