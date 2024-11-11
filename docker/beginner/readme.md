## Discover Docker: The Beginner’s Journey

### Docker Container

#### Run a container of image `docker/welcome-to-docker` without exposing port, detach

```sh
docker run docker/welcome-to-docker
```

#### Run a container of image `docker/welcome-to-docker`

```sh
docker run -d -p 8080:80 --name welcome docker/welcome-to-docker
```

- `-d`: it will run the container in the background instead (alias `--detach`)
- `-p <HOST_PORT>:<CONTAINER_PORT>`: expose its ports. `-p 8080:80` Now,
  any traffic sent to port 8080 on your host machine will be forwarded to port 80 within the container.
- `--name <NAME>`: Custom Container Name
- `docker/welcome-to-docker`: Image name

### See Last 10 lines log and follow

```sh
docker logs welcome
docker logs -f --tail 10 welcome
```

- `-f`: follow
- `--tail n`: show latest n lines

### Execute command on the container

```sh
docker exec welcome cat /usr/share/nginx/html/index.html
```

### Execute an interactive sh shell on the container

```sh
docker exec -it welcome sh
```

### Stop the container welcome

```sh
docker stop welcome
```

## Building Image

### Build Dockerfile to Image

```sh
cd go_sample
docker build --tag docker-go-ping .
docker build -t docker-go-ping -f Dockerfile.alpine .

# check the image is exists
docker image ls
```

### Tag image

```sh
# tag latest docker-go-ping image
docker image tag docker-go-ping:latest docker-go-ping:v1.0
```

### Run image to container

```sh
docker run -d -p 8080:8080 --name my-docker-go-ping-container docker-go-ping
```

## Run Container with Persisting Data With Volumes (e.g., Database)

### Create volume

```sh
docker volume create postgres-data
```

### Run postgres with volume

```sh
# https://github.com/docker-library/docs/blob/master/postgres/README.md
docker run -d \
    --name postgres \
    --env POSTGRES_USER=postgres \
    --env POSTGRES_PASSWORD=password \
    --volume postgres-data:/var/lib/postgresql/data \
    --restart unless-stopped \
    --publish 5432:5432 \
    postgres

# If you're using git bash, add MSYS2_ENV_CONV_EXCL='*' MSYS2_ARG_CONV_EXCL='*' to avoid auto conversion from arguments and env
# https://www.msys2.org/docs/filesystem-paths/#environment-variables
MSYS2_ARG_CONV_EXCL='*' MSYS2_ENV_CONV_EXCL='*' docker run -d \
    --name postgres \
    --env POSTGRES_USER=postgres \
    --env POSTGRES_PASSWORD=password \
    --volume postgres-data:/var/lib/postgresql/data \
    --restart unless-stopped \
    --publish 5432:5432 \
    postgres
```

### Run mySQL

```sh
docker run --name some-mysql -v some-mysql-data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=password -d mysql
```

## Docker Compose

### Create and Start Docker Compose + Rebuild Image

```sh
docker compose up -d --build
docker compose -f docker-compose.yml up -d --build
docker compose -f docker-compose.alpine.yml up -d --build
```

### Stop compose containers

```sh
docker compose stop
docker compose -f docker-compose.yml stop
```

### Start stopped compose containers

```sh
docker compose start
docker compose -f docker-compose.yml start
```

### Stop and Destroy Docker Compose Containers

```sh
docker compose down
docker compose -f docker-compose.yml down -d --build
```

## Vue with Docker

```sh
npx create-vue vue_sample
cd vue_sample
npm install
code .
npm run dev


docker build -t vue_sample .
docker run -d -p 80:80 --name vue_sample vue_sample
```

## Other Things

Syntax ENTRYPOINT vs CMD: https://spacelift.io/blog/docker-entrypoint-vs-cmd

- ENTRYPOINT is the process that’s executed inside the container.
- CMD is the default set of arguments that are supplied to the ENTRYPOINT process.

There are also differences in how you override these values when you start a container:

- CMD is easily overridden by appending your own arguments to the docker run command.
- ENTRYPOINT can be changed using the --entrypoint flag, but this should rarely be necessary for container images being used in the way they were intended. If you do change the ENTRYPOINT, you’ll almost certainly need to set a custom CMD too—as otherwise, your new ENTRYPOINT is likely to receive arguments that it doesn’t understand.
