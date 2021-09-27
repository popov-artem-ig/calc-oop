package structures

import (
	"calc-oop/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type ReversePolishNotation struct {
	items []string
	rwLock sync.RWMutex
}

// New -- Конструктор создания экземляра типа
func (rpn *ReversePolishNotation) New(value string) *ReversePolishNotation {

	//if isNeedConvert(value) {
		if validate(value) {
			rpnVal, err := convert(value)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			rpn.items = rpnVal
		}
	//}
	//rpn.item = value
	return rpn
}

// Calc -- Метод вычисления обратной польской записи
func (rpn *ReversePolishNotation) Calc() (float64, error) {
	stackOp := Stack{} // Стек операндов

	for i := range rpn.items {
		s := rpn.items[i]
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

// ToString -- Метод получения строки обратной польской записи
func (rpn *ReversePolishNotation) ToString() (string, error) {
	if len(rpn.items) > 0 {
		return strings.Join(rpn.items, " "), nil
	}
	return "", fmt.Errorf("error. current value empty")
}

func (rpn *ReversePolishNotation) IsEmpty() bool {
	// Acquire read lock
	rpn.rwLock.RLock()
	// defer operation of unlock.
	defer rpn.rwLock.RUnlock()
	return len(rpn.items) == 0
}

/**
-------- Скрытая реализация --------
 */
func convert(str string) ([]string, error) {
	np := 0                        // Инкремент открытых скобок
	stackOp := Stack{}  		   // Стек операторов
	resQueue := Queue{} 		   // Результирующая очередь
	strIn := str                   // Выражение для разбора
	hoarder := ""                  // Накопитель для сборки/обработки дробных и больших чисел

	for i := range strIn {
		c := strIn[i]
		//utils.DebugLog(&stackOp)
		//utils.DebugLog(&resQueue)
		//fmt.Println(string(c))
		//fmt.Println(stackOp)
		//fmt.Println(resQueue)

		if utils.IsDigit(c) || c == '.' {
			hoarder += string(c)
			if i != len(strIn)-1 {
				continue
			}
		}

		if len(hoarder) > 0 {
			resQueue.Enqueue(hoarder)
			hoarder = ""
		}

		switch c {
		case '(':
			{
				if i == len(strIn)-1 {
					return []string{}, fmt.Errorf("syntax error. opening parenthesis at the end of an expression")
				}
				stackOp.Push(string(c))
				np++
			}
		case '*', '/', '+', '-':
			{
				if i == len(strIn)-1 {
					return []string{}, fmt.Errorf("syntax error. operator at the end of an expression")
				}

				if stackOp.Peek() == nil { //Если стек операторов пуст, алгоритм помещает входящий оператор в стек.
					stackOp.Push(string(c))
				} else {
					//Если приоритет входящего оператора ниже,
					//верхний оператор извлекается из стека и выводится в очередь,
					//после чего входящий оператор сравнивается с новой вершиной стека.
					for stackOp.Peek() != nil && prior(string(c)) < prior(*stackOp.Peek()) {
						v := stackOp.Pop()
						resQueue.Enqueue(*v)
					}
					//Если входящий оператор имеет более высокий приоритет,
					//чем тот оператор, что в настоящее время находится на вершине стека,
					//входящий оператор помещается на вершину стека.
					if stackOp.Peek() == nil || prior(string(c)) > prior(*stackOp.Peek()) {
						stackOp.Push(string(c))
					} else if stackOp.Peek() != nil && prior(string(c)) == prior(*stackOp.Peek()) {
						//Если входящий оператор имеет такой же приоритет,
						//верхний оператор извлекается из стека и выводится в очередь,
						//а входящий оператор помещается в стек.
						v := stackOp.Pop()
						if v != nil {
							resQueue.Enqueue(*v)
							stackOp.Push(string(c))
						}
					}
				}
			}
		case ')':
			{
				// До тех пор, пока верхним элементом стека не станет открывающая скобка, выталкиваем элементы из стека в выходную строку.
				for *stackOp.Peek() != "(" && *stackOp.Peek() != "" {
					v := stackOp.Pop()
					resQueue.Enqueue(*v)
				}
				// Если стек закончился раньше, чем мы встретили открывающую скобку, это означает, что в выражении либо неверно поставлен разделитель,
				// либо не согласованы скобки.
				if *stackOp.Peek() == "" {
					return []string{}, fmt.Errorf("syntax error. inconsistent parentheses")
				}
				// При этом открывающая скобка удаляется из стека, но в выходную строку не добавляется.
				if *stackOp.Peek() == "(" {
					stackOp.Pop()
					np--
				}
				// Если после этого шага на вершине стека оказывается символ функции, выталкиваем его в выходную строку.

				if stackOp.Peek() != nil && utils.IsContainsFunc(*stackOp.Peek()) {
					v := stackOp.Pop()
					resQueue.Enqueue(*v)
				}
			}
			/*default:
			return "", fmt.Errorf("syntax error")*/
		}
	}
	for stackOp.Peek() != nil {
		v:= stackOp.Pop()
		resQueue.Enqueue(*v)
	}
	/*if np > 0 {
		return "", fmt.Errorf("syntax error")
	}*/
	return resQueue.All(), nil
}

func evaluateOperator(operator string, stackOp Stack) (float64, Stack, error) {
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

func validate(value string) bool {
	var count = 0
	for i := 0; i < len(value); i++ {
		if value[i] == '(' {
			count++
		} else if value[i] == ')' {
			if count < 0 {
				return false
			}
			count--
		}
	}

	if count == 0 {
		return true
	}
	return false
}

func isNeedConvert(value string) bool {
	for i := 0; i < len(value); i++ {
		if value[i] == '(' {
			return true
		}
	}
	return false
}

func prior(s string) int {
	switch s {
	case "(":
		return 1
	case "+", "-":
		return 2
	case "*", "/":
		return 3
	default:
		return 0
	}
}