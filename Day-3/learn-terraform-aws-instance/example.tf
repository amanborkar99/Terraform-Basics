terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }
}

provider "aws" {
  profile = "amanborkar"
  access_key = "access key"
  secret_key = "secret key"
  region  = "ap-south-1"
}

resource "aws_instance" "example" {
  ami           = "ami-0d758c1134823146a"
  instance_type = "t2.micro"

  tags = {
    Name = "ExampleInstance"
  }
}