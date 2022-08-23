package driver

import (
	"context"
	"encoding/json"

	"task/models"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

// SendDispatchEventMsg send dispatch event msg.
func SendDispatchEventMsg(event *models.DispatchEvent) error {
	var (
		msg []byte
		err error
	)
	msg, _ = json.Marshal(event)
	log.Printf("SendDispatchEventMsg msg: %s", string(msg))
	if err = dispatchWriter.WriteMessages(context.Background(),
		kafka.Message{Value: msg},
	); err != nil {
		log.Printf("SendDispatchEventMsg send msg error(%v)", err)
		return err
	}
	return nil
}

// SendExecuteEventMsg send execute event msg.
func SendExecuteEventMsg(events []*models.ExecuteEvent) error {
	var (
		msg []byte
		err error
	)
	messages := make([]kafka.Message, 0)
	for _, event := range events {
		msg, _ = json.Marshal(event)
		log.Printf("SendExecuteEventMsg msg: %s", string(msg))
		messages = append(messages, kafka.Message{Value: msg})
	}
	if err = executeWriter.WriteMessages(context.Background(), messages...); err != nil {
		log.Printf("SendExecuteEventMsg send msg error(%v)", err)
		return err
	}
	return nil
}
