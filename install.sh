#!/bin/bash

DEST="${HOME}/.docker-workspace"
TERMINAL_TOOL=""
NETWORK=""
PORTS=" -p 80:80"
NAME=" --name workspace"
VERSION="latest"
if [ "$1" ]; then VERSION="$1"; fi
VOLUMES=" -v ${HOME}/Workspace:/home/commnerd/Workspace -v ${HOME}/.ssh:/root/.ssh -v /var/run/docker.sock:/var/run/host.docker.sock"
IMAGE=" commnerd/workspace:${VERSION}"

install_to_dest() {
    if [ ! "$(grep ${DEST} ${1})" ]
    then
        echo ". ${DEST}" >> ${1}
        exit
    fi
}

case "$OSTYPE" in
"linux-gnu")
    PORTS=""
    NETWORK=" --network host"
    ;;
"msys")
    VOLUMES=" -v C:\\\\Users\\\\miller_m\\\\Workspace:/home/commnerd/Workspace -v C:\\\\Users\\\\miller_m\\\\.ssh:/root/.ssh -v //var/run/docker.sock:/var/run/host.docker.sock -v C:\\\\Users\\\\miller_m\\\\scripts:/home/commnerd/scripts"
    TERMINAL_TOOL="winpty "
    ;;
"win32")
    VOLUMES=" -v C:\\\\Users\\\\miller_m\\\\Workspace:/home/commnerd/Workspace -v C:\\\\Users\\\\miller_m\\\\.ssh:/root/.ssh -v //var/run/docker.sock:/var/run/host.docker.sock -v C:\\\\Users\\\\miller_m\\\\scripts:/home/commnerd/scripts"
    TERMINAL_TOOL="winpty "
    ;;
esac

echo "alias run-workspace='docker run -d --privileged --restart always ${NETWORK}${NAME}${PORTS}${VOLUMES}${IMAGE}'" > ${DEST}
echo "alias workspace='${TERMINAL_TOOL}docker exec -it -u commnerd --workdir //home/commnerd workspace bash'" >> ${DEST}

if [ -f ${HOME}/.bash_profile ]
then
    install_to_dest ${HOME}/.bash_profile
fi

if [ -f ${HOME}/.bashrc ]
then
    install_to_dest ${HOME}/.bashrc
fi
