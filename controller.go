package main

import (
  "net/http"
  "strconv"
)

type controller struct {
  hubs []*hub
}

//Global controller struct.
var contr = controller {
  hubs : make([]*hub, 0),
}

func (c *controller) addHub() {
  h := hub{
  	broadcast:   make(chan []byte),
  	register:    make(chan *connection),
  	unregister:  make(chan *connection),
  	connections: make(map[*connection]bool),
  }
  contr.hubs = append(contr.hubs, &h)
}

//Runs all hubs in goroutines
func (c *controller) run() {
    for i, h := range contr.hubs{
      go h.run()
      num := strconv.Itoa(i)
      http.HandleFunc("/"+num, h.serveHome)
      http.HandleFunc("/ws"+"/"+num, h.serveWs)
    }
}
