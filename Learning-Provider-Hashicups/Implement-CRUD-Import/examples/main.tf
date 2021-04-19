terraform {
  required_providers {
    hashicups = {
      version = "0.2"
      source  = "hashicorp.com/edu/hashicups"
    }
  }
}

provider "hashicups" {
  username = "education"
  password = "test123"
}

resource "hashicups_order" "edu" {
  items {
    coffee {
      id = 3
    }
    quantity = 15
  }
  items {
    coffee {
      id = 2
    }
    quantity = 2
  }
}

resource "hashicups_order" "sample" {}

output "ord" {
  value = hashicups_order.sample
}

output "edu_order" {
  value = hashicups_order.edu
}



module "psl" {
  source = "./coffee"

  coffee_name = "Packer Spiced Latte"
}

output "psl" {
  value = module.psl.coffee
}

