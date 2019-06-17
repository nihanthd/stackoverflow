Install docker-machine from https://docs.docker.com/machine/install-machine/

Create docker machine

``$ docker-machine create --driver virtualbox dev``

Now import the env variables of the ***dev*** docker-machine by running the following command

```$ eval "$(docker-machine env dev)"```

Execute build.go

```$ go run build.go```