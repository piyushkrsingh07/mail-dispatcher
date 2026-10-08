package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient,w* sync.WaitGroup) {
     defer w.Done();


	for recipient := range ch {
		smtpHost:="localhost"
		smtpPort:="1025"

		// string byte ki slice hogi

		//formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\n%s\r\n",recipient.Email,"Just testing email campaign")


          
		// msg:=[]byte(formattedMsg)

				msg,err:=executeTemplate(recipient)

				if err != nil {
					fmt.Printf("Worker %d: Error parsing template for %s: %v\n",id,recipient.Email,err)
					// proper erorr handling
					continue; // if ek email fail to next execute
				}

		fmt.Printf("Worker %d: Sending email to %s\n", id, recipient.Email)
//= → assign a value to an existing variable
		err=smtp.SendMail(smtpHost+":"+smtpPort,nil,"piyush02040@gmail.com",[]string{recipient.Email},[]byte(msg)) // byte slice pass krna hota hai

		if err!=nil {
			log.Fatal(err)
		}

		// ek hi brr mai nhi krenge kyuki rate limit ki errror aa skti hai
		time.Sleep(50*time.Millisecond)
		fmt.Printf("Worker %d: Sent email to %s\n", id, recipient.Email)

		fmt.Printf("Worker %d processed: %v\n", id, recipient)
	}
	// consumer wait krta hai toh wait ktm kb hoga jb tkk channel close nhi hota tkk no ktm
	// but channel toh close hua hi nhi
	// hmne channel close kr dia toh y loop end ho jaega
}