package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// w - responceWritter(Куда писать ответ )
// r - request (Откуда брать запрос)

// Функция-обработчик
func GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hi! I`m new web-server</h1>")
}

func main() {
	http.HandleFunc("/", GetGreet)                                          // Если придет запрос на адрес, то вызывай GetGreet
	log.Fatal(http.ListenAndServe("0.0.0.0:"+os.Getenv("SERVERPORT"), nil)) // Запускаем web-server в режиме "слушания"
}
