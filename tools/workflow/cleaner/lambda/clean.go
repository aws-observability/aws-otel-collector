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

type Pair struct {
	Key   *string
	Value *int64
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
	var deleteLayerVersionsIDs []Pair
	var nextToken *string

	for {
		listLayerVersionsInput := &lambda.ListLayersInput{
			Marker: nextToken,
		}
		listLayerVersionsOutput, err := lambdaClient.ListLayers(listLayerVersionsInput)
		if err != nil {
			return fmt.Errorf("unable to retrieve layer versions: %w", err)
		}

		for _, layer := range listLayerVersionsOutput.Layers {
			layered := shouldDeleteLayer(layer)
			//fmt.Println(layered)
			if layered != nil {
				created, err := time.Parse("2006-01-02T15:04:05Z0700", *layer.LatestMatchingVersion.CreatedDate)
				if err != nil {
					return fmt.Errorf("error parsing layer created time: %w", err)
				}
				if expirationDate.After(created) {
					logger.Printf("Try to delete layer version %s created-at %s", *layer.LayerArn, created)
					deleteLayerVersionsIDs = append(deleteLayerVersionsIDs, Pair{layered, layer.LatestMatchingVersion.Version})
				}

			}

		}

		if listLayerVersionsOutput.NextMarker == nil {
			break
		}
		nextToken = listLayerVersionsOutput.NextMarker
	}

	if len(deleteLayerVersionsIDs) < 1 {
		logger.Printf("No Lambda layers to delete")
	} else {
		for _, id := range deleteLayerVersionsIDs {
			fmt.Println("Trying to delete the layer in the list that is acquired", *id.Key)
			for *id.Value > 0 {
				deleteLayerVersionInput := &lambda.DeleteLayerVersionInput{LayerName: id.Key, VersionNumber: id.Value}
				if _, err := lambdaClient.DeleteLayerVersion(deleteLayerVersionInput); err != nil {
					return fmt.Errorf("unable to delete layer version: %w", err)
				}
				*id.Value--
			}

		}
		logger.Printf("Deleted %d Lambda layer versions", len(deleteLayerVersionsIDs))
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
	fmt.Println(*layerARN)
	regexList := []string{
		".*:layer:[a-zA-Z0-9-_]+.*$",
	}

	for _, rx := range regexList {
		matched, _ := regexp.MatchString(rx, *layerARN)
		fmt.Println(matched)
		if matched {
			fmt.Printf("It has matched the regexlist %s\n", *layerARN)
			return layerARN
		}
		return nil
	}
	return nil
}
