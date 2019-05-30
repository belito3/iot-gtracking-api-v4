package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	logger "github.com/sirupsen/logrus"

	"iot-gtracking-api-v4/config"
)

// Dynamo client
var Dynamo *dynamodb.DynamoDB

func init(){
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AppConf.Region),
		Credentials: credentials.NewSharedCredentials(config.AppConf.DynamoCredentialsPath, "default"),
	})
	if err != nil {
		logger.Error("Failed connect to DynamoDB !!!")
	}

	// Session connect to DynamoDb instance
	Dynamo = dynamodb.New(sess)
}