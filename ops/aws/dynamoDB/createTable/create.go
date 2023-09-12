package createTable

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var Client *dynamodb.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	Client = dynamodb.NewFromConfig(cfg)
}

func CreateTable(tableName *string) error {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
		TableName:   tableName,
	}

	_, err := Client.CreateTable(context.TODO(), input)
	if err != nil {
		return errors.New("error creating table, " + *tableName + " - " + err.Error())
	}

	return nil
}

func Wait(tableName *string) error {
	waiter := dynamodb.NewTableExistsWaiter(Client)
	if err := waiter.Wait(context.TODO(),
		&dynamodb.DescribeTableInput{TableName: tableName}, 1*time.Minute); err != nil {
		return errors.New("wait for table exists failed")
	}
	return nil
}
