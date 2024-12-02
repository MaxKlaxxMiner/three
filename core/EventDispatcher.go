package core

import "slices"

/**
 * ported from:
 * https://github.com/mrdoob/eventdispatcher.js/
 */

type Event struct {
	TypeName string
	Target   any
}

type EventDispatcher struct {
	listeners map[string][]*func(event Event)
	This      any
}

func NewEventDispatcher() *EventDispatcher {
	this := new(EventDispatcher)
	this.This = this
	return this
}

func (e *EventDispatcher) AddListener(typeName string, listener *func(event Event)) {
	if e.listeners == nil {
		e.listeners = make(map[string][]*func(event Event))
	}
	if !e.HasEventListener(typeName, listener) {
		e.listeners[typeName] = append(e.listeners[typeName], listener)
	}
}

func (e *EventDispatcher) HasEventListener(typeName string, listener *func(event Event)) bool {
	if e.listeners == nil {
		return false
	}
	return slices.Index(e.listeners[typeName], listener) >= 0
}

func (e *EventDispatcher) RemoveListener(typeName string, listener *func(event Event)) {
	if e.listeners == nil {
		return
	}
	if i := slices.Index(e.listeners[typeName], listener); i >= 0 {
		e.listeners[typeName] = slices.Delete(e.listeners[typeName], i, i+1)
	}
}

func (e *EventDispatcher) DispatchEvent(typeName string) {
	if e.listeners == nil {
		return
	}
	if array := e.listeners[typeName]; array != nil {
		array = slices.Clone(array) // Make a copy, in case listeners are removed while iterating.
		event := Event{typeName, e.This}
		for _, run := range array {
			(*run)(event)
		}
	}
}
