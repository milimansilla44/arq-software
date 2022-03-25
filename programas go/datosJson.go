package main
import (
	"fmt"
	"errors"
	"net/http"
	"ioutil" 
	"encoding/json"
)
type Categories struct[]Category

type Category struct{
	ID string'json:"id"'
	name string 'json: "name"'

}
func main(){
	cats err := GetCategories("MLA")
	if err!= nil{
		fmt.Println("Error: ", err.Error())
		exit(1)
	}
	fmt.Printl("Las categorias de MLa son: ")
}
func GetCategories(siteID string)(Categories, error){
	response := http.Get("https://api.mercadolibre.com/sites/MLA/categories")

	bytes:= ioutil.ReadAll(response.bytes)//completar

	var cats Categories
	err := json.Unmarshal(bytes, &cats)

	return cats, nil
}

