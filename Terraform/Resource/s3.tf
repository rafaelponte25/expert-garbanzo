resource "aws_s3_bucket" "my_second_bucket" {
  bucket = "bento-s3-second-bucket"
}

provider "aws" {
  alias  = "aws.oregon"
  region = "us-west-2"
}
