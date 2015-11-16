# Go Server

[![Docker Repository on Quay](https://quay.io/repository/awakia/go_server/status?token=9ec28669-a67d-4153-b3e2-90342be83c81 "Docker Repository on Quay")](https://quay.io/repository/awakia/go_server)

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

```
$ open http://$(docker-machine ip default)
```

You can remove docker go_server by

```
$ docker rm -f go_server
```
