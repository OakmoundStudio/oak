package event

import (
	"sync"
)

// TriggerBack is a version of Trigger which returns a channel that
// informs on when all bindables have been called and returned from
// the input event. It is dangerous to use this unless you have a
// very good idea how things will synchronize, as if a triggered
// bindable itself makes a TriggerBack call, this will cause the engine to freeze,
// as the function will never end because the first TriggerBack has control of
// the lock for the event bus, and the first TriggerBack won't give up that lock
// until the function ends.
//
// This inherently means that when you call Trigger, the event will almost
// almost never be immediately triggered but rather will be triggered sometime
// soon in the future.
//
// TriggerBack is right now used by the primary logic loop to dictate logical
// framerate, so EnterFrame events are called through TriggerBack.
func (eb *Bus) TriggerBack(eventName string, data interface{}) chan struct{} {
	ch := make(chan struct{})
	go func(ch chan struct{}, eb *Bus, eventName string, data interface{}) {
		eb.trigger(eventName, data)
		close(ch)
	}(ch, eb, eventName, data)
	return ch
}

// Trigger will scan through the event bus and call all bindables found attached
// to the given event, with the passed in data.
func (eb *Bus) Trigger(eventName string, data interface{}) {
	go func(eb *Bus, eventName string, data interface{}) {
		eb.trigger(eventName, data)
	}(eb, eventName, data)
}

func (eb *Bus) TriggerCIDBack(cid CID, eventName string, data interface{}) chan struct{} {
	ch := make(chan struct{})
	go func() {
		eb.mutex.RLock()
		if idMap, ok := eb.bindingMap[eventName]; ok {
			if bs, ok := idMap[cid]; ok {
				eb.triggerDefault(bs.sl, cid, eventName, data)
			}
		}
		eb.mutex.RUnlock()
		close(ch)
	}()
	return ch
}

func (eb *Bus) trigger(eventName string, data interface{}) {
	eb.mutex.RLock()
	for id, bs := range eb.bindingMap[eventName] {
		if bs != nil {
			eb.triggerDefault(bs.sl, id, eventName, data)
		}
	}
	eb.mutex.RUnlock()
}

func (eb *Bus) triggerDefault(sl []Bindable, id CID, eventName string, data interface{}) {
	prog := &sync.WaitGroup{}
	prog.Add(len(sl))
	for i, bnd := range sl {
		if bnd == nil {
			prog.Done()
			continue
		}
		go func(bnd Bindable, id CID, eventName string, data interface{}, prog *sync.WaitGroup, index int) {
			eb.handleBindable(bnd, id, data, index, eventName)
			prog.Done()
		}(bnd, id, eventName, data, prog, i)
	}
	prog.Wait()
}

func (eb *Bus) handleBindable(bnd Bindable, id CID, data interface{}, index int, eventName string) {
	if id == 0 || eb.callerMap.HasEntity(id) {
		response := bnd(id, data)
		switch response {
		case UnbindEvent:
			UnbindAll(Event{
				Name:     eventName,
				CallerID: id,
			})
		case UnbindSingle:
			bnd := binding{
				Event: Event{
					Name:     eventName,
					CallerID: id,
				},
				index: index,
			}
			bnd.unbind(eb)
		}
	}
}
