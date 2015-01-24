package agent

import (
	"encoding/json"
	"fmt"
	"github.com/alicebob/procspy"
	"time"
)

func listener(c <-chan []byte) {
	for {
		msg := <-c
		fmt.Println(string(msg))
		//time.Sleep(time.Second * 2)
	}
}

func startMonitor(channel chan<- []byte, scanningSeconds int64) {
	for {
		cs, err := procspy.Connections(true)
		if err != nil {
			panic(err)
		}
		js, err := json.Marshal(cs)
		if err != nil {
			fmt.Println(err)
		}
		channel <- js
		time.Sleep(time.Second * time.Duration(scanningSeconds))
	}

}

func Run() {
	fmt.Println("Start agent")

	var c chan []byte = make(chan []byte)

	go listener(c)

	startMonitor(c, 1)
}
