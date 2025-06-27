package aula2gobases

import (
	"fmt"
)

func PrintPrevisaoDoTempo(temp int, press int, umid int) {
	fmt.Printf("temp %d type %T\npress %d type %T\numidade %d type %T\n", temp, temp, press, press, umid, umid)
}

// func main() {
// 	// temp := 12
// 	// press := 1
// 	// umid := 20
// 	printPrevisaoDoTempo(temp, press, umid)
// }
