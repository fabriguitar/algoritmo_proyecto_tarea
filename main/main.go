package main

import (
	cp "cp/cola_de_prioridad"
	"fmt"
)

func main() {

	prioridad := 0
	colaP, err := cp.NewColaPrioridad(4)

	if err != nil {
		fmt.Println(err)
		return
	}

	//insertando elementos en la cola de prioridad
	for i := 0; i < 55; i++ {

		if i < 10 {
			prioridad = 4
		} else if i < 20 {
			prioridad = 3
		} else if i < 30 {
			prioridad = 2
		} else if i < 40 {
			prioridad = 1
		} else {
			prioridad = 0
		}

		t1 := cp.NewTarea(i, prioridad)
		err := colaP.InsertarEnPrioridad(t1)

		if err != nil {
			fmt.Println(err)
			return
		}

	}

	//recorriendo y vaciando la cola de prioridad

	for !colaP.ColaPrioridadVacia() {
		tarea, err := colaP.QuitarMin()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Numero de prioridad: 	", tarea.NumPrioridad(), ", Elemento tarea: ", tarea)
	}

}
