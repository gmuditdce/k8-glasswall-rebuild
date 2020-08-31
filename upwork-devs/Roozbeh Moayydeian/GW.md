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



	