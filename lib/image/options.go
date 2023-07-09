package image

type LoginOptions struct {
	Domain        string
	Username      string
	Password      string
	SkipTLSVerify bool
}

type LogoutOptions struct {
	All    bool
	Domain string
}

type PushOptions struct {
	Authfile      string
	CertDir       string
	Format        string
	Image         string
	Destination   string
	Rm            bool
	Quiet         bool
	SkipTLSVerify bool
	All           bool
}

type PullOptions struct {
	CertDir       string
	Quiet         bool
	SkipTLSVerify bool
	PullPolicy    string
	Image         string
	Platform      string
}

type ImagesOptions struct {
	All       bool
	Digests   bool
	NoHeading bool
	NoTrunc   bool
	Quiet     bool
	History   bool
	JSON      bool
}

type RunOptions struct {
	Image string
	Quiet bool
}

type RmOptions struct {
	ContainerNamesOrIDs []string
	All                 bool
}

type RmiOptions struct {
	ImageNamesOrIDs []string
	Force           bool
	Prune           bool
}

type TagOptions struct {
	ImageNameOrID string
	Tags          []string
}

type InspectOptions struct {
	Format        string
	InspectType   string
	ImageNameOrID string
}

type BuildOptions struct {
	Kubefile       string
	DockerFilePath string
	ContextDir     string
	PullPolicy     string
	ImageType      string
	//Manifest          string
	Tag               string
	BuildArgs         []string
	Platforms         []string
	Labels            []string
	Annotations       []string
	NoCache           bool
	Base              bool
	ImageList         string
	ImageListWithAuth string
	IgnoredImageList  string

	//BuildMode means whether to download container image during the build process
	// default value is download all container images.
	BuildMode string
}

type Spec struct {
	// implements me
}
