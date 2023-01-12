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

func Clean(sess *session.Session, expirationDate time.Time) error {
	logger.Printf("Begin to clean Lambda functions")

	lambdaclient := lambda.New(sess)
	var deleteFunctionIDs []*string
	var nextToken *string
	for {
		ListFunctionsInput := &lambda.ListFunctionsInput{Marker: nextToken}
		ListFunctionsOutput, err := lambdaclient.ListFunctions(ListFunctionsInput)
		if err != nil {
			return fmt.Errorf("unable to retrieve APIs: %w", err)
		}
		for _, lf := range ListFunctionsOutput.Functions {
			doesNameMatch, err := shouldDelete(lf)
			if err != nil {
				return fmt.Errorf("error during pattern match: %w", err)
			}
			created, err := time.Parse("2006-01-02T15:04:05Z0700", *lf.LastModified)
			if err != nil {
				return fmt.Errorf("error parting last modified time: %w", err)
			}
			if expirationDate.After(created) && doesNameMatch {
				logger.Printf("Try to delete function %s created-at %s", *lf.FunctionArn, *lf.LastModified)
				deleteFunctionIDs = append(deleteFunctionIDs, lf.FunctionArn)
			}
		}
		if ListFunctionsOutput.NextMarker == nil {
			break
		}
		nextToken = ListFunctionsOutput.NextMarker
	}
	if len(deleteFunctionIDs) < 1 {
		logger.Printf("No Lambda functions to delete")
		return nil
	}

	for _, id := range deleteFunctionIDs {
		terminateFunctionInput := &lambda.DeleteFunctionInput{FunctionName: id}
		if _, err := lambdaclient.DeleteFunction(terminateFunctionInput); err != nil {
			return fmt.Errorf("unable to delete function: %w", err)
		}
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
