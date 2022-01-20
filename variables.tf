variable "ami" {
  type    = string
  default = "ami-0af25d0df86db00c1"
}

variable "awsRegion" {
  type    = string
  default = "ap-south-1"
}

variable "availabilty_zone" {
  type    = string
  default = "ap-south-1a"
}

variable "instance_type" {
  type    = string
  default = "t2.micro"
}

variable "size" {
  type = number
  default = 10
}