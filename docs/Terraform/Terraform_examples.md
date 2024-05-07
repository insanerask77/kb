## Terraform



```
#main.tf
module "alg" {
  source = "./terraform_alg"
  environment = "dev"
  region = "us-east4"
  zone = "us-east4-a"
}

provider "github" {
  token = var.GH_TOKEN
  owner = var.GH_OWNER
}

output "workspace" {
  value = terraform.workspace
}

output "alg_cloud_instance_name" {
  description = "The name of the alg_cloud instance from the alg module"
  value       = module.alg.alg_cloud_instance_name
}

output "alg_cloud_instance_private_ip" {
  description = "The private IP address of the alg_cloud instance from the alg module"
  value       = module.alg.alg_cloud_instance_private_ip
}

# output "elk_instance_name" {
#   description = "The name of the elk instance from the alg module"
#   value       = module.alg.elk_instance_name
# }

# output "elk_instance_private_ip" {
#   description = "The private IP address of the elk instance from the alg module"
#   value       = module.alg.elk_instance_private_ip
# }

output "wireguard_instance_name" {
  description = "The name of the wireguard instance from the alg module"
  value       = module.alg.wireguard_instance_name
}

output "wireguard_instance_private_ip" {
  description = "The private IP address of the wireguard instance from the alg module"
  value       = module.alg.wireguard_instance_private_ip
}

```

```
#Computer Engine
# VM alg cloud
resource "google_compute_instance" "alg_cloud" {
  name         = "cloud-luigi-preprod"
  machine_type = "e2-medium"
  allow_stopping_for_update = true
  zone         = "${var.zone}"
  labels = {
    environment = "${var.environment}"
  }

  boot_disk {
    auto_delete = true
    device_name = "alg-cloud-${var.environment}"
    initialize_params {
      image = "ubuntu-2204-jammy-v20221018"
      size  = 30
    }
  }
  metadata_startup_script = file("terraform_alg/start.sh")
  
  shielded_instance_config {
    enable_secure_boot = true
  }

  network_interface {
    network = "car-us-brooklyn-lab-nprd-vpc"
    subnetwork = "car-brooklyn-lab-nprd-01-us-ea4"
  }

  tags = ["allow-ingress-https-from-all-clients", "allow-ingress-http-from-all-clients", "allow-ingress-icmp-from-all-clients", "allow-ingress-ssh-from-all-clients"]

}

# VM ELK

# resource "google_compute_instance" "elk" {
#   name         = "elk-${var.environment}"
#   machine_type = "n2-standard-2"
#   zone         = "${var.zone}"
#   labels = {
#     environment = "${var.environment}"
#   }

#   boot_disk {
#     auto_delete = true
#     device_name = "elk-${var.environment}"
#     initialize_params {
#       image = "ubuntu-2204-jammy-v20221018"
#       size  = 20
#     }
#   }
#   attached_disk {
#     device_name = "elk-data-${var.environment}"
#     mode        = "READ_WRITE"
#     source      = google_compute_disk.elk-data.self_link
#   }

#   metadata_startup_script = file("terraform_alg/start.sh")
  
#   shielded_instance_config {
#     enable_secure_boot = true
#   }

#   depends_on = [
#     google_compute_disk.elk-data
#   ]
#   network_interface {
#     network = "car-us-brooklyn-lab-nprd-vpc"
#     subnetwork = "car-brooklyn-lab-nprd-01-us-ea4"
#   }
  
#   tags = ["allow-ingress-https-from-all-clients", "allow-ingress-http-from-all-clients", "allow-ingress-icmp-from-all-clients", "allow-ingress-ssh-from-all-clients"]

# }

# VM Wireguard

resource "google_compute_instance" "wireguard" {
  name         = "wireguard"
  machine_type = "e2-micro"
  allow_stopping_for_update = true
  zone         = "${var.zone}"
  labels = {
    environment = "${var.environment}"
  }

  boot_disk {
    auto_delete = true
    device_name = "wireguard-${var.environment}"
    initialize_params {
      image = "ubuntu-2204-jammy-v20221018"
      size  = 10
    }
  }
  metadata_startup_script = file("terraform_alg/start.sh")

  shielded_instance_config {
    enable_secure_boot = true
  }

  network_interface {
    network = "car-us-brooklyn-lab-nprd-vpc"
    subnetwork = "car-brooklyn-lab-nprd-01-us-ea4"
  }

  tags = ["allow-ingress-https-from-all-clients", "allow-ingress-http-from-all-clients", "allow-ingress-icmp-from-all-clients", "allow-ingress-ssh-from-all-clients"]

}
```

