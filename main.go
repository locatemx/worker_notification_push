package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodb_types "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	sqs_types "github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type Notification struct {
	Consumer string `json:"consumer"`
	Filter   string `json:"filter"`
	FilterID string `json:"filter_id"`
	Content  struct {
		Headings NotificationContents `json:"headings"`
		Contents NotificationContents `json:"contents"`
		Buttons  []NotificationButton `json:"buttons"`
	} `json:"content"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
	TTL       int64  `json:"ttl"`
}

var (
	sqs_client    *sqs.Client
	dynamo_client *dynamodb.Client
	queue_url     = "http://192.168.3.200:4566/913267050402/worker_notifications_push"
	table_name    = "notifications"
	api_url       = "http://localhost:8080/sendPushNotification"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Error loading AWS configuration: %v", err)
	}

	// Inicializar clientes de SQS y DynamoDB
	sqs_client = sqs.NewFromConfig(cfg, func(o *sqs.Options) {
		o.BaseEndpoint = aws.String("http://192.168.3.200:4566")
	})
	dynamo_client = dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String("http://192.168.3.200:4566")
	})

	log.Println("Waiting for push notifications...")
	for {
		receive_messages()
		time.Sleep(5 * time.Second)
	}
}

// Función para recibir mensajes de SQS
func receive_messages() {
	output, err := sqs_client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queue_url),
		MaxNumberOfMessages: 10,
		WaitTimeSeconds:     10,
	})

	if err != nil {
		log.Printf("Error receiving messages: %v", err)
		return
	}

	for _, message := range output.Messages {
		log.Printf("Received push notification message: %s", *message.Body)

		// Extraer el campo "detail" del mensaje recibido
		var event map[string]interface{}
		if err := json.Unmarshal([]byte(*message.Body), &event); err != nil {
			log.Printf("Error parsing message body: %v", err)
			continue
		}

		// Convertir el "detail" en una estructura Notification
		detailJSON, err := json.Marshal(event["detail"])
		if err != nil {
			log.Printf("Error marshalling detail: %v", err)
			continue
		}

		var notification Notification
		if err := json.Unmarshal(detailJSON, &notification); err != nil {
			log.Printf("Error parsing detail into Notification struct: %v", err)
			continue
		}

		// Llamada a la API de notificación push (comentado hasta que la API esté lista)
		/*
			if err := send_push_notification(notification); err != nil {
				log.Printf("Error sending push notification: %v", err)
				continue
			}
		*/
		var deviceAccounts = []AccountDevice{}
		if err := NotifyAlarm(detailJSON, deviceAccounts); err != nil {
			log.Printf("Error notifying alarm: %v", err)
			continue
		}

		// Actualizar el estado en DynamoDB a 'sent'
		if err := update_notification_status(notification.FilterID, notification.Timestamp, "sent"); err != nil {
			log.Printf("Error updating status in DynamoDB: %v", err)
			continue
		}

		delete_message(message)
		log.Printf("Push notification processed and deleted with FilterID: %v", notification.FilterID)
	}
}

func update_notification_status(filterID string, timestamp int64, newStatus string) error {
	updateInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(table_name),
		Key: map[string]dynamodb_types.AttributeValue{
			"filter_id": &dynamodb_types.AttributeValueMemberS{Value: filterID},
			"timestamp": &dynamodb_types.AttributeValueMemberN{Value: fmt.Sprintf("%d", timestamp)},
		},
		UpdateExpression: aws.String("SET #s = :newStatus"),
		ExpressionAttributeNames: map[string]string{
			"#s":  "status",
			"#ts": "timestamp", // Alias para evitar el uso de la palabra reservada
		},
		ExpressionAttributeValues: map[string]dynamodb_types.AttributeValue{
			":newStatus": &dynamodb_types.AttributeValueMemberS{Value: newStatus},
		},
		ConditionExpression: aws.String("attribute_exists(filter_id) AND attribute_exists(#ts)"),
	}

	_, err := dynamo_client.UpdateItem(context.TODO(), updateInput)
	if err != nil {
		log.Printf("Error updating status in DynamoDB: %v", err)
	}
	return err
}

// Función para eliminar el mensaje de SQS
func delete_message(message sqs_types.Message) {
	_, err := sqs_client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queue_url),
		ReceiptHandle: message.ReceiptHandle,
	})
	if err != nil {
		log.Printf("Error deleting message: %v", err)
	}
}
