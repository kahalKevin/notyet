package service_handler

import "fmt"

var WorkerQueue chan chan Job

func StartDispatcher(nworkers int) {
  // First, initialize the channel we are going to but the workers' work channels into.
  WorkerQueue = make(chan chan Job, nworkers)
  
  // Now, create all of our workers.
  for i := 0; i<nworkers; i++ {
    fmt.Println("Starting worker", i+1)
    worker := NewWorker(i+1, WorkerQueue)
    worker.Start()
  }
  
  go func() {
    for {
      select {
      case work := <-WorkQueue:
        fmt.Println("Received work requeust")
        go func() {
          worker := <-WorkerQueue
          
          fmt.Println("Dispatching work request")
          worker <- work
        }()
      }
    }
  }()
}