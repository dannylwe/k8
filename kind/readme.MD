## How to Setup Kubernetes Loadbalancer on Baremetal

NOTE: [docker](https://docs.docker.com/engine/install/ubuntu/) is required to be installed on the host.

### Installation on Linux
We are going to be using kindD to create our local kubernetes cluster. You can get kind

`wget https://github.com/kubernetes-sigs/kind/releases/download/v0.9.0/kind-linux-amd64`.

For more information on kind, visit: https://kind.sigs.k8s.io/  

Create executable with `chmod +x kind-linux-amd64`  

Run: `sudo mv kind-linux-amd64 /usr/local/bin/kind`  

Confirm installation of kind, `which kind` and `kind version`. This will list kind version information.

### Installation on MacOS
`brew install kind`

## Create cluster
`kind create cluster`  

NOTE: If you have go (1.11+) and docker installed GO111MODULE="on" `go get sigs.k8s.io/kind@v0.9.0 && kind create cluster` is all you need!  

confirm cluster is running `kubectl cluster-info`  