package main

import (
   bookRoute "github.com/VictorEspiritu/goAPIRest/cmd/http"
)

func main() {
   route := bookRoute.SetupRouter()
   err := route.Run()
   if err != nil {
      panic(err)
   }
}

