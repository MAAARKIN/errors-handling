package main

import (
	"fmt"
	"log"

	"github.com/maaarkin/errors-handling/internal/service"
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {

	todoService := service.NewTodoService()

	todo, err := todoService.Save("teste")
	if err != nil { //existe error a ser tratado

		if st, ok := err.(stackTracer); ok {
			log.Fatalf("%+v", st.StackTrace()[0:2])
		} else {
			//condicional para garantir que caso o error identificado
			//não tenha sido corretamente tratado com o pkg/errors
			log.Fatal(err) //utilizando o fatal para interromper a execução do main.
		}
	}

	fmt.Println(todo.Title)
}
