package main

import (
	"fmt"
	"os"
	"sort"
)

type Persona struct {
	nombre string
}
type SortAscend []Persona

func (a SortAscend) Len() int           { return len(a) }
func (a SortAscend) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortAscend) Less(i, j int) bool { return a[i].nombre < a[j].nombre }

func main() {
	words := []Persona{}
	var n int

	fmt.Println("Ingrese un numero de palabras")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var word string
		fmt.Scanln(&word)
		words = append(words, Persona{word})
	}
	file, err := os.Create("asecendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	sort.Sort(SortAscend(words))
	for _, word := range words {
		b := []byte((word.nombre + string('\n')))
		file.Write(b)
	}
	file2, err := os.Create("descendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()
	sort.Sort(sort.Reverse(SortAscend(words)))
	for _, word := range words {
		b := []byte((word.nombre + string('\n')))
		file2.Write(b)
	}
}
