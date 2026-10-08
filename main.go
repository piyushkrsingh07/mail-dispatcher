package main

import (
	"bytes"
	"fmt"
	"html/template"
	"sync"
	//"time"
)

type Recipient struct {
	Name string
	Email string 
}

func main(){
	fmt.Println("welcome to email dispatcher")

	recipientChannel := make(chan Recipient) // unbuffered channel m size not given 
	// in case of buffered channel if consumer is not even ready toh wo queue me utna data store kr lega

//emailWorker(1,recipientChannel) // ye phle likh do phle consumer ready
  
// jb tkk srre jobs nhi ho jte tb tkk toh woh rukega
  go func(){
     loadRecipient("./users.csv",recipientChannel)
  }()

  // SYNCHRONIZATION BETWEEN THE GO ROUTINES

  // hme email workers ko sunc krne ki zaroorat h hmare main program rukna chahiye jb tkk y worker kmm nhi comlete krte


  var wg sync.WaitGroup
	
    workerCount:=5;

	for i:=1;i<= workerCount;i++ {
		wg.Add(1)
		go emailWorker(i,recipientChannel,&wg )
	}
	
	wg.Wait() // WAIT TILL wg value is 0

	//time.Sleep(3* time.Second) // 3 sec ke lye main program nhi bnd



}

func executeTemplate(r Recipient) (string,error){
	// INBUILT IN GOLABG
	t,err:=template.ParseFiles("email.tmpl") // file ko parse kr dia

	if err != nil {
		return "",err
	}
// we need to pass the buffer variable->memory space
// file ko memory space de di
var tpl bytes.Buffer
	 err=t.Execute(&tpl,r)
	 if err != nil {
		return "",err
	 }
	 ///buffer data ko string m convert

	 return tpl.String(),nil
}
