package taskManager

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestTaskManager(t *testing.T) {
	cases := []struct {
		input struct {
			tasks                        []func() error
			nParallelJobs, nErrorsToStop int
		}
		want error
	}{
		{
			input: struct {
				tasks                        []func() error
				nParallelJobs, nErrorsToStop int
			}{
				tasks: []func() error{
					func() error {
						time.Sleep(1 * time.Second)
						return nil
					},
					func() error {
						time.Sleep(5 * time.Second)
						return nil
					},
					func() error {
						time.Sleep(1 * time.Second)
						return nil
					},
					func() error {
						time.Sleep(6 * time.Second)
						return nil
					},
					func() error {
						time.Sleep(1 * time.Second)
						return nil
					},
				},
				nParallelJobs: 2,
				nErrorsToStop: 1,
			},
			want: nil,
		},
		{
			input: struct {
				tasks                        []func() error
				nParallelJobs, nErrorsToStop int
			}{
				tasks: []func() error{
					func() error {
						time.Sleep(1 * time.Second)
						return errors.New("fatal error 1")
					},
					func() error {
						time.Sleep(5 * time.Second)
						return errors.New("fatal error 2")
					},
					func() error {
						time.Sleep(1 * time.Second)
						return errors.New("fatal error 3")
					},
					func() error {
						time.Sleep(6 * time.Second)
						return errors.New("fatal error 4")
					},
					func() error {
						time.Sleep(1 * time.Second)
						return errors.New("fatal error 5")
					},
				},
				nParallelJobs: 2,
				nErrorsToStop: 1,
			},
			want: errors.New("max count of errors"),
		},
	}

	for _, c := range cases {
		got := TaskManager(c.input.tasks, c.input.nParallelJobs, c.input.nErrorsToStop)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got: %v, want: %v", got, c.want)
		}
	}

}
