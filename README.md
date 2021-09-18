# rssmq
Checks RSS feeds for new content on a regular schedule, and submits new items to a RabbitMQ queue.

# Deployment
## Kustomize
```bash
kustomize build ./kustomize | kubectl apply -f
```
