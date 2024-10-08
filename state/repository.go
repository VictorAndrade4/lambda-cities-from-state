package state

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
)

func FromAlias(state string) State {
	return getSelectedState(state)
}

func unmarshallState(result *dynamodb.GetItemOutput) State {
	selectedState := State{}
	err := dynamodbattribute.UnmarshalMap(result.Item, &selectedState)

	if err != nil {
		log.Fatalf("Got error unmarshalling: %s", err)
	}
	return selectedState
}

func getSelectedState(stateName string) State {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("cities"),
		Key: map[string]*dynamodb.AttributeValue{
			"state": {
				S: aws.String(stateName),
			},
		},
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	return unmarshallState(result)
}
