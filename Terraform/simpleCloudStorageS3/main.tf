# Define provider
provider "aws" {
  region = "us-west-2"
}

# Define Account Data
data "aws_caller_identity" "my_account" {}

# Define S3 Resource
resource "aws_s3_bucket" "my_bucket" {
  bucket = "my-s3-bucket-${data.aws_caller_identity.my_account.account_id}"

  tags = {
    Type = "LOG"
    Tier = "Standard"
  }
}

# Define object to be uploaded in S3
resource "aws_s3_bucket_object" "readme_file" {
  bucket = aws_s3_bucket.my_bucket.bucket
  key    = "files/readme.txt"

  source = "readme.txt"
}
