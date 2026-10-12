package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main() {
	app, err := NewApp()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("zero arguments")
		fmt.Println("try ./TRACKER help")
		return
	}
	switch os.Args[1] {
	case "add":
		desc := os.Args[2]
		amount, err := strconv.Atoi(os.Args[3])
		err = app.AddExpense(desc, amount)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "list":
		json, err := json.MarshalIndent(app.GetAll(), "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(json))
	case "summary":
		amount := app.GetSummary()
		fmt.Printf("Summary is $%d: \n", amount)
	case "delete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		err = app.DeleteExpense(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Deleted ID[%d]\n", id)
	case "help":
		fmt.Println("./app delete [id]")
		fmt.Println("./app summary")
		fmt.Println("./app list")
		fmt.Println("./app add DESCRIPTION AMOUNT")

	}
}
