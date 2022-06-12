package main

import (
	"fmt"
)

// STACKS
// Struktur data seperti tumpukan buku
// Method push : menambahkan item baru di akhir
// Method pop : menghilangkan item paling akhir

type Stacks struct {
	items []int
}

func (s *Stacks) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stacks) Pop() int {
	l := len(s.items)-1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

// QUEUE
// Struktur data seperti antrian
// Punya method enqueue : menambahkan item paling belakang
// Punya method dequeue : mengeluarkan item paling depan

type Queue struct {
	items []int
}

// enqueue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// dequeue
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {

	// stacks
	fmt.Println("STACKS")
	stacks := Stacks{}
	stacks.Push(1)
	stacks.Push(2)
	stacks.Push(3)
	fmt.Println("Sesudah di-Push")
	fmt.Println(stacks)
	stacks.Pop()
	fmt.Println("Sesudah di-Pop")
	fmt.Println(stacks)

	// queues
	fmt.Println("QUEUE")
	queue := Queue{}
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(6)
	fmt.Println("Sesudah di-Enqueue")
	fmt.Println(queue)
	queue.Dequeue()
	fmt.Println("Sesudah di-Dequeue")
	fmt.Println(queue)
}