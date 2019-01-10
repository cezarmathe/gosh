package shell

// Run contains the entire shell lifecycle
func Run() {

	initialize()

	defer terminate()

	for {
		interpret()
	}
}
