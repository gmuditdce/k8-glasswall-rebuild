## Download Dependencies
    curl -L https://github.com/jenkins-x/jx/releases/download/v2.0.500/jx-linux-amd64.tar.gz | tar xzv
    sudo mv jx /usr/local/bin
## Create a Github Access Token
Navigate to (https://github.com/settings/tokens/new?scopes=repo,read:user,read:org,user:email,write:repo_hook,delete_repo), Give the token a name and save this somewhere safe.

## Create empty Cluster
    gcloud auth login
    gcloud config set project 'Project Name'
    jx create cluster gke --cluster-name glasswall2ls --machine-type e2-medium --max-num-nodes 4 --min-num-nodes 3 --region us-east1-b --long-term-storage --skip-installation --skip-login

## Install JenkinsX Components
    jx boot

## Select the JX namespace
    jx ns jx

## View all deployments
    cd ~/.jx/bin/
    ./kubectl get deployments


## Deploy Existing App on JX
it will create the required files like helm charts, Dockerfile if not Existing

    jx import --url https://github.com/M-Ayman/CI-CD.git

## Check logs and Activiaty
    jx get build logs M-Ayman/CI-CD/master
    jx get activity -f CI-CD -w

## Get the url of the App
    jx get applications
## Test Build Triggers
Change any line of code and wait 2-5 minutes

## Check a new created pod with the new chnages
    ./kubectl get pods -n jx-staging
