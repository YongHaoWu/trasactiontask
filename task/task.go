package task

// Process define base task include process
type Process struct {
	Rollback string
	Goon     string
	Action   string
}

// Task define base task include process
type Task struct {
	Processes      []Process
	ActionCallback func(string) bool
}

// Exec .....
func (t *Task) Exec() {
	var idx = 0
	needRollback := false
	for i, p := range t.Processes {
		idx = i
		if !t.ActionCallback(p.Goon) {
			needRollback = true
			break
		}
	}
	if needRollback {
		for i := idx; i >= 0; i-- {
			p := t.Processes[i]
			if p.Action == "continue" {
				t.ActionCallback(p.Rollback)
			}
		}
	}
}
