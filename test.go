package main
type Redis struct{

}

func (r Redis) Set(){

}

func (r Redis) Del(){

}

type Master interface{
	Set()
	Del()
}

type MasterTask struct{
	Master Master
}

func (m MasterTask) Do(ctx context.context) error{
	m.Master.Set()
}

func (m MasterTask) Undo(ctx context.context) error{
	m.Master.Del()
}


type AndProcess struct{
	Process []Proc
}

func NewAndProcess()
func (a AndProcess) Do(ctx context.context) error{
	for _, p := range a.Process{
		if err := p.Do(ctx); err != nil{
			return err
		}
	}

	return nil
}


RedisTask = AndProcess(MasterTask{Master: redis}, SlaveTask{Slave:redis}, )
