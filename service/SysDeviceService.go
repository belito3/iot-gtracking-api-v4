package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"iot-gtracking-api-v4/db/dynamo"
)

type SysDeviceService struct {
	dynamoClient *dynamodb.DynamoDB
}

func NewSysDeviceService() *SysDeviceService {
	return &SysDeviceService{dynamo.Dynamo}
}

func(s *SysDeviceService) GetLastReport(deviceId string, customerId string) (rs []map[string]interface{}, err error){

	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":deviceId": {
				S: aws.String(deviceId),
			},
			":customerId":{
				S: aws.String(customerId),
			},
		},
		KeyConditionExpression: aws.String("deviceId = :deviceId and customerId = :customerId"),
		//ProjectionExpression:   aws.String("SongTitle"),
		TableName:              aws.String("iot-gtracking-last-location"),
	}

	result, err := s.dynamoClient.Query(input)

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &rs)
	if err != nil {
		return nil, err
	}

	return  rs,nil
}