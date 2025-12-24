package apigw

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/stretchr/testify/assert"
)

func TestShouldDelete(t *testing.T) {
	namesToTest := []struct {
		testName    string
		mockApiName string
		expected    bool
	}{
		{
			testName:    "aws-sdk",
			mockApiName: "lambda-python-aws-sdk-wrapper-arm64-1234-ap-south-1",
			expected:    true,
		},
		{
			testName:    "ok-http",
			mockApiName: "lambda-java-okhttp-wrapper-arm64-1234-ap-south-1",
			expected:    true,
		},
		{
			testName:    "aws-sdk-agent",
			mockApiName: "lambda-java-aws-sdk-agent-amd64-1234-ap-south-1",
			expected:    true,
		},
		{
			testName:    "hello-lambda",
			mockApiName: "hello-lambda-java-okhttp-wrapper-1234-ap-south-1",
			expected:    true,
		},
		{
			testName:    "invalid",
			mockApiName: "i-am-a-bad-apigateway-name",
			expected:    false,
		},
		{
			testName: "empty",
			expected: false,
		},
		{
			testName:    "almost matching",
			mockApiName: "bad-lambda-worse-aws-sdk-agent-maybe",
			expected:    false,
		},
	}
	for _, test := range namesToTest {
		t.Run(test.testName, func(*testing.T) {
			test := test
			mockInput := types.RestApi{
				Name: aws.String(test.mockApiName),
			}
			actual, err := shouldDelete(mockInput)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
