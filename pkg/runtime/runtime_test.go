package runtime

import (
	"fmt"
	"testing"
	"webscrapper/pkg/runtimectrl"
)

type controllerMock struct {
}

func (ctrl *controllerMock) Clone() runtimectrl.RuntimeController {
	return ctrl
}

func Test(t *testing.T) {
	New(&controllerMock{}).Run("Run", func(run Runtime[*controllerMock]) {

		run.Before(func(ctx *Context[*controllerMock]) {
			fmt.Println("Before [1.1]")
		})

		run.AfterEach(func(ctx *Context[*controllerMock]) {
			fmt.Println("AfterEach [1.1]")
		})

		run.BeforeEach(func(ctx *Context[*controllerMock]) {
			fmt.Println("BeforeEach [1.1]")
		})

		run.BeforeEach(func(ctx *Context[*controllerMock]) {
			fmt.Println("BeforeEach [1.2]")
		})

		run.Do("Do 1", func(ctx *Context[*controllerMock]) {
			fmt.Println("Do [1.1]")
		})

		run.Do("Do 2", func(ctx *Context[*controllerMock]) {
			fmt.Println("Do [1.2]")
		})

		run.Do("Do 3", func(ctx *Context[*controllerMock]) {
			fmt.Println("Do [1.3]")
		})

		run.Context("Context 3", func(run Runtime[*controllerMock]) {
			fmt.Println("Context [2.1]")

			run.BeforeEach(func(ctx *Context[*controllerMock]) {
				fmt.Println("BeforeEach [2.1]")
			})

			run.Do("Do 4", func(ctx *Context[*controllerMock]) {
				fmt.Println("Do [2.1]")
			})

			run.Do("Do 5", func(ctx *Context[*controllerMock]) {
				fmt.Println("Do [2.2]")
			})

			run.Do("Do 6", func(ctx *Context[*controllerMock]) {
				fmt.Println("Do [2.3]")
			})

			run.BeforeEach(func(ctx *Context[*controllerMock]) {
				fmt.Println("BeforeEach [2.2]")
			})

			run.Before(func(ctx *Context[*controllerMock]) {
				fmt.Println("Before [2.1]")
			})

			run.After(func(ctx *Context[*controllerMock]) {
				fmt.Println("After [2.1]")
			})

		})

		run.After(func(ctx *Context[*controllerMock]) {
			fmt.Println("After [1.1]")
		})

		run.AfterEach(func(ctx *Context[*controllerMock]) {
			fmt.Println("AfterEach [1.2]")
		})

		run.Do("Do 7", func(ctx *Context[*controllerMock]) {
			fmt.Println("Do [1.4]")
		})
	})
}
