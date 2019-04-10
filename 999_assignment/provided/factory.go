package provided

/* Student must implement Factory interface in seperate package */
type Factory interface {

	// Starts the workers in the factory
	// Non-Blocking
	StartUp()

	// Adds an additional Item to the processing queue
	// returns false if Shutdown has been called
	// Non-Blocking
	Add(string) bool

	// Blocks the ability to add additional items to the
	// Factory Queue. The remaining Queue must finish flushing
	// before shutdown, cannot have data loss.
	// Non-Blocking
	Shutdown()

	// this allows main to wait for execution to finish
	// Blocking
	Wait()
}
