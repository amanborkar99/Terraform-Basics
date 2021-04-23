terraform {
  required_providers {
    zoom = {
      version = "0.1"
      source  = "zoom.com/temp/zoom"
    }
  }
}

provider "zoom" {
  token = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTkxMDE5NjksImlhdCI6MTYxOTA5NjU2OX0.WjPXHfy4bhZpcPdSPa0U33mLARQqobaoifGs03m_4Jk"
}

data "zoom_corp_data" "xyz"{

}
output "ptq" {
  value = data.zoom_corp_data.xyz
}