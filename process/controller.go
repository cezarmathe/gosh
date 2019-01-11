package process

import (
	"fmt"
	"os"
)

// ProcessController is a controler for shell processes
type ProcessController struct {
	processList      map[int]*os.Process
	ActiveProcessPid int
}

// NewProcessController creates a new process controller
func NewProcessController() ProcessController {
	processController := ProcessController{
		processList:      make(map[int]*os.Process),
		ActiveProcessPid: os.Getpid(),
	}
	return processController
}

// StartProcess starts a process and waits for it to finish
func (processController *ProcessController) StartProcess(name string, argv []string, attr *os.ProcAttr) {
	// add the name of the executable at the beggining of argv
	argv = append([]string{name}, argv...)

	// create & start the process
	process, err := os.StartProcess(name, argv, attr)

	// report any errors
	if err != nil {
		fmt.Println(err)
		return
	}

	// set a new active process pid and add the process to the map
	processController.ActiveProcessPid = process.Pid
	processController.processList[process.Pid] = process

	// wait for the process to finish and clean up after
	processPid := process.Pid
	defer delete(processController.processList, processPid)
	process.Wait()

	// return to the shell
	processController.ActiveProcessPid = os.Getpid()
}

// StartBgProcess starts a background process and returns its pid
func (processController *ProcessController) StartBgProcess(name string, argv []string, attr *os.ProcAttr) (int, error) {
	// add the name of the executable at the beggining of argv
	argv = append([]string{name}, argv...)

	// create & start the process
	process, err := os.StartProcess(name, argv, attr)

	// report any errors
	if err != nil {
		return 0, err
	}

	// add the process to the map
	processController.processList[process.Pid] = process

	// wait for the process to finish in the background and clean up after
	go func() {
		processPid := process.Pid
		defer delete(processController.processList, processPid)
		process.Wait()
	}()

	return process.Pid, nil
}

// SignalActiveProcess signals the current active process
func (processController *ProcessController) SignalActiveProcess(signal os.Signal) error {
	return processController.processList[processController.ActiveProcessPid].Signal(signal)
}

// SignalProcessPid signals a process by it's pid
func (processController *ProcessController) SignalProcessPid(pid int, signal os.Signal) error {
	return processController.processList[pid].Signal(signal)
}
