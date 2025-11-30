package main

import "errors"

func coutElement(arr []int) map[int]int {
	coutMap := make(map[int]int)
	for _, num := range arr {
		coutMap[num]++
	}
	return coutMap
}

type calculateFunc func(a, b float64) (float64, error)

func calc(opreation string) calculateFunc {
	switch opreation {
	case "add":
		return func(a, b float64) (float64, error) {
			return a + b, nil
		}
	case "subtract":
		return func(a, b float64) (float64, error) {
			return a - b, nil
		}
	case "multiply":
		return func(a, b float64) (float64, error) {
			return a * b, nil
		}
	case "divide":
		return func(a, b float64) (float64, error) {
			if b == 0 {
				return 0, errors.New("除数不能为零")
			}
			return a / b, nil
		}
	default:
		return nil
	}

}
