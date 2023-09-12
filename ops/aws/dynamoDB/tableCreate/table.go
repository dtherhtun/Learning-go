package tableCreate

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

func CreateTable(name *string) error {

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("NAME"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("NAME"),
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName:   name,
		BillingMode: types.BillingModePayPerRequest,
	}

	_, err := Client.CreateTable(context.TODO(), input)
	if err != nil {
		return errors.New("error creating table, " + *name + " - " + err.Error())
	}
	return nil
}

func Wait(tableName *string) error {
	waiter := dynamodb.NewTableExistsWaiter(Client)
	err := waiter.Wait(context.TODO(), &dynamodb.DescribeTableInput{
		TableName: tableName}, 1*time.Minute)
	if err != nil {
		return errors.New("Wait for table exists failed")
	}
	return nil
}
