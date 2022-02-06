# rssmq
Checks RSS feeds for new content on a regular schedule, and submits new items to several customizable endpoints.

# Deployment
## Bare Metal
1. Download the binary `rssmq` and place it in your path.
2. Create a config file at `/opt/rssmq.json` (or pass in a custom path with `--config`)
3. Run `rssmq` on the command line.
## Docker
1. Create config file `rssmq.json`.
2. Run `docker run -v <config_path>:/opt/rssmq.json -d ghcr.io/armmaster17/rssmq:latest`
## Kubernetes Manifest
```bash
apiVersion: apps/v1
kind: Deployment
metadata:
  name: rssmq-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: rssmq
  template:
    metadata:
      labels:
        app: rssmq
    spec:
      containers:
      - name: rssmq
        image: ghcr.io/armmaster17/rssmq:latest
        imagePullPolicy: Always
        resources:
          limits:
            cpu: "25m"
            memory: "64Mi"
```
