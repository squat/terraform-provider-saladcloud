terraform {
  required_providers {
    saladcloud = {
      source  = "squat/saladcloud"
      version = "0.3.0"
    }
  }
}

provider "saladcloud" {
  # Configuration options
  api_key = "some_saladcloud_api_key"
}