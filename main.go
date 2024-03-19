package main

import "fmt"

func main() {
	marco("Tipo de Vehículo ║ Tarifa diurna(6 a 17) ║ Tarifa noturna(18 a 5)")

}

func marco(msj string) {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Print("╔")
	for x := 0; x < len(msj); x++ {
		fmt.Print("═")
	}
	fmt.Print("╗")
	fmt.Println("")
	fmt.Print("║", msj, "     ║")
	fmt.Println("")
	fmt.Print("╚")
	for x := 0; x < len(msj); x++ {
		fmt.Print("═")
	}
	fmt.Print("╝")
	fmt.Println("")
	fmt.Print("║", msj, "     ║")
	fmt.Println("")
	fmt.Print("╚")
	for x := 0; x < len(msj); x++ {
		fmt.Print("═")
	}
	fmt.Print("╝")

	fmt.Println("")
	fmt.Print("║", msj, "     ║")
	fmt.Println("")
	fmt.Print("╚")
	for x := 0; x < len(msj); x++ {
		fmt.Print("═")
	}
	fmt.Print("╝")
	fmt.Println("")
	fmt.Print("║", msj, "     ║")
	fmt.Println("")
	fmt.Print("╚")
	for x := 0; x < len(msj); x++ {
		fmt.Print("═")
	}
	fmt.Print("╝")

	fmt.Println("")
	fmt.Print("║", msj, "     ║")
	fmt.Println("")
	fmt.Print("╚")
	for x := 0; x < len(msj); x++ {
		fmt.Print("═")
	}
	fmt.Print("╝")

	fmt.Println("")
	fmt.Print("║", msj, "     ║")
	fmt.Println("")
	fmt.Print("╚")
	for x := 0; x < len(msj); x++ {
		fmt.Print("═")
	}

	fmt.Print("╝")
	fmt.Println("")
	fmt.Print("║", msj, "     ║")
	fmt.Println("")
	fmt.Print("╚")
	for x := 0; x < len(msj); x++ {
		fmt.Print("═")
	}
	fmt.Print("╝")
}
}
