package main

import "fmt"

type Queue struct {
	Queue []interface{}
}

func (q *Queue) Enqueue(data interface{}) {
	q.Queue = append(q.Queue, data)

	return
}

func (q *Queue) Dequeue() interface{} {
	data := q.Queue[0]
	q.Queue = q.Queue[1:]

	return data
}

func main() {
	queue := new(Queue)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	fmt.Println(queue)

	fmt.Println("==> dequeuing...")
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue)
}
