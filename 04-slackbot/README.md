# Slack Bot with GoLang that Calculates Age

This Slack bot is built using GoLang and calculates the age of a user based on their birthdate.

## Features
- Responds to user messages in Slack
- Calculates age based on provided birthdate

## Prerequisites
- Go 1.16 or higher
- Slack API token

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/lokeshjha/golang.git
    cd golang
    cd 04-slackbot
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up your Slack API token:
    ```sh
    export SLACK_BOT_TOKEN='your-slack-bot-token'
    ```

## Usage

Run the bot:
```sh
go run main.go
```

## Example

User: `@bot calculate age 1990-01-01`

Bot: `Your age is 32 years.`

## License

This project is licensed under the MIT License.