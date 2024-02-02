package lambda

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

const Type = "lambda"

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)

type Layer struct {
	LayerARN *string
	Version  *int64
}

func Clean(sess *session.Session, expirationDate time.Time) error {
	logger.Printf("Begin to clean Lambda functions")
	if err := cleanFunctions(sess, expirationDate); err != nil {
		return err
	}

	logger.Printf("Begin to clean Lambda layers")
	if err := cleanLayers(sess, expirationDate); err != nil {
		return err
	}

	return nil
}

func cleanFunctions(sess *session.Session, expirationDate time.Time) error {
	lambdaClient := lambda.New(sess)
	var deleteFunctionIDs []*string
	var nextToken *string

	for {
		listFunctionsInput := &lambda.ListFunctionsInput{Marker: nextToken}
		listFunctionsOutput, err := lambdaClient.ListFunctions(listFunctionsInput)
		if err != nil {
			return fmt.Errorf("unable to retrieve functions: %w", err)
		}

		for _, lf := range listFunctionsOutput.Functions {
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
				deleteFunctionIDs = append(deleteFunctionIDs, lf.FunctionArn)
			}
		}
		if listFunctionsOutput.NextMarker == nil {
			break
		}
		nextToken = listFunctionsOutput.NextMarker
	}
	if len(deleteFunctionIDs) < 1 {
		logger.Printf("No Lambda functions to delete")
	} else {
		for _, id := range deleteFunctionIDs {
			terminateFunctionInput := &lambda.DeleteFunctionInput{FunctionName: id}
			if _, err := lambdaClient.DeleteFunction(terminateFunctionInput); err != nil {
				return fmt.Errorf("unable to delete function: %w", err)
			}
		}
		logger.Printf("Deleted %d Lambda functions", len(deleteFunctionIDs))
	}

	return nil
}

func cleanLayers(sess *session.Session, expirationDate time.Time) error {
	lambdaClient := lambda.New(sess)
	var deleteLayerVersions []Layer
	var nextToken *string

	for {
		listLayerVersionsInput := &lambda.ListLayersInput{
			Marker: nextToken,
		}

		// Retrieve layer versions from Lambda service
		listLayerVersionsOutput, err := lambdaClient.ListLayers(listLayerVersionsInput)
		if err != nil {
			return fmt.Errorf("unable to retrieve layer versions: %w", err)
		}

		// Loop through retrieved layer versions
		for _, layer := range listLayerVersionsOutput.Layers {
			layerARN := shouldDeleteLayer(layer)
			if layerARN != nil {
				created, err := time.Parse("2006-01-02T15:04:05Z0700", *layer.LatestMatchingVersion.CreatedDate)
				if err != nil {
					return fmt.Errorf("error parsing layer created time: %w", err)
				}
				if expirationDate.After(created) {
					logger.Printf("Try to delete layer version %s created-at %s", *layer.LayerArn, created)
					deleteLayerVersions = append(deleteLayerVersions, Layer{layerARN, layer.LatestMatchingVersion.Version})
				}
			}
		}

		if listLayerVersionsOutput.NextMarker == nil {
			break
		}
		nextToken = listLayerVersionsOutput.NextMarker
	}

	if len(deleteLayerVersions) < 1 {
		logger.Printf("No Lambda layers to delete")
	} else {
		for _, id := range deleteLayerVersions {
			for *id.Version > 0 {
				// Prepare input for deleting a specific layer version
				deleteLayerVersionInput := &lambda.DeleteLayerVersionInput{
					LayerName:     id.LayerARN,
					VersionNumber: id.Version,
				}
				if _, err := lambdaClient.DeleteLayerVersion(deleteLayerVersionInput); err != nil {
					return fmt.Errorf("unable to delete layer version: %w", err)
				}
				// Decrement the version number for the next iteration
				*id.Version--
			}

		}
		logger.Printf("Deleted %d Lambda layer versions", len(deleteLayerVersions))
	}

	return nil
}

func shouldDelete(lf *lambda.FunctionConfiguration) (bool, error) {
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

func shouldDeleteLayer(layerList *lambda.LayersListItem) *string {
	layerARN := layerList.LayerArn
	regexList := []string{
		".*:layer:[a-zA-Z0-9-_]+.*$",
	}

	for _, rx := range regexList {
		matched, _ := regexp.MatchString(rx, *layerARN)
		if matched {
			return layerARN
		}
		return nil
	}
	return nil
}
