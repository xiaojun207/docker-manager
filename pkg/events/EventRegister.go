package events

// 事件注册
type EventRegister struct {
	// 监听事件
	Listen map[string][]func(interface{})
}

func NewEventRegister() *EventRegister {
	return &EventRegister{
		Listen: make(map[string][]func(interface{})),
	}
}

func (this *EventRegister) On(event string, fn func(interface{})) {
	this.Listen[event] = append(this.Listen[event], fn)
}
func (this *EventRegister) Emit(event string, data interface{}) {
	for _, fn := range this.Listen[event] {
		fn(data)
	}
}
func (this *EventRegister) Off(event string, fn func(interface{})) {
	delete(this.Listen, event)
}
