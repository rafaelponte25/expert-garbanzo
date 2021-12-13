# Classless inter-domain routing (CIDR)
variable "vpc_cidr_block" {
  description = "VPC CIDR Block"
}

variable "public_subnet_cidr" {
  description = "Public Subnet CIDR"
}

variable "private_subnet_cidr" {
  description = "Private subnet CIDR"
}
