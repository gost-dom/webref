package events_test

import (
	"testing"

	"github.com/gost-dom/webref/events"
	"github.com/onsi/gomega"
)

func TestElementEvents(t *testing.T) {
	expect := gomega.NewWithT(t).Expect
	e, err := events.Load("uievents")
	expect(err).ToNot(gomega.HaveOccurred())
	result := e.EventsForType("Element")
	expect(result).To(gomega.ContainElement(events.Event{
		Type:      "click",
		Interface: "PointerEvent",
		Options: map[events.EventOption]bool{
			"bubbles":    true,
			"cancelable": true,
		},
	}))
}
