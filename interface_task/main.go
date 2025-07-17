package main

import "fmt"

//----------- ex1
// type Foo struct {}

// func (f *Foo) A(){}
// func (f *Foo) B(){}
// func (f *Foo) C(){}

// type AB interface {
// 	A()
// 	B()
// }

// type BC interface {
// 	C()
// 	B()
// }

// func main() {
// 	var f AB = &Foo{}
// 	y := f.(BC)
// 	y.A()
// }

// ----------- ex2
// Вопрос: будет ли работать? Что если Print() был бы func (d *Doc) Print()?
// type Printer interface {
// 	Print()
// }

// type Doc struct{}

// func (d Doc) Print() {
// 	fmt.Println("printing")
// }

// func main() {
// 	var p Printer

// 	fmt.Println(p == nil) // true

// 	var d *Doc = &Doc{}
// 	// var d *Doc // nil == true
// 	fmt.Println(d == nil) // false

// 	p = d
// 	p.Print()
// }

// ----------- ex3
// Вопрос: что выведет?
// type Printer interface {
// 	Print()
// }

// type Doc struct{}

// func (d *Doc) Print() {
// 	fmt.Println("doc")
// }

// func main() {
// 	var d *Doc = nil
// 	var p Printer = d
// 	fmt.Println(p == nil) // false
// 	p.Print()             // doc
// }

// ----------- ex4
// Вопрос: будет ли работать вызов a.Bar()?
// type A interface {
// 	Foo()
// }

// type B interface {
// 	A
// 	Bar()
// }

// type Impl struct{}

// func (Impl) Foo() {}
// func (Impl) Bar() {}

// func main() {
// 	var b B = Impl{}
// 	var a A = b
// 	a.Foo()
// 	a.Bar() // ❌ ?
// }

// ----------- ex5
// type Worker interface {
// 	Work()
// }

// type Engineer struct{}

// func (e *Engineer) Work() {
// 	fmt.Println("Working...")
// }

// func main() {
// 	var e *Engineer = nil
// 	var w Worker = e
// 	if w == nil {
// 		fmt.Println("Worker is nil")
// 	} else {
// 		fmt.Println("Worker is NOT nil")
// 	}
// }

// ----------- ex6
// type Printer interface {
// 	Print()
// }

// type Doc struct{}

// func (d *Doc) Print() {}

// func getNilPrinter() Printer {
// 	var d *Doc = nil
// 	return d
// }

// func main() {
// 	var p Printer = getNilPrinter()
// 	fmt.Println(p == nil)
// }

// ----------- ex7
// type Nothing interface{}

// func main() {
// 	var x interface{} = nil
// 	var y Nothing = nil
// 	fmt.Println(x == y) // true
// }

// ----------- ex8
// type A interface {
// 	Foo()
// }

// type B struct{}

// func main() {
// 	var a A = nil
// 	_, ok := a.(B)
// 	fmt.Println(ok)
// }

// ----------- ex9
// type A interface {
// 	Foo()
// }

// type B interface {
// 	A
// 	Bar()
// }

// type Impl struct{}

// func (Impl) Foo() {}
// func (Impl) Bar() {}

// func main() {
// 	var b B = Impl{}
// 	var a A = b

// 	a.Foo()
// 	// a.Bar() ❌ compile-time error
// }

// ----------- ex10
// type Worker interface {
// 	Work()
// }

// type Manager struct{}

// func (m *Manager) Work() {
// 	fmt.Println("Managing...")
// }

// func doWork(w Worker) {
// 	w.Work()
// }

// func main() {
// 	var m Manager
// 	doWork(m) // ❌ почему?
// }

// ----------- ex11
type Animal interface {
	Speak()
}

type Cat struct{}

func (Cat) Speak() {
	fmt.Println("meow")
}

func main() {
	cats := []Cat{{}, {}}
	animals := make([]Animal, len(cats))
	for i, c := range cats {
		animals[i] = c // vse ok
	}
}
