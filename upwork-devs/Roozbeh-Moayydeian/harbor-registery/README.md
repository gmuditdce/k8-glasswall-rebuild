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
