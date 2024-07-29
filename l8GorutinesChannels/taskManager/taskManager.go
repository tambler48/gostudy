package taskManager

import (
	"errors"
	"sync"
)

func TaskManager(tasks []func() error, nParallelJobs int, nErrorsToStop int) error {
	var tasksChannel = make(chan func() error, len(tasks))
	var errorsChannel = make(chan error, len(tasks))
	var doneChannel = make(chan bool)
	var wg sync.WaitGroup

	for _, task := range tasks {
		tasksChannel <- task
	}
	close(tasksChannel)

	for i := 0; i < nParallelJobs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasksChannel {
				select {
				case <-doneChannel:
					return
				default:
					err := task()
					if err != nil {
						errorsChannel <- err
					}
				}
			}
		}()
	}

	go func() {
		errorsCount := 0
		for range errorsChannel {
			errorsCount++
			if errorsCount >= nErrorsToStop {
				close(doneChannel)
				return
			}
		}
	}()

	wg.Wait()
	close(errorsChannel)

	errorCount := len(errorsChannel)
	if errorCount >= nErrorsToStop {
		return errors.New("max count of errors")
	}

	return nil
}
