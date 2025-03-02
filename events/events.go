package events

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"slices"

	"github.com/gost-dom/webref/internal/specs"
)

type Properties map[string]bool

type Event struct {
	// The event type, such as "click". Corresponds to the `type` property on the event
	Type string
	// Interface is the name of the IDL interface that represents the event.
	// E.g., click events on elements are represented by "PointerEvent" events.
	Interface string
	// Properties represents the properties specified, e.g. 'bubbles' or
	// 'cancelable'. Only properties that are explicitly specified will exist in
	// the list.
	Properties
}

type rawJsonFile struct {
	Events []rawEventData
}

type rawEventDataProps struct {
	Type      string
	Interface string
	Targets   []string
}

type rawEventData struct {
	Props rawEventDataProps
	Properties
}

func (d *rawEventData) UnmarshalJSON(data []byte) error {
	values := make(map[string]interface{})
	d.Properties = make(Properties)
	err1 := json.Unmarshal(data, &d.Props)
	err2 := json.Unmarshal(data, &values)
	for k, v := range values {
		if b, ok := v.(bool); ok {
			d.Properties[k] = b
		}
	}
	return errors.Join(err1, err2)
}

func (d rawEventData) hasTarget(name string) bool {
	return slices.Contains(d.Props.Targets, name)
}

type Events struct {
	events []rawEventData
}

// EventsForType returns all events in the spec that relate to elements of a
// specific type. E.g., for an HTML form, this will contain the "formdata", and
// "submit" events.
func (a Events) EventsForType(tagName string) []Event {
	var result []Event
	for _, raw := range a.events {
		if raw.hasTarget(tagName) {
			result = append(result, Event{
				Type:       raw.Props.Type,
				Interface:  raw.Props.Interface,
				Properties: raw.Properties,
			})
		}
	}
	return result
}

func parseFile(reader io.Reader) (rawJsonFile, error) {
	spec := rawJsonFile{}
	b, err := io.ReadAll(reader)
	if err == nil {
		err = json.Unmarshal(b, &spec)
	}
	// spec.initialize()
	return spec, err
}

func Load(apiName string) (Events, error) {
	file, err := specs.Open(fmt.Sprintf("events/%s.json", apiName))
	defer file.Close()

	if err != nil {
		return Events{}, err
	}
	data, err := parseFile(file)
	fmt.Println("DATA", data)
	return Events{data.Events}, err
}
