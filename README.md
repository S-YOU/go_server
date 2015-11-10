# Go Server

Simple server implementation by golang

## Run in Docker

First, install [docker toolbox](https://www.docker.com/docker-toolbox)

Then, type the commands below to run docker

```
$ eval $(docker-machine env default)
$ docker build -t docker_go .
$ docker run -d -p 80:8080 --name go_server docker_go
```

From now, you can access `docker-machine ip default` in browser

You can remove docker go_server by

```
$ docker rm -f go_server
```
