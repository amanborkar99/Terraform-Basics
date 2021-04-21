# Add Authentication to a provider

## Prerequisites
* Golang installation
* Terraform
* Docker

## Steps to follow : 
1. Clone the `implement-Read` Branch from [Terraform Hashicup Provider Repository](https://github.com/hashicorp/terraform-provider-hashicups "link to Repository").

2. Add **Vendor** directory that holds all provider dependencies.

3. Build the provider Binary i.e file exe and add it to the appropriate userpath location.

``` GO
%APPDATA%\terraform.d\plugins\hashicorp.com\edu\hashicups\0.2\windows_amd64
```
4. Initiate the Docker container which conatains the demoapp API image, The **HashiCups provider** requires an instance of HashiCups. Navigate to the `docker_compose` directory then run `docker-compose up` to spin up a local instance of HashiCups on _port :19090_.

5. HashiCups requires a username and password to generate an **JSON web token (JWT)** which is used to authenticate against protected endpoints. Create a user on HashiCups named `education` with the password `test123`.

```BASH
$curl -X POST localhost:19090/signup -d '{"username":"education", "password":"test123"}'
```
6. Set a `HASHICUPS_TOKEN` environment variable to the token generated in the previous step.

7. Write the required **provider Schema** and **provider configuration**.

8. Add the `education` user credentials your provider block in main.tf.
