# rssmq
Checks RSS feeds for new content on a regular schedule, and submits new items to a RabbitMQ queue.

# Deployment
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
        env:
        - name: RSSMQ_DB_HOST
          value: <db-ip>
        - name: RSSMQ_DB_USER
          value: <db-user>
        - name: RSSMQ_DB_PASSWORD
          value: <db-password>
        - name: RSSMQ_DB_DATABASE
          value: rssmq
        - name: RSSMQ_MQ_URL
          value: "amqp://rssmq:rssmq@<your-rabbitmq-ip>:5672"
        - name: RSS_MQ_QUEUE
          value: "rssmq"
        resources:
          limits:
            cpu: "25m"
            memory: "64Mi"
```
