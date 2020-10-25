## Download promtail binary

## Launch
```
curl https://raw.githubusercontent.com/grafana/loki/v1.6.0/production/docker-compose.yaml -O docker-compose.yaml
```
`docker-compose up`  

This will launch grafana on port 3000 and loki on 3100. default creds {admin}:{admin}  

Promtail server running on port 80. Can be found at "/".

## Send Logs
./promtail -config.file=config.yaml
