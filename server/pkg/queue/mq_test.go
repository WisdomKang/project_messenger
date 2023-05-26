package queue_test

import (
	"messenger_server/pkg/queue"
	"testing"
)

func TestPublish(t *testing.T) {
	queue.PubMessage()

	queue.SubMessage()
}
