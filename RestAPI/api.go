package main

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	
)


type Program_response struct  {
	Sequencial string 		 `json:"Sequencial"`
	TimeToExecute float64	 `json:"TimeToExecute"`

}

func Seq (counter int) float64{

	start := time.Now()

	for i := 0; i < counter ; i++ {

		time.Sleep(time.Millisecond)
	}
	elapsed :=  time.Since(start)

	return elapsed.Seconds()

}
func Wait(Counter int ,c chan int){

	for i := 0; i < Counter ; i++ {

		time.Sleep(time.Millisecond)


	}

	c <- 1



}

func Parallel() float64{

	channel := make(chan int) 

	start := time.Now()

	go Wait(1,channel)
	go Wait(1,channel)
	go Wait(1,channel)
	go Wait(1,channel)

	<-channel;<-channel;<-channel;<-channel

	elapsed :=  time.Since(start)

	return elapsed.Seconds()

}
func main() {

	

	fmt.Printf("Starting the Program Timer ... \n")

	http.HandleFunc("/sequencial", func(rw http.ResponseWriter, r *http.Request) {

		time := Seq(4)

		rw.Header().Set("Content-Type", "application/json")

		response := Program_response{
			Sequencial: "True", 
			TimeToExecute: time}
         json.NewEncoder(rw).Encode(response)	
		
	})
	http.HandleFunc("/parallel", func(rw http.ResponseWriter, r *http.Request) {

		time := Parallel()

		rw.Header().Set("Content-Type", "application/json")

		response := Program_response{
			Sequencial: "False", 
			TimeToExecute: time}
         json.NewEncoder(rw).Encode(response)	
		
	})

	// Listen for connections on all ip addresses (0.0.0.0)
	// port 9090
	fmt.Printf("Starting Server...")
	err := http.ListenAndServe(":9090", nil)
	fmt.Println(err)


	


	
   
	

	
}