FROM python:3.10-alpine3.15
COPY . /rssmq
CMD /rssmq --config /opt/rssmq.json
