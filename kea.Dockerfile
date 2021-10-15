FROM ubuntu
RUN apt-get update && apt-get -y install curl apt-transport-https gnupg
RUN curl -1sLf 'https://dl.cloudsmith.io/public/isc/kea-1-8/cfg/setup/bash.deb.sh' | bash
RUN apt-get -y install isc-kea-common isc-kea-dhcp4-server isc-kea-ctrl-agent
ADD --chown=_kea:_kea ./test-data/kea/kea-dhcp4.conf /etc/kea/kea-dhcp4.conf
ADD --chown=_kea:_kea ./test-data/kea/kea-ctrl-agent.conf /etc/kea/kea-ctrl-agent.conf
USER _kea
ENV KEA_PIDFILE_DIR=/run/kea
ENV KEA_LOCKFILE_DIR=/run/lock/kea
ENV KEA_LOGGER_DESTINATION=/var/log/kea
EXPOSE 8000/tcp
COPY ./entrypoint.kea.sh /docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]
