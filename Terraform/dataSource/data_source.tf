#Example of Data Source
data "aws_instance" "ec2_data" {
  instance_id = "234524frqewff4534"
}

resource "aws_instance" "ec2_instance" {
  ami           = "ami-43242fsd"
  instance_type = "t2.micro"
  key_name      = data.aws_instance.ec2_data.key_name

  tags = {
    Name = "rafael-ec2-instance"
  }
}
