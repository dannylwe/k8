## Build Image
docker build -t k8 .

## Run Image
docker run -it --rm -p 9001:9001 k8

## Using Dockerhub
- Get started with dockerhub [HERE](https://docs.docker.com/docker-hub/).  
- We use dockerhub to store our image in a public repo.