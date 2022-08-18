package driver

import (
	"encoding/json"
	"time"

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
	_ = dispatchConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if _, err = dispatchConn.WriteMessages(
		kafka.Message{Value: msg},
	); err != nil {
		log.Printf("SendDispatchEventMsg send msg error(%v)", err)
		return err
	}
	return nil
}
