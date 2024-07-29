package main

import (
	"errors"
	"fmt"
	"l8GorutinesChannels/taskManager"
	"time"
)

func main() {

	tasks := []func() error{
		func() error {
			fmt.Println("1 start")
			time.Sleep(1 * time.Second)
			fmt.Println("1 end 1 sec")
			return nil
		},
		func() error {
			fmt.Println("2 start")
			time.Sleep(5 * time.Second)
			fmt.Println("2 end 5 sec")
			return nil
		},
		func() error {
			fmt.Println("3 start")
			time.Sleep(1 * time.Second)
			fmt.Println("3 end 1 sec")
			return nil
		},
		func() error {
			fmt.Println("4 start")
			time.Sleep(6 * time.Second)
			fmt.Println("4 end 6 sec")
			return nil
		},
		func() error {
			fmt.Println("5 start")
			time.Sleep(1 * time.Second)
			fmt.Println("5 end 1 sec")
			return nil
		},
	}

	err1 := taskManager.TaskManager(tasks, 2, 3)
	if err1 != nil {
		fmt.Printf("Finished with error: %v\n", err1)
	} else {
		fmt.Println("All tasks finished")
	}

	tasks2 := []func() error{
		func() error {
			fmt.Println("1 start")
			time.Sleep(1 * time.Second)
			fmt.Println("1 end 1 sec")
			return errors.New("fatal error 1")
		},
		func() error {
			fmt.Println("2 start")
			time.Sleep(5 * time.Second)
			fmt.Println("2 end 5 sec")
			return errors.New("fatal error 2")
		},
		func() error {
			fmt.Println("3 start")
			time.Sleep(1 * time.Second)
			fmt.Println("3 end 1 sec")
			return errors.New("fatal error 3")
		},
		func() error {
			fmt.Println("4 start")
			time.Sleep(6 * time.Second)
			fmt.Println("4 end 6 sec")
			return errors.New("fatal error 4")
		},
		func() error {
			fmt.Println("5 start")
			time.Sleep(1 * time.Second)
			fmt.Println("5 end 1 sec")
			return errors.New("fatal error 5")
		},
	}
	err2 := taskManager.TaskManager(tasks2, 2, 3)
	if err2 != nil {
		fmt.Printf("Finished with error: %v\n", err2)
	} else {
		fmt.Println("All tasks finished")
	}

}

/*Write a function to execute N jobs in parallel (i.e. in N parallel goroutines).
The function takes as input:
- slice with jobs `[]func() error`;
- number of jobs that can be executed in parallel (`N`);
- the maximum number of errors after which processing should be suspended.*/
