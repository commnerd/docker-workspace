#!/bin/bash

DEST="${HOME}/.bash_profile"
TERMINAL_TOOL=""
NETWORK=""
PORTS="-p 80:80"
NAME="--name workspace"
VOLUMES="-v ${HOME}/Workspace:/home/commnerd/Workspace -v ${HOME}/.ssh:/root/.ssh -v /var/run/docker.sock:var/run/host.docker.sock"
IMAGE="commnerd/workspace:work"

case "$OSTYPE" in
"linux-gnu")
    PORTS=""
    NETWORK="--network host"
    ;;
"msys")
    VOLUMES="-v C:\\\\Users\\\\miller_m\\\\Workspace:/home/commnerd/Workspace -v C:\\\\Users\\\\miller_m\\\\.ssh:/root/.ssh -v //var/run/docker.sock:/var/run/host.docker.sock -v C:\\\\Users\\\\miller_m\\\\scripts:/home/commnerd/scripts"
    TERMINAL_TOOL="winpty "
    ;;
"win32")
    VOLUMES="-v C:\\\\Users\\\\miller_m\\\\Workspace:/home/commnerd/Workspace -v C:\\\\Users\\\\miller_m\\\\.ssh:/root/.ssh -v //var/run/docker.sock:/var/run/host.docker.sock -v C:\\\\Users\\\\miller_m\\\\scripts:/home/commnerd/scripts"
    TERMINAL_TOOL="winpty "
    ;;
esac

echo "alias run-workspace='docker run -d --privileged --restart always ${NAME} ${PORTS} ${VOLUMES} ${IMAGE}'" >> ${DEST}
echo "alias workspace='${TERMINAL_TOOL}docker exec -it -u commnerd --workdir //home/commnerd workspace bash'" >> ${DEST}
