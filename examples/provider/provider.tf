terraform {
  required_providers {
    saladcloud = {
      source  = "squat/saladcloud"
      version = "0.1.1"
    }
  }
}

provider "saladcloud" {
  # Configuration options
  api_key_auth = "some_saladcloud_api_key"
}