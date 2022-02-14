resource "aws_iam_role" "test_role" {
  name = "test_role"

  assume_role_policy = jsonencode({
    "Version": "2012-10-17",
    "Id": "sqspolicy",
    "Statement": [
        {
            "Action": [
                "sqs:GetQueueUrl",
                "sqs:SendMessage"
            ],
            "Effect": "Allow",
            "Resource": "${aws_sqs_queue.CARTHUS.arn}"
        }
    ]
  })
  tags = {
    tag-key = "tag-value"
  }
}

resource "aws_lambda_function" "farron_keep_lambda" {
  function_name = "lambda_function_name"
  role          = aws_iam_role.test_role.arn
  filename      = "lambda_function_payload.zip"
  handler       = "index.test"
  runtime       = "go1.x"
}

resource "aws_sqs_queue" "CARTHUS" {
  name = "CARTHUS"
}