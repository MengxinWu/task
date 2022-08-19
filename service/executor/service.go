package executor

import (
	"task/models"
)

var executeEventCh = make(chan *models.ExecuteEvent, 100000)

// InitExecuteEvent init execute event.
func InitExecuteEvent() {
	for {
		select {
		case event := <-executeEventCh:
			var err error
			if err = run(event); err != nil {
				handleProcessError(event, err)
			}
		}
	}
}

// ReceiveExecuteEvent receive execute event.
func ReceiveExecuteEvent(event *models.ExecuteEvent) {
	executeEventCh <- event
	return
}
