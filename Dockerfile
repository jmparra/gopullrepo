FROM debian:latest

MAINTAINER Antarkt Inc.
LABEL vendor=antarkt.inc
LABEL inc.antarkt.santopanel.role=script

ENV PACKAGES "git-core ca-certificates openssh-client ssh bash"
ENV minimal_apt_get_args '-y -qq --no-install-recommends'

ADD puller /usr/local/bin/puller
RUN apt-get update -qq && apt-get install $minimal_apt_get_args $PACKAGES && \
    chmod +x /usr/local/bin/puller && \
    mkdir -p /root/.ssh/ && \
    echo "    IdentityFile /root/.ssh/id_rsa" >> /etc/ssh/ssh_config && \
    ssh-keyscan -p22 bitbucket.org > /root/.ssh/known_hosts

ENTRYPOINT ["puller"]
