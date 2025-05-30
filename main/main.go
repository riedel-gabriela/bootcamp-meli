package main

// pegar nome do modulo para exportar no go.mod
import printer "github.com/riedel-gabriela/bootcamp-meli/aula2gobases"

func main() {
	// printer.PrintNameAndAddress("Gabriela", "Rua Gentil Sandin, 380")

	// var temp int = 12
	// press := 1
	// umid := 20
	// printer.PrintPrevisaoDoTempo(temp, press, umid)

	// printer.PrintContadorLetras("paralelepipedo")
	// printer.ConcedeEmprestimo(22, true, 1.2, 50000.0)
	// printer.ConcedeEmprestimo(22, true, 0.6, 100000.0)
	// printer.ConcedeEmprestimo(22, true, 2, 100001.0)
	printer.VerificaMesFor(2)
	printer.VerificaMesSwitch(5)
	printer.VerificaMesMap(7)

}
