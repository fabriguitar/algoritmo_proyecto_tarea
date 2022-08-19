package ColaPrioridad

import (
	"errors"
)

type nodo struct {
	elemento  Tarea // dato almacenado
	siguiente *nodo // enlace al siguiente nodo
}

func newNodo(x Tarea) *nodo {
	return &nodo{
		elemento:  x,
		siguiente: nil,
	}
}

//declaración de la estructura ColaLista
type ColaLista struct {
	frente *nodo
	fin    *nodo
}

// constructor: crea cola vacía
func NewColaLista() *ColaLista {
	return &ColaLista{
		frente: nil,
		fin:    nil,
	}
}

// insertar: pone elemento por el final
func (cola *ColaLista) Insertar(elemento Tarea) {

	a := newNodo(elemento)

	if cola.ColaVacia() {

		cola.frente = a
	} else {
		cola.fin.siguiente = a

	}

	cola.fin = a
}

// quitar: sale el elemento frente
func (cola *ColaLista) Quitar() (tarea Tarea, err error) {
	//var aux Tarea

	if !cola.ColaVacia() {
		tarea = cola.frente.elemento
		cola.frente = cola.frente.siguiente
	} else {
		err = errors.New("eliminar de una cola vacia")
	}
	return tarea, err
	//return aux
}

// libera todos los nodos de la cola
func (cola *ColaLista) BorrarCola() {

	for cola.frente != nil {
		cola.frente = cola.frente.siguiente
	}
}

// acceso al primero de la cola
func (cola ColaLista) FrenteCola() (tarea Tarea, err error) {
	if cola.ColaVacia() {
		err = errors.New("error cola vacia")
	} else {
		tarea = cola.frente.elemento
	}

	return tarea, err
}

// verificación del estado de la cola
func (cola ColaLista) ColaVacia() bool {
	return (cola.frente == nil)
}
