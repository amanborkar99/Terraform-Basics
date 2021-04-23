terraform {
  required_providers {
    zoom = {
      version = "0.1"
      source  = "zoom.com/temp/zoom"
    }
  }
}

provider "zoom" {
  token = "Bearer Token"
}

data "zoom_corp_data" "xyz"{

}
output "ptq" {
  value = data.zoom_corp_data.xyz
}
