// Copyright (c) avijit bhattacharjee 2024

package JSON_Handling


import (
	"fmt"
	"encoding/json"
)

type Salary struct {
	Fixed int `json: "fixed"`
	Variable int `json: "variable"`
}

type OtherIncome struct {
	Interest int `json: "interest"`
	CapitalGain int `json: "capitalGain"`
}

type Income struct {
	IncomeSalary Salary
	OtherSOurce OtherIncome
}

func Marshall() {

	var salary = Salary{
		Fixed : 10,
		Variable: 2,

	}

	var other = OtherIncome{
		Interest : 2,
		CapitalGain: 1,

	}


	var income = Income{
		IncomeSalary: salary,
		OtherSOurce: other,
	}

	byteArr, _ := json.MarshalIndent(income, "", " ")
	fmt.Println(byteArr)

	
}