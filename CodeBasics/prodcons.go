package main

import (
	"fmt"
	"sync"
	"time"
)

type Buffer struct {
	Size        int
	Data        []float64
	Head        int // next index to write in buffer
	Tail        int // next index to read in buffer
	N           int // number of floats in the buffer
	mutexVide   sync.Mutex
	mutexPlein  sync.Mutex
	bufferVide  *sync.Cond
	bufferPlein *sync.Cond
}

func MakeBuffer(size int) *Buffer {
	buffer := &Buffer{
		Size: size,
		Data: make([]float64, size),
		Head: 0,
		Tail: 0,
		N:    0,
	}
	buffer.bufferVide = sync.NewCond(&buffer.mutexVide)
	buffer.bufferPlein = sync.NewCond(&buffer.mutexPlein)
	return buffer
}

func (buffer *Buffer) Depose(v float64) {
	// test si buffer plein => attendre
	if buffer.N == buffer.Size {
		buffer.bufferPlein.Wait()
	}
	buffer.Data[buffer.Head] = v
	buffer.Head = (buffer.Head + 1) % buffer.Size
	// buffer.lock.Lock()
	buffer.N++
	// buffer.lock.Unlock()
	if buffer.N == 1 {
		buffer.bufferVide.Signal()
	}
}

func (buffer *Buffer) Read() float64 {
	// test si buffer est vide => attendre
	if buffer.N == 0 {
		buffer.bufferVide.Wait()
	}
	v := buffer.Data[buffer.Tail]
	buffer.Data[buffer.Tail] = 0.0
	buffer.Tail = (buffer.Tail + 1) % buffer.Size
	// buffer.lock.Lock()
	buffer.N--
	// buffer.lock.Unlock()
	if buffer.N == buffer.Size-1 {
		buffer.bufferPlein.Signal()
	}
	return v
}

// goroutine producteur
// ecrit dans buffer si buffer non plein
func producteur(buffer *Buffer, startDelay int, deltaDelay int) {
	time.Sleep(time.Duration(startDelay) * time.Second)
	for x := 1; x <= 3*buffer.Size; x++ {
		data := float64(x) * 1.1
		fmt.Println("Produce:", data)
		buffer.Depose(data)
		time.Sleep(time.Duration(deltaDelay) * time.Second)
	}
}

// goroutine consommateur
// lit dans le buffer si buffer non vide
// remplacer la valeur lue par 0
func consommateur(buffer *Buffer, startDelay int, deltaDelay int) {
	time.Sleep(time.Duration(startDelay) * time.Second)
	for x := 1; x <= 3*buffer.Size; x++ {
		data := buffer.Read()
		fmt.Println("\tConsume:", data)
		time.Sleep(time.Duration(deltaDelay) * time.Second)
	}
}

func main() {
	size := 10
	buffer := MakeBuffer(size)
	fmt.Println(buffer)
	go producteur(buffer, 1, 1)
	go consommateur(buffer, 0, 2)
	fmt.Println(buffer)
	time.Sleep(40 * time.Second)
}
