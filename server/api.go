package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "fmt"
  "io/ioutil"
  "encoding/json"
)

func initApi () {

  ginInstance := gin.Default ()
  ginInstance.GET ("/catalogue/:category/:alpha/:page", func (context *gin.Context) {
    context.JSON (
      200, 
      getCatalogue (
        context.Param ("category"),
        context.Param ("alpha"),
        context.Param ("page"),
      ),
    )
  })
  ginInstance.Run ()

}

func getCatalogue (category string, alpha string, page string) CataloguePage {

  url := "https://secure.runescape.com/" +
    "m=itemdb_rs/api/catalogue/items.json?category=" +
    category +
    "&alpha=" +
    alpha + 
    "&page=" +
    page

  response, responseError := http.Get (url)

  if responseError != nil {

    fmt.Println (responseError)

  }

  responseBody, bodyError := ioutil.ReadAll (response.Body)

  if bodyError != nil {

    fmt.Println (bodyError)

  }
 
  var catalogue CataloguePage
  bindingError := json.Unmarshal ([]byte(string (responseBody)), &catalogue)

  if bindingError != nil {

    fmt.Println (bindingError)

  }

  return catalogue

}
