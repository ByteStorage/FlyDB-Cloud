package image

// Image is the interface for image operations.
type Image interface {
	// Login logs in to a container registry.
	Login(opts *LoginOptions) error
	// Logout logs out of a container registry.
	Logout(opts *LogoutOptions) error
	// Push pushes an image to a container registry.
	Push(opts *PushOptions) error
	// Pull pulls an image from a container registry.
	Pull(opts *PullOptions) (string, error)
	// Images lists images.
	Images(opts *ImagesOptions) error
	// Run runs an image.
	Run(opts *RunOptions) error
	// Rm removes containers.
	Rm(opts *RmOptions) error
	// Rmi removes images.
	Rmi(opts *RmiOptions) error
	// Tag tags an image.
	Tag(opts *TagOptions) error
	// Inspect inspects an image.
	Inspect(opts *InspectOptions) (*Spec, error)
	// Build builds an image.
	Build(opts *BuildOptions) (string, error)
}
