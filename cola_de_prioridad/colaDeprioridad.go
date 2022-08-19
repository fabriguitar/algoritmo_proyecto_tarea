package ColaPrioridad

import (
	"errors"
	"fmt"
	"strconv"
)

type ColaPrioridad struct {
	tabla        []*ColaLista
	maxPrioridad int
}

/*CONSTRUCTOR
 */

func NewColaPrioridad(n int) (colaP *ColaPrioridad, err error) {
	if n < 1 {
		err = errors.New("error en prioridad: " + strconv.Itoa(n))
	} else {
		tabla1 := make([]*ColaLista, n+1)

		for i := 0; i <= n; i++ {
			tabla1[i] = NewColaLista()
		}

		colaP = &ColaPrioridad{
			maxPrioridad: n,
			tabla:        tabla1,
		}
	}

	return colaP, err
}

/*INSERTAR
La operación añade una nueva tarea, un elemento, a la cola de prioridades. La tarea se inserta
en la cola tabla[prioridad] , siendo prioridad la asociada a la tarea.*/
func (colaP *ColaPrioridad) InsertarEnPrioridad(t Tarea) (err error) {
	p := t.NumPrioridad()

	if p >= 0 && p <= colaP.maxPrioridad {
		colaP.tabla[p].Insertar(t)
	} else {
		err = errors.New("tarea con prioridad fuera de rango")
	}

	return err
}

/*busca, en primer lugar, el elemento de máxima prioridad*/
func (cp ColaPrioridad) ElementoMin() (tarea Tarea, err error) {
	i := 0
	indiceCola := -1
	// búsqueda de la primera cola no vacía

	for {
		if !cp.tabla[i].ColaVacia() {
			indiceCola = i
			i = cp.maxPrioridad + 1 // termina el bucle

		} else {
			i++
		}

		if i > cp.maxPrioridad {
			break
		}
	}

	if indiceCola != -1 {
		tareaAux, err := cp.tabla[indiceCola].FrenteCola()
		if err == nil {
			tarea = tareaAux
		} else {
			fmt.Println(err)
		}
	} else {
		err = errors.New("cola de prioridades vacía")
	}

	return tarea, err

}

/*busca, en primer lugar, el elemento de máxima prioridad y lo elimina*/
func (cp *ColaPrioridad) QuitarMin() (tarea Tarea, err error) {
	i := 0
	indiceCola := -1
	// búsqueda de la primera cola no vacía

	for {
		if !cp.tabla[i].ColaVacia() {
			indiceCola = i
			i = cp.maxPrioridad + 1 // termina el bucle

		} else {
			i++
		}

		if i > cp.maxPrioridad {
			break
		}
	}

	if indiceCola != -1 {
		tareaAux, err := cp.tabla[indiceCola].Quitar()
		if err == nil {
			tarea = tareaAux
		} else {
			fmt.Println(err)
		}
	} else {
		err = errors.New("cola de prioridades vacía")
	}

	return tarea, err

}

//comprueba que cada una de las colas está vacía.
func (cp ColaPrioridad) ColaPrioridadVacia() bool {
	i := 0
	for cp.tabla[i].ColaVacia() && i < cp.maxPrioridad {
		i++
	}

	return cp.tabla[i].ColaVacia()
}
