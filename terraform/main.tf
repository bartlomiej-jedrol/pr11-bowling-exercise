terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "s3" {
    bucket  = "bj-terraform-states"
    key     = "state-pr11-aws-serverless-api/terraform.tfstate"
    region  = "eu-central-1"
    encrypt = true
  }
}

provider "aws" {
  region = var.region
}

# ========== Lambda ==========
resource "aws_iam_role" "lambda_iam_role" {
  name = var.lambda_iam_role
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ]
  })
}

resource "aws_lambda_function" "lambda_function" {
  function_name = var.lambda_function_name
  handler       = "main"
  runtime       = "provided.al2023"
  timeout       = 10
  filename      = "../build/main.zip"

  role = aws_iam_role.lambda_iam_role.arn
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.lambda_iam_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

# ========== API Gateway ==========
resource "aws_api_gateway_rest_api" "api_gateway" {
  name = var.api_gateway_name
}

# API Gateway Role
resource "aws_iam_role" "api_gateway_role" {
  name = var.api_gateway_role
  path = "/"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "apigateway.amazonaws.com"
        }
      }
    ]
  })
}

# API Gateway Role Policy attachments
resource "aws_iam_role_policy_attachment" "api_gateway_policy" {
  role       = aws_iam_role.api_gateway_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonAPIGatewayPushToCloudWatchLogs"
}

# API Gateway Lambda permission
resource "aws_lambda_permission" "api_gateway_lambda" {
  statement_id  = "AllowAPIGatewayInvokeLambda"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_function.function_name
  principal     = "apigateway.amazonaws.com"

  # The /*/* part allows invocation from any stage, method and resource path
  # within API Gateway.
  source_arn = "${aws_api_gateway_rest_api.api_gateway.execution_arn}/*/*"
}

# API Gateway Deployment
resource "aws_api_gateway_deployment" "api_gateway_deployment" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
}

# API Gateway Stage
resource "aws_api_gateway_stage" "api_stage" {
  deployment_id = aws_api_gateway_deployment.api_gateway_deployment.id
  rest_api_id   = aws_api_gateway_rest_api.api_gateway.id
  stage_name    = var.api_gateway_stage
  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.api_gateway_cloud_watch_group.arn
    format = jsonencode({
      requestId               = "$context.requestId"
      sourceIp                = "$context.identity.sourceIp"
      requestTime             = "$context.requestTime"
      protocol                = "$context.protocol"
      httpMethod              = "$context.httpMethod"
      resourcePath            = "$context.resourcePath"
      routeKey                = "$context.routeKey"
      status                  = "$context.status"
      responseLength          = "$context.responseLength"
      integrationErrorMessage = "$context.integrationErrorMessage"
    })
  }
}

# ========== CloudWatch ==========
resource "aws_cloudwatch_log_group" "api_gateway_cloud_watch_group" {
  name = "/aws/apigateway/${var.api_gateway_name}"
}

resource "aws_cloudwatch_log_group" "lambda_cloud_watch_group" {
  name = "/aws/lambda/${var.lambda_function_name}"
}
