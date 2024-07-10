package pkg

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

type item struct {
	TaskName string
	Done     bool
}

type Todo struct {
	array []item
}

func (t *Todo) Add(task string) {
	addItem := &item{TaskName: task, Done: false}
	t.array = append(t.array, *addItem)
}

func (t *Todo) Delete(index int) {
	array1 := t.array[:index-1]

	array2 := t.array[index:]

	array1 = append(array1, array2...)

	t.array = array1
}

func (t *Todo) UpdateTask(index int, task string) {
	t.array[index-1].TaskName = task
	fmt.Println(task)
	fmt.Println(t.array)
}

func (t *Todo) Done(index int) {
	t.array[index-1].Done = true
}

func (t Todo) ToPrint() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "TASKS"},
			{Align: simpletable.AlignCenter, Text: "DONE?"},
		},
	}
	if len(t.array) == 0 {
		fmt.Println("Empty")
		return
	}

	for i, v := range t.array {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", i+1)},
			{Text: v.TaskName},
			{Text: fmt.Sprintf("%v ", v.Done)},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	fmt.Println(table.String())
}
