package runtime

import (
	"fmt"
	"strings"

	"webscrapper/pkg/runtimectrl"
)

const numOfIndentationSpace = 4

type Runtime[T runtimectrl.RuntimeController] interface {
	After(fn RuntimeContextCallback[T])
	AfterEach(fn RuntimeContextCallback[T])
	Before(fn RuntimeContextCallback[T])
	BeforeEach(fn RuntimeContextCallback[T])
	Context(name string, fn RuntimeCallback[T])
	Do(name string, fn RuntimeContextCallback[T])
}

type RuntimeCallback[T runtimectrl.RuntimeController] func(run Runtime[T])

type RuntimeContextCallback[T runtimectrl.RuntimeController] func(ctx *Context[T])

type runtime[T runtimectrl.RuntimeController] struct {
	zone *runtimeZone[T]
}

func newRuntime[T runtimectrl.RuntimeController](space *runtimeZone[T]) *runtime[T] {
	return &runtime[T]{space}
}

func (run *runtime[T]) After(fn RuntimeContextCallback[T]) {
	run.zone.after = append(run.zone.after, fn)
}

func (run *runtime[T]) AfterEach(fn RuntimeContextCallback[T]) {
	run.zone.afterEach = append(run.zone.afterEach, fn)
}

func (run *runtime[T]) Before(fn RuntimeContextCallback[T]) {
	run.zone.before = append(run.zone.before, fn)
}

func (run *runtime[T]) BeforeEach(fn RuntimeContextCallback[T]) {
	run.zone.beforeEach = append(run.zone.beforeEach, fn)
}

func (run *runtime[T]) Context(name string, fn RuntimeCallback[T]) {
	run.zone.tasks = append(run.zone.tasks, func(ctx *Context[T]) {
		newSection := newSection(name, run.zone.depth+1, ctx)
		newZone := newZone(newSection, run.zone.runtimeSpace)
		newRun := newRuntime(newZone)
		fn(newRun)
		newZone.run(run)
	})
}

func (run *runtime[T]) Do(name string, fn RuntimeContextCallback[T]) {
	run.zone.tasks = append(run.zone.tasks, func(ctx *Context[T]) {
		newSection := newSection(name, run.zone.depth, ctx)
		task := newTask(newSection, run.zone.runtimeSpace, fn)
		task.run(run)
	})
}

func (run *runtime[T]) print(value string, indentation int) {
	spaces := strings.Repeat(" ", indentation*numOfIndentationSpace)
	fmt.Printf("%s%s\n", spaces, value)
}

type runtimeSection[T runtimectrl.RuntimeController] struct {
	name  string
	depth int
	ctx   *Context[T]
}

func newSection[T runtimectrl.RuntimeController](name string, depth int, ctx *Context[T]) *runtimeSection[T] {
	return &runtimeSection[T]{name: name, depth: depth, ctx: ctx}
}

type runtimeSpace[T runtimectrl.RuntimeController] struct {
	before     []RuntimeContextCallback[T]
	after      []RuntimeContextCallback[T]
	beforeEach []RuntimeContextCallback[T]
	afterEach  []RuntimeContextCallback[T]
	tasks      []RuntimeContextCallback[T]
}

func newSpace[T runtimectrl.RuntimeController]() *runtimeSpace[T] {
	return &runtimeSpace[T]{}
}

type runtimeZone[T runtimectrl.RuntimeController] struct {
	*runtimeSection[T]
	*runtimeSpace[T]
}

func newZone[T runtimectrl.RuntimeController](section *runtimeSection[T], space *runtimeSpace[T]) *runtimeZone[T] {
	return &runtimeZone[T]{
		runtimeSection: section,
		runtimeSpace: &runtimeSpace[T]{
			beforeEach: append([]RuntimeContextCallback[T]{}, space.beforeEach...),
			afterEach:  append([]RuntimeContextCallback[T]{}, space.afterEach...),
		},
	}
}

func (zone *runtimeZone[T]) run(run *runtime[T]) {
	run.print(zone.name, zone.depth)
	ctx := zone.ctx.clone()
	for _, before := range zone.before {
		run.print("Before", zone.depth)
		before(ctx)
	}
	for _, task := range zone.tasks {
		task(ctx)
	}
	for _, after := range zone.after {
		run.print("After", zone.depth)
		after(ctx)
	}
}

type runtimeTask[T runtimectrl.RuntimeController] struct {
	*runtimeSection[T]
	*runtimeSpace[T]
	fn RuntimeContextCallback[T]
}

func newTask[T runtimectrl.RuntimeController](section *runtimeSection[T], space *runtimeSpace[T], fn RuntimeContextCallback[T]) *runtimeTask[T] {
	return &runtimeTask[T]{
		fn:             fn,
		runtimeSection: section,
		runtimeSpace: &runtimeSpace[T]{
			beforeEach: space.beforeEach,
			afterEach:  space.afterEach,
		},
	}
}

func (task *runtimeTask[T]) run(run *runtime[T]) {
	ctx := task.ctx.clone()
	for _, before := range task.beforeEach {
		run.print("Before Each", task.depth)
		before(ctx)
	}
	run.print("Task "+task.name, task.depth)
	task.fn(ctx)
	for _, after := range task.afterEach {
		run.print("After Each", task.depth)
		after(ctx)
	}
}

type Context[T runtimectrl.RuntimeController] struct {
	C T
}

func newContext[T runtimectrl.RuntimeController](c T) *Context[T] {
	return &Context[T]{c}
}

func (ctx *Context[T]) clone() *Context[T] {
	return &Context[T]{ctx.C.Clone().(T)}
}

type RuntimeWrapper[T runtimectrl.RuntimeController] struct {
	ctx *Context[T]
}

func New[T runtimectrl.RuntimeController](c T) *RuntimeWrapper[T] {
	return &RuntimeWrapper[T]{newContext(c)}
}

func (runw *RuntimeWrapper[T]) Run(name string, fn RuntimeCallback[T]) {
	section := newSection(name, 0, runw.ctx)
	space := newSpace[T]()
	zone := newZone(section, space)
	run := newRuntime(zone)
	fn(run)
	zone.run(run)
}
