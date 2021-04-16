# CRUD Operation Using Provider

## Prerequisites :-
   * Terraform 0.14
   * Docker Installed

## Initialise Hashicups locally :-

1. Learn [HashiCups Provider repository](https://github.com/hashicorp/learn-terraform-hashicups-provider) and navigate to it.

2. Navigate to the docker_compose directory and run docker-compose up to spin up a local instance of HashiCups on port :19090.
``` BASH
cd docker_compose && docker-compose up
```
3. Verify that HashiCups is running by sending a request to its health check endpoint.
``` JAVASCRIPT
curl localhost:19090/health
```

## Installing Hashicup Provider :-

* First, download the HashiCups provider binary. Because providers are Go binaries, you can download a pre-compiled binary or build it directly from source.

#### Download Binary

``` 
curl -Lo terraform-provider-hashicups https://github.com/hashicorp/terraform-provider-hashicups/releases/download/v0.3.1/terraform-provider-hashicups_0.3.1_windows_amd64.zip
```
1. Download the Zip file and Unzip it in location `%APPDATA%\terraform.d\plugins\hashicorp.com\edu\hashicups\0.3.1\[OS_ARCH]`

2. And make it executable i.e **.exe** .

## Create new HashiCups user :-

* Follow the command line instruction for signing up.

1. `curl -X POST localhost:19090/signup -d '{"username":"xyz", "password":"123"}'`.

2. Then again sign in with same credentials. This will return the userID, username, and a JSON Web Token (JWT) authorization token. 

3. Set the `HASHICUPS_TOKEN` environment variable to the token you retrieved from invoking the /signin endpoint.
``` 
export HASHICUPS_TOKEN= 'copy the token string here which we go above'
```

## Initialize the Workspace :-

1. Add the required version and source in **main.tf** .

2. Implement the Terraform initialization, apply and destroy infrastruture.




