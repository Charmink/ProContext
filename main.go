package main

import "fmt"

func main()  {
	var n, k int
	fmt.Print("Size of matrix: ")
	fmt.Scan(&n)// считываем размерность
	fmt.Print("Position of figure in the last array: ")
	fmt.Scan(&k)// считываем позицию
	var first []int64
	var second []int64
	var help []int64 // используется при переназначении массивов
	if n < 1 {
		fmt.Printf("Parametr n = %d is incorrect!", n) // проверка входных данных
		return
	}

	if k > n || k < 1{
		fmt.Printf("Parametr k = %d is incorrect!", k)
		return
	}

	for i := 0; i < n; i ++{ // инициализация используемых массивов(фрагмент шахматной доски)
		help = append(help, 0)
		first = append(first, 0)
		second = append(second, 0)
	}

	first[k - 1] = 1
	for i := 0; i < n - 1; i ++{
		for j := 0; j < n; j ++{
			if j - 1 >= 0{
				second[j] += first[j - 1]
			}
			if j + 1 < n{
				second[j] += first[j + 1]
			}
		}
		copy(first, second) // переназначаем массивы чтобы перейти к следующему шагу
		copy(second, help)

	}
	var cnt int64 = 0
	for _, el := range first{ // считаем результат(сумма элементов верхнего ряда доски)
		cnt += el
	}
	fmt.Printf("Result: %d", cnt)
}
