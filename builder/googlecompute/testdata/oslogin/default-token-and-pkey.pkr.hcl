# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

variable "project" {
  type    = string
  default = env("GOOGLE_PROJECT_ID")
}

variable "ssh_private_key" {
  type    = string
  default = ""
}

variable "ssh_username" {
  type    = string
  default = "root"
}

variable "zone" {
  type    = string
  default = "us-central1-a"
}

locals { timestamp = regex_replace(timestamp(), "[- TZ:]", "") }

# No provided access_token or account_file should read contents of env GOOGLE_APPLICATION_CREDENTIALS
source "googlecompute" "autogenerated_1" {
  image_name           = "packer-oslogin-tester-${local.timestamp}"
  project_id           = var.project
  source_image_family  = "centos-7"
  ssh_username         = var.ssh_username
  ssh_private_key_file = "%s"
  ssh_timeout          = "30s"
  use_os_login         = true
  skip_create_image    = true
  zone                 = var.zone
}

build {
  sources = ["source.googlecompute.autogenerated_1"]

  provisioner "shell" {
    inline = ["echo hello from the other side, username is $(whoami)"]
  }
}
