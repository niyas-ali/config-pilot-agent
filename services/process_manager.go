package services

type ProcessManager struct {
	jobs []GitProcess
}

func (p *ProcessManager) AddProcess(process GitProcess) {
	p.jobs = append(p.jobs, process)
}
func (p *ProcessManager) Run() {
	for _, job := range p.jobs {
		job.Run()
	}
}
