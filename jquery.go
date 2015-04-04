package angularjs

import "github.com/gopherjs/gopherjs/js"

type JQueryElement struct{ *js.Object }

func ElementById(id string) *JQueryElement {
	return &JQueryElement{js.Global.Get("angular").Call("element", js.Global.Get("document").Call("getElementById", id))}
}

func (e *JQueryElement) Prop(name string) *js.Object {
	return e.Call("prop", name)
}

func (e *JQueryElement) SetProp(name, value interface{}) {
	e.Call("prop", name, value)
}

func (e *JQueryElement) On(events string, handler func(*Event)) {
	e.Call("on", events, func(e *js.Object) {
		handler(&Event{Object: e})
	})
}

func (e *JQueryElement) Val() *js.Object {
	return e.Call("val")
}

func (e *JQueryElement) SetVal(value interface{}) {
	e.Call("val", value)
}