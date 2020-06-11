package main

import (
	"fmt"
	localCalc "server/src/calc"
	"github.com/Julio-Assis/nummanip/calc"
	calcNew "github.com/Julio-Assis/nummanip/v2/calc"
)

func main() {
	result := calc.Add(1, 2, 3)
	fmt.Println("calc.Add(1, 2, 3) =>", result)

	newResult, err := calcNew.Add(1, 2, 3, 4)

	if err != nil {
		fmt.Println("Error: => ", err)
	} else {
		fmt.Println("calcNew.add(1, 2, 3, 4) => ", newResult)
	}

	localResult, err := localCalc.Add(12,3,4)

	fmt.Println("calcNew.add(12, 3, 4) => ", localResult)
}
