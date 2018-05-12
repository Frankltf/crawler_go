package scheduler

import "engine"

type SimpleScheduler struct {
	Workchan chan engine.Request
}

func (s *SimpleScheduler)Submit(r engine.Request)  {
	go func() {
		s.Workchan<-r
	}()

}
func (s *SimpleScheduler) ConfigureWorkchan(c chan engine.Request)  {
	 s.Workchan=c
}