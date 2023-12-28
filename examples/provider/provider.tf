terraform {
  required_providers {
    saladcloud = {
      source  = "squat/saladcloud"
      version = "0.2.0"
    }
  }
}

provider "saladcloud" {
  # Configuration options
  api_key = "some_saladcloud_api_key"
}