```
#Diks.tf
## Disks + Snapshot policies applied


# alg-cloud snapshot policy
resource "google_compute_disk_resource_policy_attachment" "attachment-alg-cloud" {
  name = google_compute_resource_policy.daily_snapshot.name
  disk = google_compute_instance.alg_cloud.name
  zone = var.zone
}

# wireguard snapshot policy
resource "google_compute_disk_resource_policy_attachment" "attachment-wireguard" {
  name = google_compute_resource_policy.daily_snapshot.name
  disk = google_compute_instance.wireguard.name
  zone = var.zone
}

# elk snapshot policy

# resource "google_compute_disk_resource_policy_attachment" "attachment-elk" {
#   name             = google_compute_resource_policy.daily_snapshot.name
#   disk             = google_compute_instance.elk.name
#   zone             = var.zone
# }

# resource "google_compute_disk_resource_policy_attachment" "attachment-elk-data" {
#   name             = google_compute_resource_policy.daily_snapshot.name
#   disk             = google_compute_disk.elk-data.name
#   zone             = var.zone
# }

# resource "google_compute_disk" "elk-data" {
#   name  = "elk-data-${var.environment}"
#   type  = "pd-ssd"
#   zone  = var.zone
#   image = ""
#   labels = {
#     environment = "${var.environment}"
#   }
#   physical_block_size_bytes = 4096
# }
```

```
main.tf

# GCP Provider

provider "google" {
  project     = var.project_id
  region      = var.region
  zone        = var.zone
}

```

