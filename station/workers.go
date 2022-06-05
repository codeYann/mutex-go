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
