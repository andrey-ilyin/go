package safeslice

type SafeSlice interface {
	Append(interface{})
	At(int) interface{}
	Close() []interface{}
	Delete(int)
	Len() int
	Update(int, UpdateFunc)
}
type UpdateFunc func(interface{}) interface{}

type safeSlice chan command

type command struct {
	action  action
	index   int
	value   interface{}
	updater UpdateFunc
	result  chan interface{}
	data    chan []interface{}
}

type action int

type findResult struct {
	value interface{}
}

const (
	add action = iota
	at
	remove
	length
	update
	end
)

func (s safeSlice) Append(value interface{}) {
	s <- command{action: add, value: value}
}

func (s safeSlice) At(index int) interface{} {
	reply := make(chan interface{})
	s <- command{action: at, index: index, result: reply}
	result := (<-reply).(findResult)
	return result.value
}

func (s safeSlice) Delete(index int) {
	s <- command{action: remove, index: index}
}

func (s safeSlice) Len() int {
	reply := make(chan interface{})
	s <- command{action: length, result: reply}
	return (<-reply).(int)
}

func (s safeSlice) Update(index int, updater UpdateFunc) {
	s <- command{action: update, index: index, updater: updater}
}

func (s safeSlice) Close() []interface{} {
	data := make(chan []interface{})
	s <- command{action: end, data: data}
	return <-data
}

func New() safeSlice {
	s := make(safeSlice)
	s.run()
	return s
}

func (s safeSlice) run() {
	go func() {
		data := make([]interface{}, 0)
		for {
			command := <-s
			switch command.action {
			case add:
				data = append(data, command.value)
			case at:
				result := findResult{}
				if command.index < len(data) && command.index >= 0 {
					result.value = data[command.index]
				} else {
					result.value = nil
				}
				command.result <- result
			case remove:
				if command.index < len(data) && command.index >= 0 {
					data = append(data[:command.index], data[command.index+1:]...)
				}
			case length:
				command.result <- len(data)
			case update:
				if command.index < len(data) && command.index >= 0 {
					value := data[command.index]
					data[command.index] = command.updater(value)
				}
			case end:
				close(s)
				command.data <- data
				return
			}
		}
	}()
}
