package eventbus_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/latavin243/goutils/eventbus"
)

func TestBus(t *testing.T) {
	eventBus := eventbus.NewEventBus[string]()
	topic := "stringEvent"

	stringSubscriber := eventbus.Subscriber[string]{
		Name: "string sub",
		Notify: func(events []string) {
			fmt.Printf("Received string event: %+v\n", events)
		},
	}

	eventBus.Subscribe(topic, stringSubscriber)

	eventBus.Publish(topic, []string{"Hello, Event Bus!"})

	eventBus.Unsubscribe(topic, "string sub")

	eventBus.Publish(topic, []string{"This event won't be received."})

	time.Sleep(time.Second)
}
