#!/bin/bash

# Detect system architecture
UNAME_M=$(uname -m)

if [ "$UNAME_M" = "x86_64" ]; then
    LAMBDA_RIE_URL="https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie"
elif [ "$UNAME_M" = "arm64" ]; then
    #LAMBDA_RIE_URL="https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie-arm64"
    LAMBDA_RIE_URL="https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie"
else
    echo "Unsupported architecture"
    exit 1
fi

AWS_LAMBDA_RIE_DIR=./.aws-lambda-rie
AWS_LAMBDA_RIE_BINARY=${AWS_LAMBDA_RIE_DIR}/aws-lambda-rie

# Install AWS Lambda RIE
if [ ! -f "$AWS_LAMBDA_RIE_BINARY" ]; then
    echo "Installing AWS Lambda RIE in $AWS_LAMBDA_RIE_DIR from: $LAMBDA_RIE_URL"
    mkdir -p $AWS_LAMBDA_RIE_DIR
    curl -Lo $AWS_LAMBDA_RIE_BINARY $LAMBDA_RIE_URL
    chmod +x $AWS_LAMBDA_RIE_BINARY
    echo "AWS Lambda RIE successfully installed"
else
    echo "AWS Lambda RIE is already installed in $AWS_LAMBDA_RIE_DIR"
fi
