resource "telegram_bot_webhook" "bot" {
    url             = module.api_gateway.apigatewayv2_api_api_endpoint
    allowed_updates = ["message"]
}