// Terraform variables.
variable "region" {
  type    = string
  default = "eu-central-1"
}

variable "lambda_iam_role" {
  type    = string
  default = "pr11-lambda-role"
}

variable "lambda_inline_policy" {
  type    = string
  default = "pr11-lambda-inline-policy"
}

variable "lambda_function_name" {
  type    = string
  default = "pr11-lambda"
}

variable "api_gateway_name" {
  type    = string
  default = "pr11-api-gateway"
}

variable "api_gateway_role" {
  type    = string
  default = "pr11-api-gateway-role"
}

variable "api_gateway_stage" {
  type    = string
  default = "dev"
}
