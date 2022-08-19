package ColaPrioridad

type Tarea struct {
	prioridad int
	item      int
}

func NewTarea(q int, n int) Tarea {
	return Tarea{
		item:      q,
		prioridad: n,
	}
}

func (t Tarea) NumPrioridad() int {
	return t.prioridad
}

func (n1 Tarea) IgualQue(n2 Tarea) bool {
	return n1.prioridad == n2.prioridad
}

func (n1 Tarea) MenorQue(n2 Tarea) bool {
	// orden inverso, es decir, prioridad 0 > prioridad 1
	return n1.prioridad > n2.prioridad
}

func (n1 Tarea) MenorIgualQue(n2 Tarea) bool {
	// orden inverso, es decir, prioridad 0 > prioridad 1
	return n1.prioridad >= n2.prioridad
}

func (n1 Tarea) MayorQue(n2 Tarea) bool {

	return n1.prioridad < n2.prioridad
}

func (n1 Tarea) MayorIgualQue(n2 Tarea) bool {

	return n1.prioridad <= n2.prioridad
}
