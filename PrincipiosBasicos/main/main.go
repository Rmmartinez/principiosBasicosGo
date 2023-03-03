package main

import "fmt"

type registro struct {
	nombre      string
	apellid     string
	universidad string
}

// esto sería un método COPY para REGISTRO
func (p registro) copy(copy string) {
	fmt.Println(copy + ", " + p.getUniversidad())
}

// get y set para Registro
func (p registro) getNombre() string {
	return p.nombre
}

func (p registro) getApellido() string {
	return p.apellid
}

func (p registro) getUniversidad() string {
	return p.universidad
}

func (p registro) setNombre(nom string) {
	p.nombre = nom
}

func (p registro) setApellido(ape string) {
	p.apellid = ape
}

func (p registro) setUniversidad(uni string) {
	p.universidad = uni
}
func main() {

	// manejo de listas
	paises := []string{"Argentina", "Uruguay", "Brasil"}
	// agrego elementos con append
	paises = append(paises, "Chile", "Perú")
	fmt.Println("\nPaises donde está disponible el sistema: ")
	// recorro e imprimo
	for i := 0; i < len(paises); i++ {
		fmt.Println(" -", paises[i])
	}
	fmt.Println("")

	/*
		var entero int = 2
		var decimal float64 = 2.5
		var texto string = "Hola Mundo"
		var booleana bool = true

		//Para asignar valor a variable, pero sin declararla con var
		//Durante el resto del programa, no puedo cambiarle el tipo. Ahora es int
		numero := 3

		fmt.Println("Entero: ", entero, "\nDecimal: ", decimal, "\nTexto: ", texto, "\nBooleanoo: ", booleana,
		"\n\nNumero sin declarar var: ", numero)
	*/

	//repaso operadores lógicos
	bool1 := false
	bool2 := true
	fmt.Println("¿Son iguales los valores booleanos?: ")
	if bool1 == bool2 {
		fmt.Println("Si, son iguales.")
	} else {
		fmt.Println("No, son distintos.")
	}

	fmt.Println("\n¿Ambos valores son true?: ")
	if bool1 == true && bool2 == true {
		fmt.Println("Si, ambos valores son true.")
	} else {
		if bool1 == false && bool2 == false {
			fmt.Println("No, ambos son false.")
		} else {
			fmt.Println("No, un valor es true y el otro false.")
		}
	}
	fmt.Println("\nFinalizando...")
	cuentaRegresiva(5)
	//fmt.Println("\nResultado:", sumoValores(4, 3))

	//trabajo con Struct
	persona := registro{"Magali", "Martinez", "UNSL"}
	fmt.Println("\n\nAutor:", persona.getNombre(), persona.getApellido())
	persona.copy("Todos los derechos reservados")
}

// recibe el numero desde el que quiero empezar la cuenta regresiva
func cuentaRegresiva(num int) {
	//En Golang NO hay while
	for i := num; i >= 0; i-- {
		fmt.Println(i)
	}
}

// recibe dos enteros, y retorna un entero que es el resultado de la suma
func sumoValores(val1 int, val2 int) int {
	return val1 + val2
}
