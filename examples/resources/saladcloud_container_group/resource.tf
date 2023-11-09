resource "saladcloud_container_group" "my_containergroup" {
  container = {
    command = [
      "...",
    ]
    environment_variables = {
      "Northeast" = "..."
      "Cadillac"  = "..."
    }
    image = "...my_image..."
    logging = {
      new_relic = {
        host          = "...my_host..."
        ingestion_key = "...my_ingestion_key..."
      }
      splunk = {
        host  = "...my_host..."
        token = "...my_token..."
      }
      tcp = {
        host = "...my_host..."
        port = 6
      }
    }
    registry_authentication = {
      aws_ecr = {
        access_key_id     = "...my_access_key_id..."
        secret_access_key = "...my_secret_access_key..."
      }
      basic = {
        password = "...my_password..."
        username = "Gregorio42"
      }
      docker_hub = {
        personal_access_token = "...my_personal_access_token..."
        username              = "Letitia20"
      }
      gcp_gcr = {
        service_key = "...my_service_key..."
      }
    }
    resources = {
      cpu       = 6
      gpu_class = "...my_gpu_class..."
      memory    = 9
    }
  }
  country_codes = [
    "la",
  ]
  display_name = "...my_display_name..."
  liveness_probe = {
    exec = {
      command = [
        "...",
      ]
    }
    failure_threshold = 2
    grpc = {
      port    = 10
      service = "...my_service..."
    }
    http = {
      headers = [
        {
          name  = "Maxine Aufderhar"
          value = "...my_value..."
        },
      ]
      path   = "...my_path..."
      port   = 5
      scheme = "http"
    }
    initial_delay_seconds = 1
    period_seconds        = 5
    success_threshold     = 3
    tcp = {
      port = 10
    }
    timeout_seconds = 5
  }
  name = "Rick Dickinson"
  networking = {
    auth     = false
    dns      = "...my_dns..."
    port     = 8
    protocol = "http"
  }
  organization_name = "...my_organization_name..."
  project_name      = "...my_project_name..."
  readiness_probe = {
    exec = {
      command = [
        "...",
      ]
    }
    failure_threshold = 2
    grpc = {
      port    = 10
      service = "...my_service..."
    }
    http = {
      headers = [
        {
          name  = "Mr. Patsy Kuhn"
          value = "...my_value..."
        },
      ]
      path   = "...my_path..."
      port   = 1
      scheme = "http"
    }
    initial_delay_seconds = 6
    period_seconds        = 9
    success_threshold     = 3
    tcp = {
      port = 2
    }
    timeout_seconds = 10
  }
  replicas       = 8
  restart_policy = "always"
  startup_probe = {
    exec = {
      command = [
        "...",
      ]
    }
    failure_threshold = 10
    grpc = {
      port    = 5
      service = "...my_service..."
    }
    http = {
      headers = [
        {
          name  = "Patty Bergnaum"
          value = "...my_value..."
        },
      ]
      path   = "...my_path..."
      port   = 2
      scheme = "http"
    }
    initial_delay_seconds = 6
    period_seconds        = 7
    success_threshold     = 5
    tcp = {
      port = 9
    }
    timeout_seconds = 2
  }
}