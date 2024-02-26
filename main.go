package main

import (
	"log"
	"net/http"
)

func main() { //Função Principal

	fs := http.FileServer(http.Dir("./estatico")) //Inserir localização do diretório do servidor estático
	http.Handle("/", fs)
	log.Print("Listening on: 3000...") //Pacote log irá chamar a porta definida
	err := http.ListenAndServe(":3000", nil)
	if err != nil { //Se houver o erro, chama a variável erro
		log.Fatal(err)

	}

}
