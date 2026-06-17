package main

import "log"

func main() {
	log.Println("App started")

	log.Fatal("DB connection failed")//make it exit

	log.Println("This will NEVER run")
}
//difference between log and fmt print is it adds time stamp and date