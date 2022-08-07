terraform {
    backend "s3" {
        bucket = "telegram-id-bot-terraform-state"
        key    = "tfstate"
        region = "eu-west-3"
    }
    required_providers {
        aws = {
            source = "hashicorp/aws"
        }
        telegram = {
            source = "yi-jiayu/telegram"
            version = "0.2.1"
        }
    }
}

variable "telegram_api_token" {
    type        = string
    description = "Token of the Telegram API bot."
}

provider "aws" {
    # Paris
    region = "eu-west-3"
}

provider "telegram" {
    bot_token = var.telegram_api_token
}
