package apigw

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const Type = "apigw"

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)

func Clean(ctx context.Context, cfg aws.Config, expirationDate time.Time) error {
	logger.Printf("Begin to clean APIGW clients")

	apigwclient := apigateway.NewFromConfig(cfg)
	var deleteGatewayIds []string
	var nextToken *string

	for {
		getRestApisOutput, err := apigwclient.GetRestApis(ctx, &apigateway.GetRestApisInput{Position: nextToken})
		if err != nil {
			return fmt.Errorf("unable to retrieve APIs: %w", err)
		}
		for _, gw := range getRestApisOutput.Items {
			doesNameMatch, err := shouldDelete(gw)
			if err != nil {
				return fmt.Errorf("error during pattern match: %w", err)
			}
			if expirationDate.After(*gw.CreatedDate) && doesNameMatch {
				logger.Printf("Try to delete gateway %s created-at %v", *gw.Id, gw.CreatedDate)
				deleteGatewayIds = append(deleteGatewayIds, *gw.Id)
			}
		}
		if getRestApisOutput.Position == nil {
			break
		}
		nextToken = getRestApisOutput.Position
	}

	if len(deleteGatewayIds) < 1 {
		logger.Printf("No API Gateways to delete")
		return nil
	}

	for _, id := range deleteGatewayIds {
		_, err := apigwclient.DeleteRestApi(ctx, &apigateway.DeleteRestApiInput{RestApiId: aws.String(id)})
		if err != nil {
			return fmt.Errorf("unable to delete gateway: %w", err)
		}
		// 30 second request limit to delete rest api
		// https://docs.aws.amazon.com/apigateway/latest/developerguide/limits.html#api-gateway-control-service-limits-table
		time.Sleep(31 * time.Second)
	}
	return nil
}

func shouldDelete(gw types.RestApi) (bool, error) {
	gwName := gw.Name
	regexList := []string{
		"\\Alambda-[a-z]+-aws-sdk-.*$",
		"\\Ahello-lambda-[a-z]+-okhttp-.*",
		"\\Alambda-[a-z]+-okhttp-wrapper-.*$",
	}

	for _, rx := range regexList {
		matches, err := regexp.MatchString(rx, *gwName)
		if err != nil {
			return false, fmt.Errorf("error during regex match: %w", err)
		}
		if matches {
			return true, nil
		}
	}
	return false, nil
}
