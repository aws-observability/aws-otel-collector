package lambda

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

const Type = "lambda"

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)

type Layer struct {
	LayerARN string
	Version  int64
}

func Clean(ctx context.Context, cfg aws.Config, expirationDate time.Time) error {
	return errors.Join(cleanFunctions(ctx, cfg, expirationDate), cleanLayers(ctx, cfg, expirationDate))
}

func cleanFunctions(ctx context.Context, cfg aws.Config, expirationDate time.Time) error {
	logger.Printf("Begin to clean Lambda functions")
	lambdaClient := lambda.NewFromConfig(cfg)
	var deleteFunctionIDs []string

	paginator := lambda.NewListFunctionsPaginator(lambdaClient, &lambda.ListFunctionsInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("unable to retrieve functions: %w", err)
		}

		for _, lf := range output.Functions {
			doesNameMatch, err := shouldDelete(lf)
			if err != nil {
				return fmt.Errorf("error during pattern match: %w", err)
			}
			created, err := time.Parse("2006-01-02T15:04:05Z0700", *lf.LastModified)
			if err != nil {
				return fmt.Errorf("error parsing last modified time: %w", err)
			}
			if expirationDate.After(created) && doesNameMatch {
				logger.Printf("Try to delete function %s created-at %s", *lf.FunctionArn, *lf.LastModified)
				deleteFunctionIDs = append(deleteFunctionIDs, *lf.FunctionArn)
			}
		}
	}

	if len(deleteFunctionIDs) < 1 {
		logger.Printf("No Lambda functions to delete")
		return nil
	}
	for _, id := range deleteFunctionIDs {
		_, err := lambdaClient.DeleteFunction(ctx, &lambda.DeleteFunctionInput{FunctionName: aws.String(id)})
		if err != nil {
			return fmt.Errorf("unable to delete function: %w", err)
		}
		logger.Printf("Deleted %d Lambda functions", len(deleteFunctionIDs))
	}

	return nil
}

func cleanLayers(ctx context.Context, cfg aws.Config, expirationDate time.Time) error {
	logger.Printf("Begin to clean Lambda layers")
	lambdaClient := lambda.NewFromConfig(cfg)
	var deleteLayerVersions []Layer

	paginator := lambda.NewListLayersPaginator(lambdaClient, &lambda.ListLayersInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("unable to retrieve layer versions: %w", err)
		}

		// Loop through retrieved layer versions
		for _, layer := range output.Layers {
			if shouldDeleteLayer(layer, expirationDate) {
				logger.Printf("Try to delete layer version %s created-at %s", *layer.LayerArn, *layer.LatestMatchingVersion.CreatedDate)
				deleteLayerVersions = append(deleteLayerVersions, Layer{*layer.LayerArn, layer.LatestMatchingVersion.Version})
			}
		}
	}

	if len(deleteLayerVersions) < 1 {
		logger.Printf("No Lambda layers to delete")
		return nil
	}
	for _, id := range deleteLayerVersions {
		for id.Version > 0 {
			// Prepare input for deleting a specific layer version
			_, err := lambdaClient.DeleteLayerVersion(ctx, &lambda.DeleteLayerVersionInput{
				LayerName:     aws.String(id.LayerARN),
				VersionNumber: aws.Int64(id.Version),
			})
			if err != nil {
				return fmt.Errorf("unable to delete layer version: %w", err)
			}
			// Decrement the version number for the next iteration
			id.Version--
		}
		logger.Printf("Deleted %d Lambda layer versions", len(deleteLayerVersions))
	}

	return nil
}

func shouldDelete(lf types.FunctionConfiguration) (bool, error) {
	lfName := lf.FunctionName
	regexList := []string{
		"\\Alambda-[a-z]+-aws-sdk-.*$",
		"\\Ahello-lambda-[a-z]+-okhttp-.*",
		"\\Alambda-[a-z]+-okhttp-wrapper-.*$",
	}

	for _, rx := range regexList {
		matches, err := regexp.MatchString(rx, *lfName)
		if err != nil {
			return false, fmt.Errorf("error during regex match: %w", err)
		}
		if matches {
			return true, nil
		}
	}
	return false, nil
}

func shouldDeleteLayer(layerList types.LayersListItem, expirationDate time.Time) bool {
	layerARN := layerList.LayerArn
	regexList := []string{
		".*:layer:aws-otel-collector.*$",
		".*:layer:aws-otel-java.*$",
		".*:layer:aws-otel-lambda.*$",
		".*:layer:aws-otel-nodejs.*$",
		".*:layer:aws-otel-go-wrapper.*$",
		".*:layer:opentelemetry.*$",
		".*:layer:aws-observability.*$",
		".*:layer:aws-distro-for-opentelemetry.*$",
	}

	for _, rx := range regexList {
		matched, _ := regexp.MatchString(rx, *layerARN)
		if matched {
			created, err := time.Parse("2006-01-02T15:04:05Z0700", *layerList.LatestMatchingVersion.CreatedDate)
			if err != nil {
				logger.Printf("Error Parsing the created time for layer %s", err)
				return false
			}
			if expirationDate.After(created) {
				return true
			}
		}
	}
	return false
}
