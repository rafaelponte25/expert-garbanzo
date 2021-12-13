# When a variable has no default value it will grab one
# from the file.tfvars!

variable "bucket_name" {
  description = "Bucket name for s3"
}

# Interpolation
resource "aws_s3_bucket" "variable_s3_bucket" {
  bucket = "${var.bucket_name}-rafael"
}

# Conditional with interpolation syntax
resource "aws_s3_bucket" "variable_interpolation_s3_bucket" {
  bucket = var.bucket_name == "" ? "if-vaule-is-true" : "if-value-is-false-${var.bucket_name}"
}

# Locals are customized data types
locals {
  instance_name = "dev-instance"
  instance_type = "t2.micro"
}

resource "aws_instance" "my_instance" {
  ami           = "ami-435jhkg54vl543"
  instance_type = local.instance_type

  tags {
    Name = local.instance_name
  }
}
