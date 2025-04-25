package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"taskMangementService/model"

	"github.com/segmentio/kafka-go"
)

type User struct {
    ID   uint   `json:"id"`
    Name string `json:"name"`
	Email string `json:"email"`
	ContactNumber string `json:"contactNumber"`
}

type TaskNotification struct {
    TaskID      uint   `json:"task_id"`
    TaskTitle   string `json:"task_title"`
    UserID      uint   `json:"user_id"`
    UserName    string `json:"user_name"`
    Email       string `json:"email"`
    PhoneNumber string `json:"phone_number"`
}

type UserClient interface {
    ValidateAndGetUserByID(id uint) (*User, error)
	notifyUser(user *User, task *model.Task)
}

func ValidateAndGetUserByID(userID uint) (*User, error) {

	// Calling other microservice user-service which handles user auth and details
    resp, err := http.Get(fmt.Sprintf("http://localhost:8082/users/%d", userID))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("user not found")
    }

    var user User
    err = json.NewDecoder(resp.Body).Decode(&user)
    return &user, err
}

func notifyUser(user *User, task *model.Task) {

	// Notifyting users through email and sms with message queue like kafka
    notification := TaskNotification{
        TaskID:      task.ID,
        TaskTitle:   task.Title,
        UserID:      user.ID,
        UserName:    user.Name,
        Email:       user.Email,        
        PhoneNumber: user.ContactNumber,
    }

	data, err := json.Marshal(notification)
    if err != nil {
        log.Println("Error marshalling notification:", err)
        return
    }
	
	 writer := kafka.Writer{
        Addr:     kafka.TCP("localhost:9092"), 
        Topic:    "task-notifications",
        Balancer: &kafka.LeastBytes{},
    }

	err = writer.WriteMessages(context.Background(),
        kafka.Message{
            Key:   []byte(fmt.Sprintf("%d", task.ID)),
            Value: data,
        },
    )

	if err != nil {
        log.Println("Error sending Kafka message:", err)
    } else {
        log.Println("Notification sent to Kafka for task:", task.ID)
    }

    writer.Close()
}