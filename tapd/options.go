package tapd

type Options struct {
	username string
	password string

	endpoint string

	workspaceID int
}

type Option func(*Options)

func WithBasicAuth(username, password string) Option {
	return func(o *Options) {
		o.username = username
		o.password = password
	}
}

func WithEndpoint(endpoint string) Option {
	return func(o *Options) {
		o.endpoint = endpoint
	}
}

func WithWorkspaceID(workspaceID int) Option {
	return func(o *Options) {
		o.workspaceID = workspaceID
	}
}
