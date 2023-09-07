package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
)


type fuzzer struct {
  Routes []string
  URL string
}

// creating a new fuzzer with the specified routes
func NewFuzzer(url string,path string)*fuzzer{
  var routes []string
  f, err := os.Open(path)
  if err != nil {
  panic(err)
  }
  scanner := bufio.NewScanner(f)
  for scanner.Scan(){

    routes = append(routes,scanner.Text())
  }
  if err != nil{
    panic("Could Not Read file")
  }


  return &fuzzer{Routes:routes,URL:url} 
}

//checking the given route 
func (f *fuzzer ) CheckRoutes(){
  for _,route := range f.Routes {
      newUrl := f.URL+"/"+route
      res, err := http.Get(newUrl)
      if err!=nil{
        fmt.Println(err)
      }

      if res.StatusCode == 200{
      fmt.Println(route,"--- successful connection ---", res.StatusCode)
      }
      if res.StatusCode == 404{
        fmt.Println(route,"--- route doesnt exist ---", res.StatusCode)
      }

  }
 }

 
func main() {
  //getting vlag values
  url := flag.String("url", "", "website url")
  flag.Parse()

  fuzzer := NewFuzzer(*url, "./list.txt")
  fuzzer.CheckRoutes()

}

