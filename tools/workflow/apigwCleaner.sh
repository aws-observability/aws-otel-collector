#!/bin/bash
: <<'END_DOCUMENTATION'
`delete-lambda-apigw-rest-apis.sh`
This script deletes lingering APIGW REST APIs in the currently configured aws
account as read by `AWS_PROFILE`. Resources that failed to pass tests don't get
deleted on AWS, and we are limited to 120 endpoints per account.
See more:
https://docs.aws.amazon.com/apigateway/latest/developerguide/limits.html
END_DOCUMENTATION
IFS=$'\n'
name_and_id_apis=$(aws apigateway get-rest-apis --query 'items[*].[name,id]' --output text)
for n_i in $name_and_id_apis; do
    name=$(cut -f 1 <<< $n_i)
    if [[ $name == *lambda-*-wrapper-* ]]; then
        id=$(cut -f 2 <<< $n_i)
        echo "Delete APIGW Rest API: name=$name, id=$id"
        aws apigateway delete-rest-api --rest-api-id $id
        # Docs say only 1 APIGW REST API can be delete every 30 seconds :)
        # https://docs.aws.amazon.com/apigateway/latest/developerguide/limits.html
        sleep 30
    fi
done
unset IFS