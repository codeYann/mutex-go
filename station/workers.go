// Nome: Yan Rodrigues da Silva
// Matrícula: 495532

package station

type Worker struct {
	name string
	id   int
	role string
}

func CreateWorker(nameWorker, roleWorker string, idWorker int) *Worker {
	return &Worker{
		name: nameWorker,
		id:   idWorker,
		role: roleWorker,
	}
}
