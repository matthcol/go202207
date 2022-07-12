package main

import "fmt"

// goroutine producteur
// ecrit dans buffer si buffer non plein
func producteur(buffer []float64, head int, n *int) {
	for x := 1; x <= 3*len(buffer); x++ {
		data := float64(x) * 1.1
		buffer[head] = data
		head = (head + 1) % len(buffer)
		(*n)++
	}
}

// goroutine consommateur
// lit dans le buffer si buffer non vide
// remplacer la valeur lue par 0
func consommateur(buffer []float64, tail int, n *int) {
	for x := 1; x <= 3*len(buffer); x++ {
		data := buffer[tail]
		buffer[tail] = 0.0
		tail = (tail + 1) % len(buffer)
		(*n)--
		fmt.Println("Consume:", data)
	}
}

func main() {
	size := 10
	buffer := make([]float64, size)
	head := 0 // next index to write in buffer
	tail := 0 // next index to read in buffer
	n := 0    // number of floats in the buffer
	fmt.Println(buffer)
	go producteur(buffer, head, &n)
	go consommateur(buffer, tail, &n)
	fmt.Println(buffer)
}
