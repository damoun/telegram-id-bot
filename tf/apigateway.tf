module "api_gateway" {
    source                 = "terraform-aws-modules/apigateway-v2/aws"
    version                = "2.1.0"

    name                   = "telegram-id-bot"
    description            = "HTTP to receive request for telegram id bot"
    protocol_type          = "HTTP"

    create_api_domain_name = false

    cors_configuration     = {
        allow_headers = []
        allow_methods = ["POST"]
        allow_origins = ["*"]
    }

    integrations           = {
        "ANY /{proxy+}" = {
            lambda_arn             = module.lambda_function.lambda_function_arn
            payload_format_version = "2.0"
            timeout_milliseconds   = 12000
        }
    }
}
