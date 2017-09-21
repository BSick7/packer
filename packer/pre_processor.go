package packer

// A PreProcessor is responsible for performing operations before a build.
// An example of a pre-processor would be pulling a source artifact.
type PreProcessor interface {
	// Configure is responsible for setting up configuration,
	// storing the state for later, and returning any errors,
	// such as validation errors
	Configure(...interface{}) error

	// PreProcess performs an action before a builder executes.
	PreProcess(Ui) (err error)
}
