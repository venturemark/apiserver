package project

var (
	description = "Daemon for serving the venturemark grpc api."
	gitSHA      = "n/a"
	name        = "apiserver"
	source      = "https://github.com/venturemark/apiserver"
	version     = "n/a"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
