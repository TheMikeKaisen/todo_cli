package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	todo "github.com/TheMikeKaisen/Todo_Cli"
)

const (
	todoFile = ".todo.json" // '.' at the start -> hidden file
)

func main() {

	// flags
	add := flag.Bool("add", false, "Add a todo!")
	complete := flag.Int("complete", 0, "Completed the task!")
	delete := flag.Int("delete", 0, "Delete the task related to id.")
	list := flag.Bool("list", false, "list all todos.")

	flag.Parse()

	t := &todo.Todos{}

	err := t.Load(todoFile)
	if err != nil {
		log.Println("Error while loading the todos file.")
		os.Exit(1)
	}

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			log.Println("Error while getting input:", err)
			os.Exit(1)
		}

		t.Add(task)
		err = t.Store(todoFile)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	case *complete > 0:

		err := t.Complete(*complete)
		if err != nil {
			log.Println("Error completing the task")
			os.Exit(1)
		}

		// store the updated value.
		err = t.Store(todoFile)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	case *delete > 0:

		err := t.Delete(*delete)
		if err != nil {
			log.Println("Error deleting the task")
			os.Exit(1)
		}

		// store the updated value.
		err = t.Store(todoFile)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

	case *list:

		t.Print()

	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)
	}
}

// for when user wants to add something to todo!
func getInput(r io.Reader, args ...string) (string, error) {

	if len(args) > 0 {

		// combine individual words into one sentence
		return strings.Join(args, " "), nil
	}

	Scanner := bufio.NewScanner(r)
	Scanner.Scan()

	if err := Scanner.Err(); err != nil {
		return "", err
	}

	text := Scanner.Text()
	if len(text) == 0 {
		return "", errors.New("todo cannot be empty")
	}
	return text, nil
}
