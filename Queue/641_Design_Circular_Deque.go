package main

type MyCircularDeque struct {
	queue       []int
	front, rear int
	size        int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		queue: make([]int, k+1, k+1),
		front: 0,
		rear:  0,
		size:  k + 1,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	// because both front and rear starts from 0, so update fron first then assign value
	this.front = (this.front - 1 + this.size) % this.size
	this.queue[this.front] = value

	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	this.queue[this.rear] = value
	this.rear = (this.rear + 1 + this.size) % this.size

	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.front = (this.front + 1 + this.size) % this.size

	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.rear = (this.rear - 1 + this.size) % this.size

	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[this.front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[(this.rear-1+this.size)%this.size]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.front == this.rear
}

func (this *MyCircularDeque) IsFull() bool {
	return (this.rear+1+this.size)%this.size == this.front
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
