package main

import (
	"fmt"

	"github.com/osamikoyo/ascii-gallery/internal/data"
	"github.com/osamikoyo/ascii-gallery/internal/data/models"
)

func main(){
	storage, err := data.Connect()
	if err != nil{
		fmt.Print(err)
	}

	err = storage.AddArt(models.Art{
		Author: "hello",
		Title: "test",
		Content: "test",
		Desc: "test",
	})
	if err != nil{
		fmt.Println(err)
	}
}