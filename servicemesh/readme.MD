## Run redis docker container
docker run --name redis-server -p 6379:6379 -d redis

## Build images
docker build -t acehanks/getall:v1 .
docker build -t acehanks/getbyid:v1 .
docker build -t acehanks/getcust:v1 .

## Push images to Dockerhub
docker push acehanks/getall
docker push acehanks/getbyid
docker push acehanks/getcust

## TODO  
[ ] Add k8 yaml  
[ ] Deploy to k8  
[ ] use linkerd 2.X service mesh  
[ ] use istio service mesh  