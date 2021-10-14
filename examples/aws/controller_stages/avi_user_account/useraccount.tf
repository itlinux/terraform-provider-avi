provider "aws" {
  shared_credentials_file = var.aws_creds_file
  region                  = var.aws_region
}

data "aws_instance" "avi_controller" {
  filter {
    name   = "tag:Name"
    values = ["${var.project_name}-terraform-controller"]
  }
}

provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_current_password
  avi_controller = data.aws_instance.avi_controller.private_ip
  avi_tenant     = "admin"
}

resource "avi_useraccount" "avi_user" {
  username     = var.avi_username
  old_password = var.avi_current_password
  password     = var.avi_new_password
}

