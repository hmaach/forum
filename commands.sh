#!/bin/bash

# Stop and remove existing container if it exists
docker stop forum-con 
docker rm forum-con

# Remove unused data
docker system prune -f

# Remove the old image
docker rmi forum-img 

# Build a new image
docker build --no-cache -f dockerfile -t forum-img .

# Run a new container
docker run -d -p 8080:8080 --name forum-con forum-img

docker logs forum-con
