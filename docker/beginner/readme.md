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

### Run postgres with existing volume

```sh
docker run -d \
    --name postgres \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=password \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v postgres-data:/var/lib/postgresql/data/pgdata \
   --restart unless-stopped \
    -p 5432:5432 \
    postgres
```

## Docker Compose

### Create and Start Docker Compose + Rebuild Image

```sh
docker compose up -d --build
docker compose -f docker-compose.yml up -d --build
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
