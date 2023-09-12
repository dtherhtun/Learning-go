package tableCreate

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type BarJoke struct {
	Name     string  `json:"name" dynamodbav:"NAME"`
	Rating   float64 `json:"rating"`
	Headline string  `json:"headline"`
	Content  string  `json:"content"`
}

func FillTable(table *string) error {
	items, err := ReadFile()
	if err != nil {
		return err
	}

	for _, record := range items {
		av, err := attributevalue.MarshalMap(record)
		if err != nil {
			return fmt.Errorf("failed to marshal item: %w", err)
		}

		input := &dynamodb.PutItemInput{
			TableName: table,
			Item:      av,
		}

		_, err = Client.PutItem(context.TODO(), input)
		if err != nil {
			return fmt.Errorf("failed to write item to DynamoDB: %w", err)
		}
	}
	return nil
}

func ReadFile() ([]BarJoke, error) {
	fileBytes, err := os.ReadFile("items.json")
	if err != nil {
		fmt.Println("failed to read JSON file:", err)
		return nil, err
	}

	var items []BarJoke
	err = json.Unmarshal(fileBytes, &items)
	if err != nil {
		fmt.Println("failed to unmarshal JSON:", err)
	}
	return items, nil
}
