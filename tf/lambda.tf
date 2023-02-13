module "lambda_function" {
    source                 = "terraform-aws-modules/lambda/aws"
    version                = "4.10.0"

    function_name          = "telegram-id-bot"
    description            = "Function to respond for telegram id bot"
    handler                = "main"
    runtime                = "go1.x"

    source_path            = "${path.module}/../main"

    publish                = true
    maximum_retry_attempts = 0

    allowed_triggers       = {
        APIGatewayAny = {
            service    = "apigateway"
            source_arn = "${module.api_gateway.default_apigatewayv2_stage_execution_arn}/*/*"
        }
    }

    environment_variables = {
        TELEGRAM_API_TOKEN = var.telegram_api_token
    }
}
