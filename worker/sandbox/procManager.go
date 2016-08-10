package sandbox

import (
    "fmt"

    "github.com/open-lambda/open-lambda/worker/config"
)

type ProcManager struct {
    registryName string
    opts         *config.Config
}

func NewProcManager(opts *config.Config) (manager *ProcManager) {
    manager = new(ProcManager)

    manager.opts = opts
    manager.registryName = fmt.Sprintf("%s:%s", opts.Registry_host, opts.Registry_port)

    return manager
}

func (pm *ProcManager) Create(name string, port int) (Sandbox, error) {
    return NewProcSandbox(name, port, pm)
}

func (pm *ProcManager) Pull(name string) error {
    // TODO

    return nil
}