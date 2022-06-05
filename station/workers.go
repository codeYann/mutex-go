package station

// Estrutura para os funcionários
type Worker struct {
	name string
	id   int
	role string
}

// Criando um funcionário
func CreateWorker(nameWorker, roleWorker string, idWorker int) *Worker {
	return &Worker{
		name: nameWorker,
		id:   idWorker,
		role: roleWorker,
	}
}
