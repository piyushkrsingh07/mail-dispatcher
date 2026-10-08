package main

import (
	"fmt"
	"sync"
)

func emailWorker(id int, ch chan Recipient,w* sync.WaitGroup) {
     defer w.Done();
	for recipient := range ch {
		fmt.Println(id,recipient)
	}
}