package expense

import (
	"time"
)

type Expense struct {
	Id          int
	Date        time.Time
	Description string
	Amount      int
}

func (exp *Expense) ExpenseInit(id int, description string, amount int) {
	exp.Id = id
	exp.Description = description
	exp.Amount = amount
	exp.Date = time.Now()
}

type ExpenseList struct {
	List []Expense
}

func (e *ExpenseList) AddExpense(description string, amount int) {
	exp := Expense{}
	exp.ExpenseInit(len(e.List)+1, description, amount)

	e.List = append(e.List, exp)
}

func (e *ExpenseList) DeleteExpense(id int) {
	for index := range e.List {
		if e.List[index].Id == id {
			e.List = append(e.List[:index], e.List[index+1:]...)
		}
	}
}

func (e *ExpenseList) Summary() int {
	sum := 0
	for index := range e.List {
		sum += e.List[index].Amount
	}
	return sum
}

func (e *ExpenseList) All() []Expense {
	return e.List
}
