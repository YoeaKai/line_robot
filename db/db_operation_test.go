// db_operation_test verifies the functions of db_operation.go.

package db_operation

import (
	"context"
	"testing"
	"time"
)

func TestInsertMessageToDB(t *testing.T) {
	ctx := context.Background()

	// Testing file.
	document := UserMessage{
		UserId:    "userID1",
		Timestamp: time.Now(),
		Message: Message{
			MessageType: "message'",
			Text:        "Test message",
		},
	}

	// Test insertMessageToDB.
	if err := InsertMessageToDB(ctx, document, "messageId"); err != nil {
		t.Errorf("Failed to run test: %v", err)
	} else {
		t.Log("success")
	}
}
