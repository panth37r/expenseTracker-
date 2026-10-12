package main

import (
	"expenseTracker/expense"
	"expenseTracker/storage"
)

type App struct {
	List *expense.ExpenseList
}

func NewApp() (*App, error) {
	a := &App{List: &expense.ExpenseList{}}
	err := storage.ReadFromFile(a.List)
	if err != nil {
		return a, err
	}
	return a, nil
}

func (a *App) AddExpense(description string, amount int) error {
	a.List.AddExpense(description, amount)
	err := storage.SaveToFile(a.List)
	return err
}

func (a *App) DeleteExpense(id int) error {
	a.List.DeleteExpense(id)
	err := storage.SaveToFile(a.List)
	return err
}

func (a *App) GetAll() []expense.Expense {
	return a.List.All()
}

func (a *App) GetSummary() int {
	return a.List.Summary()
}
