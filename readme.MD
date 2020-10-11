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
- `cd /deployments`
- `kubectl apply -f ./server.yaml`

## To view endpoint
- Go to `localhost:30643`

## Create Loadbalancer with Kubectl
- `kubectl expose deployment go-deployment --port=7989 --target-port=9001 --name=go-service --type=LoadBalancer -o yaml > loadbalancer.yaml` 
