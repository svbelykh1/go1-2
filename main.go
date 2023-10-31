package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	fmt.Println("Test")

	//Запрос

	client := &http.Client{Timeout: time.Second}

	resOne, err := client.Get("http://golang.org/doc")
	if err != nil {
		log.Fatal(err)
	}
	// Свойства
	// POST =
	fmt.Println(resOne.Status, resOne.StatusCode, resOne.Request.URL)
	fmt.Println(resOne.Request.Header, resOne.Request.Response.Header)

	body, _ := io.ReadAll(resOne.Body) //Получаем тело запроса
	resOne.Body.Close()

	file, err := os.Create("out_site.txt") //Открываем файл

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()        // Закрываем файл
	_, err = file.Write(body) //Сохраняем в файл

	if err != nil {
		log.Fatal(err)
	}

}
