package sandbox

import (
	"os/exec"
	"strconv"

	"github.com/open-lambda/open-lambda/worker/handler/state"
)

type ProcSandbox struct {
	name string
	mgr  *ProcManager
	port int
	proc *exec.Cmd
}

func NewProcSandbox(name string, port int, mgr *ProcManager) (s *ProcSandbox, err error) {
	s = &ProcSandbox{name: name, mgr: mgr, port: port}

	s.proc = exec.Command("python", "../../lambda-generator/pyserver/server.py", strconv.Itoa(port))
	s.proc.Dir = "../../lambda-generator/pyserver/"
	s.proc.Start()

	return s, err
}

func (s *ProcSandbox) Start() error { return nil }

func (s *ProcSandbox) Stop() error { return nil }

func (s *ProcSandbox) Pause() error { return nil }

func (s *ProcSandbox) Unpause() error { return nil }

func (s *ProcSandbox) Remove() error {
	// pgid, err := syscall.Getpgid(s.proc.Process.Pid)
	// if err == nil {
	//     syscall.Kill(-pgid, 15)  // TODO(mike): no idea if this works
	// } else {
	//     return err
	// }

	// cmd.Wait()
	return nil
}

func (s *ProcSandbox) Logs() (string, error) {
	return "Logging not implemented for ProcSandbox", nil
}

func (s *ProcSandbox) State() (state.HandlerState, error) {
	return state.Running, nil
}

func (s *ProcSandbox) Port() (string, error) {
	return strconv.Itoa(s.port), nil
}
