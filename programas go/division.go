package main
import (
	"errors"
	"fmt"
)
func main(){

	div, err := division(7,0)
	if err!=nil(

		fmt.Println("Error: ", err.Error())
		exit(1)

	)
}
func division(a int , b int)(int, error ){
	
	if b==0 (
		return -1, errors.New("no se puede dividir por cero")
	)
	return a / b, nil  
}