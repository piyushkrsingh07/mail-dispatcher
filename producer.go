package main

import (
	"encoding/csv"
	//"fmt"
	"os"
)

func loadRecipient(filePath string, ch chan Recipient) error{
   // ------------------ AFTER ALL THE MAILS HAVE BEEN READ ----------------------
   defer close(ch)
	// to read the file
	f,err:=os.Open(filePath)

	if err != nil {
		return err
	}

	defer f.Close() // jaise hi load recipent function over due to anything then use defer keyword


	r:=csv.NewReader(f) // return object

	records,err:=r.ReadAll()
	if err != nil {
		return err
	}

	

	for _,record := range records[1:]{
		// 1st index k aage wle return 

		//fmt.Println(record)
		// send to the consumer -> channel
		// record look like this 

		// [User 1 user1@example.com]

		ch <- Recipient{
			Name: record[0],
			Email: record[1],
		}

		// jb tkk y koi read krne wla nhi rtha tb tkk ye blocking hoga

	//	fatal error: all goroutines are asleep - deadlock!->THIS WILL BE THE ERROR

	}

	return nil
	

}