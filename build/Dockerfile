FROM riet/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/aliyun-monitor-operator \
    USER_UID=1001 \
    USER_NAME=aliyun-monitor-operator

# install operator binary
COPY build/_output/bin/aliyun-monitor-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
