# SONiC-restapi

## Description
This is a configuration agent which exposes HTTP endpoints to perform dynamic network configuration on switches running SONiC. It restful API server is `pemgr-server`

## Getting Started
### Build Rest-API
  1. Run `docker build -f Dockerfile.release . -t rest-api-image`
  2. The above should generate Docker image which is used for local development on your VM
  3. Run `docker images` to check if Emgr-restapi docker was generated <br/>
      		<pre>`REPOSITORY                     TAG       IMAGE ID       CREATED        SIZE` <br/>
              	 `rest-api-image                 latest    49df32203088   18 hours ago   336MB`<br/>
            </pre>
  4. `rest-api-image` is for local testing on a dev VM
### Running Rest-API container
#### Run Rest-API container locally on a VM
  1. `docker run --name rest-api --privileged -p 8880:8880 -p 442:442 -v /var/run/redis:/var/run/redis -itd rest-api-image:latest`
  
####  Login to Rest-API container and check logs
  1. `docker exec -it rest-api bash`
  
#### Run Rest-API container on a switch
  1. `docker run --name rest-api --privileged --net=host --uts=host -v /var/run/redis:/var/run/redis -itd rest-api-image:latest`