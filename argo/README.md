# Glasswall Rebuild usage with argo workflow and argo events

## Architecture (taken from https://argoproj.github.io/argo-events/triggers/argo-workflow/)

![Glasswall Rebuild architecture overview](https://github.com/argoproj/argo-events/blob/master/docs/assets/argo-workflow-trigger.png)

Basically:

- Argo event will receive events from multiple possible sources (rabbitmq, minio, s3 etc.)
- Once an event is received, it will trigger execution of a workflow, passing the filename to rebuild to that workflow
- The workflow will then be executed, i.e. generate a report

## Setup

### Overview

- Install argo workflows and argo events following official documentation
- Setup argo events bus
- Setup Minio event source. Example manifest is in this repo: minio-eventsource.yml
- Setup Minio event sensor. Example manifest is in this repo: minio-event-sensor.yml (everything happens in the sensor, we can see that it will trigger the flow around line 36. and we are passing the notifications parameters to the flow. See from line 19.)

### Detailed

Install Argo Events

```bash
kubectl create ns argo-events
kubectl apply -n argo-events -f https://raw.githubusercontent.com/argoproj/argo/stable/manifests/quick-start-postgres.yaml
kubectl apply -f https://raw.githubusercontent.com/argoproj/argo-events/stable/manifests/install.yaml
kubectl apply -n argo-events -f https://raw.githubusercontent.com/argoproj/argo-events/stable/examples/eventbus/native.yaml
```

Set parameters in `minio-event-sensor.yml`

```yaml
parameters:
  # Set the same endpoint in minio-eventsource.yml
  - name: inputEndpoint
    value: YOUR_INPUT_MINIO_ENDPOINT
  - name: outputEndpoint
    value: YOUR_OUTPUT_MINIO_ENDPOINT
```

Supply the MinIO credentials

```bash
kubectl -n argo-events create secret generic my-input-credentials --from-literal=accesskey=YOUR_INPUT_ACCESS_KEY --from-literal=secretkey=YOUR_INPUT_SECRET_KEY
kubectl -n argo-events create secret generic my-output-credentials --from-literal=accesskey=YOUR_OUTPUT_ACCESS_KEY --from-literal=secretkey=YOUR_OUTPUT_SECRET_KEY
```

Install the listener

```bash
kubectl apply -n argo-events -f minio-eventsource.yml -f minio-event-sensor.yml
```
