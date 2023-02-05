package main

import "fmt"

//封裝-屬性
type Triangle struct {
	B, H float32
}
type Square struct {
	len float32
}

//多形
type Shape interface {
	Area() float32
}

//封裝-方法
func (t *Triangle) Area() float32 {
	fmt.Print("Triangle Area: ")
	return t.B * t.H / 2
}

func (sq *Square) Area() float32 {
	fmt.Print("Square Area: ")
	return sq.len * sq.len
}

func main() {
	fmt.Println("======== 封裝-屬性&方法 ========")
	t := Triangle{6, 8}
	fmt.Println(t.Area())
	fmt.Println("======== 獲取和設定屬性 ========")
	s := new(Student)
	s.name = "Chi"
	fmt.Println(s.name)
	s.SetName("Eve")
	fmt.Println(s.GetName())
	fmt.Println("======== 繼承 ========")
	bus := Bus{}
	bus.Working()
	fmt.Println("======== 多形 ========")
	sq := &Square{8}
	shapes := []Shape{sq, &t}
	for i := range shapes {
		fmt.Println(shapes[i].Area())
	}

}

//獲取和設定屬性
type Student struct {
	name  string
	score float32
	Age   int
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) SetName(newName string) {
	s.name = newName
}

//繼承
func (b *Bus) Run() {
	fmt.Println("bus is runing ")
}
func (b *Bus) Stop() {
	fmt.Println("bus stop")
}

type Engine interface {
	Run()
	Stop()
}

type Bus struct {
	Engine
}

func (b *Bus) Working() {
	b.Run()
	b.Stop()
}
