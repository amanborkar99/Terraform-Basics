provider "hashicups" {
  username = "education"
  password = "test123"
}

resource "hashicups_order" "edu" {
  items {
    coffee {
      id = 3
    }
    quantity = 3
  }
  items {
    coffee {
      id = 2
    }
    quantity = 6
  }
}
data "hashicups_ingredients" "item1" {
  coffee_id = hashicups_order.edu.items[1].coffee[0].id
}


output "edu_order" {
  value = hashicups_order.edu
}
output "coffe_id" {
  value = data.hashicups_ingredients.item1
}
