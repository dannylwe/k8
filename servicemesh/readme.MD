## Deploy application
kubectl apply -f deploy.yaml

# ngnix install
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.41.2/deploy/static/provider/cloud/deploy.yaml 

## remove ingress error
kubectl delete -A ValidatingWebhookConfiguration ingress-nginx-admission

## Run redis docker container
docker run --name redis-server -p 6379:6379 -d redis

## Build images and Push
docker build -t acehanks/getall:v3 . && docker push acehanks/getall  
docker build --no-cache -t acehanks/getbyid:v10 . && docker push acehanks/getbyid  
docker build --no-cache -t acehanks/getcust:v10 . && docker push acehanks/getcust  

## Inject Linkerd into deployments
kubectl get deployments -n servicemesh -o yaml | linkerd inject - | kubectl apply -f -

## TODO  

[ ] Add redis 
[ ] Add ngnix ingress
[ ] use linkerd 2.X service mesh  
[ ] use istio service mesh  