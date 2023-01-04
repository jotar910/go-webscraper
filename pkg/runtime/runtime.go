package runtime

import (
	"fmt"
	"strings"
	"webscrapper/pkg/controller"
)

const numOfIndentationSpace = 4

type Runtime interface {
	After(fn RuntimeContextCallback)
	AfterEach(fn RuntimeContextCallback)
	Before(fn RuntimeContextCallback)
	BeforeEach(fn RuntimeContextCallback)
	Context(name string, fn RuntimeCallback)
	Do(name string, fn RuntimeContextCallback)
}

type RuntimeCallback func(run Runtime)

type RuntimeContextCallback func(ctx *Context)

type runtime struct {
	zone *runtimeZone
}

func newRuntime(space *runtimeZone) *runtime {
	return &runtime{space}
}

func (run *runtime) After(fn RuntimeContextCallback) {
	run.zone.after = append(run.zone.after, fn)
}

func (run *runtime) AfterEach(fn RuntimeContextCallback) {
	run.zone.afterEach = append(run.zone.afterEach, fn)
}

func (run *runtime) Before(fn RuntimeContextCallback) {
	run.zone.before = append(run.zone.before, fn)
}

func (run *runtime) BeforeEach(fn RuntimeContextCallback) {
	run.zone.beforeEach = append(run.zone.beforeEach, fn)
}

func (run *runtime) Context(name string, fn RuntimeCallback) {
	run.zone.tasks = append(run.zone.tasks, func(ctx *Context) {
		newSection := newSection(name, run.zone.depth+1, ctx)
		newZone := newZone(newSection, run.zone.runtimeSpace)
		newRun := newRuntime(newZone)
		fn(newRun)
		newZone.run(run)
	})
}

func (run *runtime) Do(name string, fn RuntimeContextCallback) {
	run.zone.tasks = append(run.zone.tasks, func(ctx *Context) {
		newSection := newSection(name, run.zone.depth, ctx)
		task := newTask(newSection, run.zone.runtimeSpace, fn)
		task.run(run)
	})
}

func (run *runtime) print(value string, indentation int) {
	spaces := strings.Repeat(" ", indentation*numOfIndentationSpace)
	fmt.Printf("%s%s\n", spaces, value)
}

type runtimeSection struct {
	name  string
	depth int
	ctx   *Context
}

func newSection(name string, depth int, ctx *Context) *runtimeSection {
	return &runtimeSection{name: name, depth: depth, ctx: ctx}
}

type runtimeSpace struct {
	before     []RuntimeContextCallback
	after      []RuntimeContextCallback
	beforeEach []RuntimeContextCallback
	afterEach  []RuntimeContextCallback
	tasks      []RuntimeContextCallback
}

func newSpace() *runtimeSpace {
	return &runtimeSpace{}
}

type runtimeZone struct {
	*runtimeSection
	*runtimeSpace
}

func newZone(section *runtimeSection, space *runtimeSpace) *runtimeZone {
	return &runtimeZone{
		runtimeSection: section,
		runtimeSpace: &runtimeSpace{
			beforeEach: append([]RuntimeContextCallback{}, space.beforeEach...),
			afterEach:  append([]RuntimeContextCallback{}, space.afterEach...),
		},
	}
}

func (zone *runtimeZone) run(run *runtime) {
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

type runtimeTask struct {
	*runtimeSection
	*runtimeSpace
	fn RuntimeContextCallback
}

func newTask(section *runtimeSection, space *runtimeSpace, fn RuntimeContextCallback) *runtimeTask {
	return &runtimeTask{
		fn:             fn,
		runtimeSection: section,
		runtimeSpace: &runtimeSpace{
			beforeEach: space.beforeEach,
			afterEach:  space.afterEach,
		},
	}
}

func (task *runtimeTask) run(run *runtime) {
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

type Context struct {
	C controller.Controller
}

func newContext(c controller.Controller) *Context {
	return &Context{c}
}

func (ctx *Context) clone() *Context {
	return &Context{ctx.C.Clone()}
}

type RuntimeWrapper struct {
	ctx *Context
}

func New(c controller.Controller) *RuntimeWrapper {
	return &RuntimeWrapper{newContext(c)}
}

func (runw *RuntimeWrapper) Run(name string, fn RuntimeCallback) {
	section := newSection(name, 0, runw.ctx)
	space := newSpace()
	zone := newZone(section, space)
	run := newRuntime(zone)
	fn(run)
	zone.run(run)
}
