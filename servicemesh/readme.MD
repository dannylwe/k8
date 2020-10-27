## Run redis docker container
docker run --name redis-server -p 6379:6379 -d redis

## Build images and Push
docker build -t acehanks/getall:v2 . && docker push acehanks/getall  
docker build -t acehanks/getbyid:v2 . && docker push acehanks/getbyid  
docker build -t acehanks/getcust:v2 . && docker push acehanks/getcust  

## Inject Linkerd into deployments
kubectl get deployments -n servicemesh -o yaml | linkerd inject - | kubectl apply -f -

## TODO  

[ ] Deploy to k8  
[ ] Add ngnix ingress
[ ] use linkerd 2.X service mesh  
[ ] use istio service mesh  