package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	if index < 0 || index > len(*t) {
		return errors.New("invalid index")
	}

	(*t)[index-1].CompletedAt = time.Now()
	(*t)[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t
	if index < 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...)
	return nil
}

func (t *Todos) Load(filename string) error {

	// read the file
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Println("error while loading the file.")
		return err
	}

	// if file is empty, return err
	if len(file) == 0 {
		return errors.New("the file is empty")
	}

	// unmarshall the json file into original t(Todos list)
	err = json.Unmarshal(file, t)
	if err != nil {
		log.Println("Error while unmarshaling the data")
		return err
	}

	return nil

}

func (t *Todos) Store(filename string) error {

	// marshall the todos
	data, err := json.Marshal(t)
	if err != nil {
		log.Println("Error while marshaling todos")
		return err
	}

	// writes the data to the filename specified. If a file with the given filename does not exist, then it will create one for you!
	// 0644 represents the permission of reading and writing the file to the owner of the operating system.

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Println("Error while writing the file!")
		return err
	}

	return nil
}

func (t *Todos) Print() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		task := blue(item.Task)
		done := blue("no")
		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green("yes")
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
}

