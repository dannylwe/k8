## Build Image
docker build -t k8 .

## Run Image
docker run -it --rm -p 9001:9001 k8

## Using Dockerhub
- Get started with dockerhub [HERE](https://docs.docker.com/docker-hub/).  
- login into dockerhub with `docker login`  
- Push `docker push acehanks/k8`  
- Pull `docker pull acehanks/k8:v1`  
- We use dockerhub to store our image in a public repo.  

## Run Kubernetes deployment
- kubectl apply -f server.yaml
