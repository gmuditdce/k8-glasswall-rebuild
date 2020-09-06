What we have in test environment (current stage):
- Simple k8s cluster with mounted NFS storage system on all workers (nfs).
- Local docker image repository for storing base image (hub.localrepo.local/stage/glasswall-test)
- Input files (/input)

- step 1: building base image and pushing to local repo
	docker build -t hub.localrepo.local/stage/glasswall-test .
	docker push hub.localrepo.local/stage/glasswall-test
- step 2: update deployment file (glasswall-k8s-deployment.yaml)
- step 3: deploy on cluster: kubectl apply -f glasswall-k8s-deployment.yaml

result: 
	input files begins scanned by Glasswall base image and store on /output directory.

### * Note about times:
	The system timezone is configured by symlinking /etc/localtime to a binary timezone identifier in the /usr/share/zoneinfo directory.
	We mount the binary file into our docker containers so we can only synchronize the host time with the NTP server and the containers are automatically synchronized with the host server.

### * Note about Harbor registery:
	with Harbor we can manage our image(s) in master level and we can scale up and out our images registery. if we want to run a job per  file, we must pull an image per file, so we need a master level of local registery.
```	
ex:
	- update docker-compose.yml using your own environment variables. (redis address, credential, ...)
 	- docker-compose up -d
	- docker build -t hub.localrepo.local/stage/glasswall-test .
	- docker login hub.localrepo.local # Login with the credential you set in docker-compose.yml
	- docker push hub.localrepo.local/stage/glasswall-test
	- docker pull hub.localrepo.local/stage/glasswall-test
	- now you can brows pannel UI using hub.localrepo.local
```

What we must have in production environment:
- Strong k8s cluster (EKS or AKS)
- Image repo on ECR or ACR
- private and limited access storage on S3 or Azure SCS
- keeping everything isolated and private by restricted access.

- step 1: building image and pushing to ECR/ACR
	docker build -t {REPO_NAME} .
	docker push {REPO_NAME}

- Creating CI/CD pipeline based on files repo or whatever we have for starting the pipeline (S3 uploads, SNSQ, ...).
- Creating one k8s job per file. jobs automatically removed after processing files and only logs can be accessible for limited/unlimited time (https://kubernetes.io/docs/concepts/workloads/controllers/job/).
- Storing results on S3/Azure SCS



	
