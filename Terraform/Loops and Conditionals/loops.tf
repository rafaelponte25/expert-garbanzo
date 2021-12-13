resource "aws_s3_bucket" "my_bucket" {
  bucket = "random-name"
  count  = var.number_of_buckets != 0 ? var.number_of_buckets : 1
}

variable "number_of_buckets" {
  default = 1
}

variable "users" {
  default = ["spiderman", "ironman", "wolverine"]
}

resource "aws_iam_user" "users" {
  # Loop each of the var elements and provide their values
  for_each = toset(var.users)
  name     = each.value
}

variable "asg_tags" {
  tyep = map(string)
  default = {
    Name = "ASG_EC2"
    Type = "m4.large"
    Team = "DEV"
  }
}

resource "aws_autoscaling_group" "asg" {
  max_size = 0
  min_size = 0

  dynamic "tag" {
    for_each = var.asg_tags
    content {
      key                 = tag.key
      value               = tag.value
      propagate_at_launch = true
    }
  }
}

# Print the users with upper case letters using "$ terraform output" command
output "uppercase_heroes" {
  value = [for hero in var.users : upper(hero) if length(herp) < 7]
}
