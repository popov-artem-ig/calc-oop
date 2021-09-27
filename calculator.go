package main

import (
	"calc-oop/structures"
	"calc-oop/utils"
	"fmt"
	"strconv"
)

func Calculate(expressionRpn []string) (float64, error) {
	stackOp := structures.Stack{} // Стек операндов

	for i := range expressionRpn {
		s := expressionRpn[i]
		//utils.DebugLog(s, stackOp, data_structures.Queue{})
		var v float64
		var err error

		if utils.IsFloat(s) {
			stackOp.Push(s)
			continue
		}
		//Если это оператор, соответствующее количество операндов перекладывается из стека во временные переменные
		v, stackOp, err = evaluateOperator(s, stackOp)
		if err != nil {
			return 0, err
		}
		stackOp.Push(fmt.Sprintf("%f", v))
	}

	if(stackOp.Size() > 1) {
		panic("Input did not fully simplify")
	}

	stRes := stackOp.Pop()
	result, _ := strconv.ParseFloat(*stRes, 64)
	return result, nil
}

func evaluateOperator(operator string, stackOp structures.Stack) (float64, structures.Stack, error) {
	af, bf := 0.0, 0.0
	var err error
	b := stackOp.Pop()
	a := stackOp.Pop()
	if a != nil {
		af, err = strconv.ParseFloat(*a, 64)
	}
	bf, err = strconv.ParseFloat(*b, 64)

	if err != nil {
		return 0, stackOp, err
	}

	switch operator {
	case "+":
		return af + bf, stackOp, nil
	case "-":
		return af - bf, stackOp, nil
	case "*":
		return af * bf, stackOp, nil
	case "/":
		if bf != 0.0 {
			return af / bf, stackOp, nil
		} else {
			return 0, stackOp, fmt.Errorf("division by 0 is prohibited")
		}
	}

	return 0, stackOp, fmt.Errorf("unexpected operator")
}

func isContainsFunc(x string) bool {
	arr := []string{"*", "/", "+", "-"}
	for _, n := range arr {
		if x == n {
			return true
		}
	}
	return false
}