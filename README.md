# Intro

This is a simple project to get introduced to Go using Redis as an in-memory cache / DB.

## Create a Redis Docker Instance

To run Redis, you can pull the official Redis image from Docker Hub and run it as a container.


```bash
docker pull redis

docker run --name redis -p 6379:6379 -d redis

```
