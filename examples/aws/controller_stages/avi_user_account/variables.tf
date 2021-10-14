variable "aws_region" {
  type    = string
  default = "us-west-2"
}

variable "avi_username" {
  type    = string
  default = "admin"
}

variable "avi_current_password" {
}

variable "avi_new_password" {
}

variable "project_name" {
}

variable "aws_creds_file" {
  default = "~/.aws/credentials"
}

