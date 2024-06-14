terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "5.33.0"
    }
  }
}

provider "google" {
  project = "eth-parser-426407"
}

import {
  to = google_cloud_run_v2_service.default
  id = "projects/eth-parser-426407/locations/asia-southeast1/services/eth-parser"
}

import {
  to = google_cloud_run_v2_service_iam_member.noauth
  id = "projects/eth-parser-426407/locations/asia-southeast1/services/eth-parser roles/run.invoker allUsers"
}


variable "image_id" {
  type        = string
  default     = "asia-southeast1-docker.pkg.dev/eth-parser-426407/eth-parser/eth-parser:latest"
  description = "Container image to deploy. Defaults to asia-southeast1-docker.pkg.dev/eth-parser-426407/eth-parser/eth-parser:latest"
}


resource "google_cloud_run_v2_service" "default" {
  name     = "eth-parser"
  location = "asia-southeast1"
  client   = "terraform"

  template {
    containers {
      image = var.image_id
    }
  }
}

resource "google_cloud_run_v2_service_iam_member" "noauth" {
  name     = "eth-parser"
  location = "asia-southeast1"
  role     = "roles/run.invoker"
  member   = "allUsers"
}
