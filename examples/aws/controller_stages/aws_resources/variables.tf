variable "aws_access_key" {
  default = ""
}

variable "aws_secret_key" {
  default = ""
}

variable "aws_region" {
  type    = string
  default = "us-west-2"
}

variable "aws_vpc_id" {
  type    = string
  default = "vpc-19295f7c"
}

variable "avi_username" {
  type    = string
  default = "admin"
}

variable "avi_password" {
}

variable "aws_availability_zone" {
  type    = string
  default = "us-west-2a"
}

variable "aws_subnet_mask" {
  default = 24
}

variable "project_name" {
}

variable "aws_creds_file" {
  default = "~/.aws/credentials"
}

variable "aws_subnets" {
  type    = list(string)
  default = [""]
}

variable "avi_controller_ami" {
  default = "ami-1426aa6c"
}

variable "aws_availability_zones" {
  type    = list(string)
  default = ["us-west-2a", "us-west-2b", "us-west-2c"]
}

