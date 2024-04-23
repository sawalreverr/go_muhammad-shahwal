## Setup

Command in your terminal

```bash
# build your own project
docker build -t go-simple-api:1.0.0 .

# test run
docker run -itd --name simple-api -p 8080:8080 go-simple-api:1.0.0

# check log
docker logs simple-api

# tagging your docker images
docker tag go-simple-api:1.0.0 sawalrever23/go-simple-api:1.0.0

# push your images
docker push sawalrever23/go-simple-api:1.0.0
```

Command in AWS EC2 Console

```bash
# docker install
sudo yum install docker -y

# check version
sudo docker version

# start docker service
sudo systemctl start docker

# pull your images
sudo docker pull sawalrever23/go-simple-api:1.0.0

# run your images
sudo docker run -d -p 8080:8080 -it --rm --name simple-api sawalrever23/go-simple-api:1.0.0

# check your container logs
sudo docker logs simple-api

# stop your container
sudo docker stop simple-api # it will be delete your container too, bcz we define --rm
```

dont forget to rename .env.example to .env and define your private .env