```
service_account.tf
# resource "google_service_account" "github-ar" {
#   account_id   = "github-actions"
#   display_name = "github-actions"
# }
# resource "google_service_account_key" "github-ar" {
#   service_account_id = google_service_account.github-ar.id
# }
resource "google_service_account" "github-r" {
  account_id   = "github-r"
  display_name = "github-r"
}

resource "google_service_account" "github-rw" {
  account_id   = "github-rw"
  display_name = "github-rw"
}

resource "google_service_account_key" "github-r" {
  service_account_id = google_service_account.github-r.id
  private_key_type  = "TYPE_GOOGLE_CREDENTIALS_FILE"
}

resource "google_service_account_key" "github-rw" {
  service_account_id = google_service_account.github-rw.id
  private_key_type  = "TYPE_GOOGLE_CREDENTIALS_FILE"
}


resource "google_project_iam_member" "storage_object-r" {
  project = "car-us-brooklyn-lab-nprd"
  member  = "serviceAccount:${google_service_account.github-r.email}"
  role    = "roles/storage.objectViewer"
}

resource "google_project_iam_member" "storage_object-rw" {
  project = "car-us-brooklyn-lab-nprd"
  member  = "serviceAccount:${google_service_account.github-rw.email}"
  role    = "roles/storage.objectCreator"
}


# resource "google_project_iam_member" "artifactregistry_admin" {
#   project = "car-us-brooklyn-lab-nprd"
#   member  = "serviceAccount:${google_service_account.github.email}"
#   role    = "roles/artifactregistry.admin"
# }


# resource "google_project_service" "artifact_registry" {
#   project = "car-us-brooklyn-lab-nprd"
#   service = "artifactregistry.googleapis.com"
# }

# resource "google_artifact_registry_repository" "github-container-registry" {
#   repository_id = "github-container-registry"
#   location     = var.zone
#   format       = "DOCKER"
#   docker_config {
#     immutable_tags = false
#   }
# }

# output "service_account_key_public" {
#   value       = jsonencode(google_service_account_key.github-rw-test)
#   description = "Service Account Key"
# }

# resource "local_file" "google_cloud_credentials_file" {
#   filename = "/google_cloud_credentials.json"
#   content  = jsonencode(google_service_account_key.github-rw-file)
# }

resource "github_actions_secret" "google_cloud_key-r" {
  repository       = "alg-backend"
  secret_name      = "GCP_PRIVATE_KEY_R"
  plaintext_value  = google_service_account_key.github-r.private_key
}

resource "github_actions_secret" "google_cloud_key-rw" {
  repository       = "alg-backend"
  secret_name      = "GCP_PRIVATE_KEY_RW"
  plaintext_value  = google_service_account_key.github-rw.private_key
}

# resource "github_actions_secret" "google_cloud_credential" {
#   repository       = "alg-backend"
#   secret_name      = "GCP_CREDENTIALS-RW"
#   plaintext_value  = jsonencode(google_service_account_key.github-rw)
# }

# resource "github_actions_variable" "google_cloud_var" {
#   repository       = "alg-backend"
#   variable_name    = "GCP_JSON"
#   value            = google_service_account_key.github-rw
# }


resource "google_storage_bucket" "alg-assets-prod" {
 name          = "alg-assets-prod"
 location      = "us-east4"
 storage_class = "STANDARD"
 force_destroy = true
 uniform_bucket_level_access = true
}

resource "google_storage_bucket" "alg-assets-preprod" {
 name          = "alg-assets-preprod" 
 location      = "us-east4"
 storage_class = "STANDARD"
 force_destroy = true
 uniform_bucket_level_access = true
}

resource "google_storage_bucket" "alg-assets-qa" {
 name          = "alg-assets-qa"
 location      = "us-east4"
 storage_class = "STANDARD"
 force_destroy = true
 uniform_bucket_level_access = true
}


```

```
# variables.tf

# VARIABLES
variable "environment" {
  description = "Environment name (e.g., dev, prod)"
  type        = string
  default = "dev"
}


variable "project_id" {
  description = "Google Cloud project ID"
  type        = string
  default = "car-us-brooklyn-lab-nprd"
}

variable "region" {
  description = "GCE instance region"
  type        = string
  default = "us-east4"
}

variable "zone" {
  description = "GCE instance zone"
  type        = string
  default = "us-east4-a"
}


```

```
outputs.tf

output "alg_cloud_instance_name" {
  description = "The name of the alg_cloud instance"
  value       = google_compute_instance.alg_cloud.name
}

output "bucket_name_prod" {
  description = "The name of the alg_cloud bucket prod"
  value       = google_storage_bucket.alg-assets-prod.name
}

output "bucket_name_preprod" {
  description = "The name of the alg_cloud bucket preprod"
  value       = google_storage_bucket.alg-assets-preprod.name
}

output "bucket_name_qa" {
  description = "The name of the alg_cloud bucket qa"
  value       = google_storage_bucket.alg-assets-qa.name
}


output "alg_cloud_instance_private_ip" {
  description = "The private IP address of the alg_cloud instance"
  value       = google_compute_instance.alg_cloud.network_interface[0].network_ip
}

# output "elk_instance_name" {
#   description = "The name of the elk instance"
#   value       = google_compute_instance.elk.name
# }

# output "elk_instance_private_ip" {
#   description = "The private IP address of the elk instance"
#   value       = google_compute_instance.elk.network_interface[0].network_ip
# }

output "wireguard_instance_name" {
  description = "The name of the elk instance"
  value       = google_compute_instance.wireguard.name
}

output "wireguard_instance_private_ip" {
  description = "The private IP address of the elk instance"
  value       = google_compute_instance.wireguard.network_interface[0].network_ip
}
```

