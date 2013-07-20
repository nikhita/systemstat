package main

import (
  "fmt"
  "os"
  "bitbucket.org/bertimus9/serverstat"
)

func main() {
  serverstatExample, err := serverstat.New(1, "gobi")
  if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

  serverstatExample.SetId(serverstatExample.Id() + 1)
  serverstatExample.SetName(serverstatExample.Name() + " is great")

  fmt.Println(serverstatExample.Id(), serverstatExample.Name())
  // Output: 2 gobi is great
}