package sandbox

type SandboxManager interface {
	Create(name string, port int) (Sandbox, error)
	Pull(name string) error
}
