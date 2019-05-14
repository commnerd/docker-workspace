# docker-workspace
A consistent Debian-based workspace setup for me using docker

## Run the following to bring up the workspace environment
`docker run -d --restart always --privileged -v ~/Workspace:/home/commnerd/workspace -v ~/.ssh:/root/.ssh -p 80:80 -p 9000:9000 --name workspace commnerd/workspace`

or

`./run.sh`

The 'run.sh' command calls the docker command above as-is.  If you need customization, please craft the 'docker run' command yourself.


## Build the environment by running
`docker build --squash -t commnerd/workspace .`

or 

`./build.sh`


## It is suggested you add the following command to the end of your .bashrc file to connect to the container and change into the home directory

`docker exec -it --workdir /home/commnerd -u commnerd workspace bash`

or

`./connect.sh`

## There is an installation script that installs the following commands:

`run-workspace`: Kicks off the docker run process for the workspace
`workspace`: Connects to the workspace from the host
