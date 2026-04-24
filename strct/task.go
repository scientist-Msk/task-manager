package strct

import (
	"time"
)

type Task struct {
	heading string
	text    string
	timeCrt string
}
type TaskD struct {
	HeadingD   string
	textD      string
	timeCrtD   string
	timeCmpltD string
}

func (f *Task) AddHeading(head string) {
	if head != "" {
		f.heading = head
	}
}

func (f *Task) AddText(txt string) {
	if txt != "" {
		f.text = txt
	}
}
func NewTask(Heading, Text string) Task {
	if Heading == "" || Text == "" {
		return Task{}
	}
	TimeCreate := time.Now().Format("2006-01-02 15:04")
	return Task{
		heading: Heading,
		text:    Text,
		timeCrt: TimeCreate,
	}
}

func DeleteTaskByHeading(tasks map[Task]bool, heading string) bool {
	for key := range tasks {
		if key.heading == heading {
			delete(tasks, key)
			return true
		}
	}
	return false
}

func TaskComplete(tasks map[Task]bool, heading string) TaskD {
	for key := range tasks {
		if key.heading == heading {
			HDNGD := heading
			TXTD := key.text
			TCRTD := key.timeCrt
			TCMPLTD := time.Now().Format("2006-01-02 15:04")
			delete(tasks, key)
			return TaskD{
				HeadingD:   HDNGD,
				textD:      TXTD,
				timeCrtD:   TCRTD,
				timeCmpltD: TCMPLTD,
			}
		}

	}
	return TaskD{}
}
