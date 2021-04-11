# Configure docker provider
#
# https://www.terraform.io/docs/providers/docker/index.html
terraform {
  required_providers {
    docker = {
      source  = "terraform-providers/docker"
      //host    = "npipe:////.//pipe//docker_engine"
    }
  }
//required_version = "~> 0.14"
}
