package todo_test

import (
	"os"
	"testing"

	"github.com/halladj/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskname := "My Task"
	l.Add(taskname)

	if l[0].Task != taskname {
		t.Errorf(
			"Expected %q, got %q instead",
			taskname,
			l[0].Task,
		)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	taskname := "My Task"
	l.Add(taskname)

	if l[0].Task != taskname {
		t.Errorf(
			"Expected %q, got %q instead",
			taskname,
			l[0].Task,
		)
	}

	if l[0].Done {
		t.Errorf(
			"New Task Should Not Be Completed.",
		)
	}

	l.Complete(1)
	if !l[0].Done {
		t.Errorf(
			"New Task Should Be Completed.",
		)
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New Task 0",
		"New Task 1",
		"New Task 2",
	}

	for _, v := range tasks {
		l.Add(v)
	}

	if l[0].Task != tasks[0] {
		t.Errorf(
			"Expected %q, got %q instead",
			tasks[0],
			l[0].Task,
		)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf(
			"Expected list length %d, got %d instead.",
			2,
			len(l),
		)
	}

	// l should be, l -> [ "New Task0", "New Task 2" ] and tasks -> ["New Task0", "New Task 1", "New Task 2"]
	if l[1].Task != tasks[2] {
		t.Errorf(
			"Expected %q, got %q instead",
			tasks[2],
			l[1].Task,
		)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskname := "My Task"
	l1.Add(taskname)

	if l1[0].Task != taskname {
		t.Errorf(
			"Expected %q, got %q instead.",
			taskname,
			l1[0].Task,
		)
	}

	tf, err := os.CreateTemp("", "")
	defer os.Remove(tf.Name())
	if err != nil {
		t.Fatalf(
			"Error Creating temp file: %s", err,
		)
	}

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf(
			"Error saving list to file: %s", err,
		)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf(
			"Error getting list from file: %s", err,
		)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf(
			"Task %q should match %q task",
			l1[0].Task,
			l2[0].Task,
		)
	}
}
