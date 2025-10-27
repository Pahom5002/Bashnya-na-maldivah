package main

import "fmt"

var a int

type stack struct {
	values []int
}

// добавляет на вершину стека
func (s *stack) push(value int) {
	s.values = append(s.values, value)
}

// размер стека
func (s *stack) size() int {
	return len(s.values)
}

// пустой ли(прада/ложь)
func (s *stack) isEmpty() bool {
	return len(s.values) == 0
}

// очистка стека
func (s *stack) clear() {
	s.values = s.values[:0]
}

// вывод
func (s *stack) write() {
	a = 0
	fmt.Println("Вывод содержимого:")

	for a < len(s.values)-1 {

		a++

		fmt.Print(s.values[a-1], ", ")
	}
	fmt.Print(s.values[a], " | ")
	fmt.Println(" ")
}

// pop заменяет крайний элемент на первый
func (s *stack) pop() {

	s.values = s.values[:len(s.values)-1]
	s.values = append(s.values, s.values[0])

}
func main() {
	stack := &stack{}
	fmt.Println("стек пуст?", stack.isEmpty())

	fmt.Println("размер стека:", stack.size())

	stack.push(10)
	stack.push(34)
	stack.push(13)
	stack.push(61)
	stack.push(1)
	stack.push(90)
	fmt.Println(" ")
	fmt.Println("стек пуст?", stack.isEmpty())
	fmt.Println("размер стека:", stack.size())
	fmt.Println(" ")
	stack.write()
	stack.pop()
	fmt.Println("размер стека:", stack.size())
	fmt.Println(" ")
	stack.write()
	fmt.Println("размер стека:", stack.size())
	stack.clear()
	fmt.Println(" ")
	fmt.Println("стек пуст?", stack.isEmpty())
	fmt.Println("размер стека:", stack.size())

}
