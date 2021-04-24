terraform {
  required_providers {
    zoom = {
      version = "0.1"
      source  = "zoom.com/temp/zoom"
    }
  }
}

provider "zoom" {
  token = "Bearer token"
}


# data "zoom_corp_data" "xyz"{
    
# }
# output "ptq" {
#   value = data.zoom_corp_data.xyz
# }
# resource "zoom_resource" "mx" {
#   email = "notema3125@gridmire.com"
#   first_name = "champaklal"
#   last_name = "gada"
#   type = 1

# }
resource "zoom_resource" "tx" {
  email = "amanborkar12345@gmail.com"
  first_name = "Aman"
  last_name = "VIPBoss"
  type = 1
}

 output "TMKOC" {
   value = zoom_resource.tx
 }

# resource "zoom_resource" "tl"{
#   email = "amanborkar12345@gmail.com"
#   first_name = "Aman"
#   last_name = "VIPBoss"
#   type = 1
# }
# resource "zoom_resource" "t3"{
#   email = "notema3125@gridmire.com"
#   first_name = "Anupam"
#   last_name = "Director"
#   type = 1
# }