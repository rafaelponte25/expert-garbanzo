provider "aws" {
  region="us-west-2"
}

resource "aws_dynamodb_table" "first_dynamodb_table" {
  name = "GameScores"
  billing_mode = "PROVISIONED"
  read_capacity = 10
  write_capacity = 10
  hash_key = "UserId"
  range_key = "GameTitle"

  attribute {
    name="UserdId"
    type="S" // type string
  }
  attribute {
    name="GameTitle"
    type="S" //string
  }
  attribute {
    name="TopScore"
    type="N" //type number
  }

  ttl {
    attribute_name="TimeToExist"
    enabled=false
  }
  global_secondary_index {
    name="GameTitle"
    hash_key="GameTitleIndex"
    projection_type = "INCLUDE"
    range_key = "TopScore"
    read_capacity = 10
    write_capacity = 10
    non_key_attributes = ["UserId"]
  }

  tags {
    Name="GameScores"
    Type="Game"
  }
}