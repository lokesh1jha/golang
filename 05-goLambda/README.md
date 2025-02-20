# Go Lambda Project

This project will be using AWS Lambda and Go.

## Prerequisites

- Go 1.16 or later
- AWS CLI
- AWS SAM CLI

## Setup

1. **Install Go**: Follow the instructions on the [official Go website](https://golang.org/doc/install).
2. **Install AWS CLI**: Follow the instructions on the [official AWS CLI website](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html).
3. **Install AWS SAM CLI**: Follow the instructions on the [official AWS SAM CLI website](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html).

## Project Structure

```
.
├── README.md
├── main.go
├── template.yaml
└── Makefile
```

## Building and Deploying

1. **Build the project**:
    ```sh
    sam build
    ```

2. **Deploy the project**:
    ```sh
    sam deploy --guided
    ```

## License

This project is licensed under the MIT License.