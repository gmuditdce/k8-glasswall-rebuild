# k8-glasswall-rebuild 

For Upwork resources, please read the [Rules of Engagement](../rules-of-engagement.md)


### Project Brief

**Objective**: Create a Kubernetes (K8) native application that is able to process 1000s of files using the Glasswall Rebuild engine

- Glasswall Rebuild is a CDR (Content Disarm and Reconstruction) tool that creates safe files by rebuilding the original file into a new file in 'known good' state (for example without Macros, Javascript or Metadata)
- Due to the fact that this engine is designed to receive potentially malicious files, the security of the execution environment is super important.   
    - In practice this means that we want to execute each file rebuild in a pristine (i.e. new) docker container
    - in K8 this means that we need 1 pod per file to be processed
- Here are the key workflows required for this project:
    - Folder exists with 100 or 1000s of files to be processed
    - K8 environments boots up and starts processing the files 
    - One pod and one container is used for each file (note that there could be other supporting containers in that pod, but only one should be using the Glasswall engine to process the target file)
    - one rebuild is complete, the engine output (rebuilt file and xml report) are stored in a separate folder, and pod is destroyed
- use CI and CD pipeline has a way to trigger workflow
- Target execution environments for the K8 environment:
    - locally (using Docker Desktop)
    - EC2 or Azure VM (with K8 installed)
    - Managed EKS (AWS or Azure)
- Implement logging solutions to visualize what is going on inside the K8 environnement

### Comms

Slack channel: https://glasswallsolu-qbp1117.slack.com/archives/C0199AQH0QK

### References

For more details about how the Glasswall Engine works, please see:

- https://github.com/filetrust/Glasswall-Rebuild-SDK-Evaluation : here you will find the engine and docker builds
- https://github.com/filetrust/Glasswall-Rebuild-SDK-Evaluation/blob/master/Getting-Started/Getting-Started.md : here are instructions for getting it running on docker
- https://engineering.glasswallsolutions.com/docs/product-descriptions/product-overview


 ## FAQ

 **Q: What to use for CI?**

 A: At the moment we are either using GitHub Actions, Azure DevOps or AWS Code Build
 
 **Q: Do you expect the developed solution to scale it up and down automatically?**
 
 A: Yes that is the idea. We want a solution that has zero 'rebuild' pods running when there are no files on the queue, and as many pods as possible (depending on config and resources) when there is files to process (with the Glasswall Rebuild engine which is already dockerized or available as an Cloud API)
 
 **Q: Do you utilize any kind of log analysis tools? If yes, could you specify which particular ones?**
 
 A: We use Elastic and DataDog in production (and the native AWS or Azure tools), but I'm open to explore new tools and workflows
 
 **Q: What does it mean "1 pod per file processed" not "running multiple pods at the same time, and allocate a file to an pod that is available"?**
 
 A: I mean that for each file to be processed by the Glassewall Engine, you need to start a new pod (with new docker container) to run the Glasswall Rebuild engine inside it
once the process is completed and the new file (and rebuild xml report) are saved in another folder, that pod (and container) needs to be terminated and destroyed
