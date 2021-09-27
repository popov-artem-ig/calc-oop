package main

import (
	"calc-oop/structures"
	"calc-oop/utils"
	"fmt"
)

type Converter struct {

}

func Convert(str string) ([]string, error) {
	np := 0                        // Инкремент открытых скобок
	stackOp := structures.Stack{}  // Стек операторов
	resQueue := structures.Queue{} // Результирующая очередь
	strIn := str                   // Выражение для разбора
	hoarder := ""                  // Накопитель для сборки/обработки дробных и больших чисел

	for i := range strIn {
		c := strIn[i]
		//utils.DebugLog(string(c), stackOp, resQueue)

		if utils.IsDigit(c) || c == '.' {
			hoarder += string(c)
			continue
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

				if stackOp.Peek() != nil && isContainsFunc(*stackOp.Peek()) {
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
