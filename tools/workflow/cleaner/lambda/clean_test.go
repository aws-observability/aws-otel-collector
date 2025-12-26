package lambda

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/stretchr/testify/assert"
)

func TestShouldDeleteLayer(t *testing.T) {
	namesToTest := []struct {
		testName     string
		mockLayerARN string
		expected     bool
	}{
		{
			testName:     "aws-otel-go",
			mockLayerARN: "arn:aws:lambda:us-east-1:611364707713:layer:aws-otel-go-wrapper-amd64-0ebbb233ce5cb004395b3d69cd753c7977784bf6:2",
			expected:     true,
		},
		{
			testName:     "aws-otel-lambda-java",
			mockLayerARN: "arn:aws:lambda:us-east-1:611364707713:layer:aws-otel-lambda-java-agent-03fe7b9dec95821fe14f0c583b25080241cfed41:1",
			expected:     true,
		},
		{
			testName:     "aws-otel-lambda-java-aws-sdk",
			mockLayerARN: "arn:aws:lambda:us-east-1:611364707713:layer:aws-otel-lambda-java-aws-sdk-agent-0faf1410a828a873c6c9a94e112815a147a3f03b:1",
			expected:     true,
		},
		{
			testName:     "aws-otel-nodejs",
			mockLayerARN: "arn:aws:lambda:us-east-1:611364707713:layer:aws-otel-nodejs-wrapper-amd64-01ccbd39e15554dc244260559f1181e9b0556287:3",
			expected:     true,
		},
		{
			testName:     "opentelemetry",
			mockLayerARN: "arn:aws:lambda:us-east-1:611364707713:layer:opentelemetry-python-aws-sdk-wrapper-arm64-1763695497:1",
			expected:     true,
		},
	}

	for _, test := range namesToTest {
		mockLayerARN := test.mockLayerARN
		t.Run(test.testName, func(t *testing.T) {
			mockInput := types.LayersListItem{
				LayerArn: aws.String(mockLayerARN),
				LatestMatchingVersion: &types.LayerVersionsListItem{
					CreatedDate: aws.String("2024-02-14T10:00:00Z"),
				},
			}
			actual := shouldDeleteLayer(mockInput, time.Now())
			assert.Equal(t, test.expected, actual)
		})
	}
}
