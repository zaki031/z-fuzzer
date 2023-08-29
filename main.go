package main

import (


  "fmt"
  "net/http"
  "bufio"
  "os"
  
  
  
)




//checking the given route 
func checkRoute(url, route string){
  newUrl := url+"/"+route
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



//looping into the list txt file and checking the StatusCode

func CheckList(url,path string){

  f, err := os.Open(path)
  if err != nil {
  fmt.Println(err)
  }
  scanner := bufio.NewScanner(f)
  for scanner.Scan(){

    checkRoute(url,scanner.Text())
  }
  if err != nil{
    fmt.Println("f")
  }


}



func userInput(){
  
  var url string

  fmt.Println("Enter the link of the website you want to scan : ")
  fmt.Scan(&url)

  CheckList(url,"./list.txt")
  

}
 
func main() {
    
  
  userInput()

}



