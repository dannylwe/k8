## Build Image
docker build -t k8 .

## Run Image
docker run -it --rm -p 9001:9001 k8