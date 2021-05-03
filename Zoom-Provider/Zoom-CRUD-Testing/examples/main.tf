terraform {
  required_providers {
    zoom = {
      version = "1.0"
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
#   email = "wohadep882@laraskey.com"
#   first_name = "Jethalal"
#   last_name = "gada"
#   type = 1

# }
# data "zoom_read_single_user" "mp"{
#   email = "francis@wizrocket.com"
# }
# output "name"{
#   value = data.zoom_read_single_user.mp
# }
resource "zoom_resource" "tp" {
  email = "amanborkar17@cse.iiitp.ac.in"
  first_name = "Aman"
  last_name = "Borkar"
  type = 1
}

#  output "TMKOC" {
#    value = zoom_resource.tp
#  }

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