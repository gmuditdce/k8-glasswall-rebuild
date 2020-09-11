# Glasswall Rebuild usage with argo workflow and argo events

## Achitecture (taken from https://argoproj.github.io/argo-events/triggers/argo-workflow/)

![Glasswall Rebuild architecture overview](https://github.com/argoproj/argo-events/blob/master/docs/assets/argo-workflow-trigger.png)

Basically :

- Argo event will receive events from multiple possible sources (rabbitmq, minio, s3, ..)
- Once an event is received, it will trigger execution of a workflow, passing parameters to that workflow
- The workflow will then be executed and rebuild the file

## Workflow View

This is a simple workflow implemented for the PoC, the intermediate steps(hasher, file checker) now don't do anything it was just to present how the flow can be orchestrated.
In the image bellow we can see the execution and each pod logs on the UI

![Controller flow diagram](workflow1.png)
![Controller flow diagram](workflow2.png)


