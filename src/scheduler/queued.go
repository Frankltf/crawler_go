package scheduler

import "engine"

type QueuedScheduler struct {
	requestchan chan engine.Request
	workchan chan chan engine.Request 
}

func (s *QueuedScheduler)WorkReady(w chan engine.Request)  {
	s.workchan<-w
}
func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestchan<-r
}

func (s *QueuedScheduler) ConfigureWorkchan(c chan engine.Request) {
	panic("implement me")
}
func (s *QueuedScheduler) Run()  {
	s.requestchan=make(chan engine.Request)
	s.workchan=make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workQ []chan engine.Request
		for  {
			var activeRequest  engine.Request
			var activeWorker chan engine.Request
			if len(requestQ)>0 && len(workQ)>0 {
				activeRequest=requestQ[0]
				activeWorker=workQ[0]
			}
			select {
				case r:=<-s.requestchan:
					requestQ=append(requestQ, r)
				case w:=<-s.workchan:
					workQ=append(workQ,w)
				case activeWorker<-activeRequest:
					requestQ=requestQ[1:]
					workQ=workQ[1:]
			}


		}
	}()
}
