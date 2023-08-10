package main

import "strconv"

func calculation(a string, b string, symbol string) string{
    answer := 0

	i, err := strconv.Atoi(a)

    if err != nil {
        // ... handle error
        panic(err)
    }

	j, err := strconv.Atoi(b)

    if err != nil {
        // ... handle error
        panic(err)
    }
    
	switch{
	case symbol == "+": answer = i + j
	case symbol == "-": answer = i - j
	case symbol == "*": answer = i * j
	case symbol == "%": answer = i / j
	}

    //fmt.Println(strconv.Itoa(answer))

	return strconv.Itoa(answer)
	}


func GoodOutput(a string) string{
	// 50 symbols
	len_of_a := len(a)

	if len_of_a < 50{
		len_of_a = 50 - len_of_a
	}else{len_of_a = 0}
    
	for i := 0; i < len_of_a; i++{
		a += " "
	}

	return a
}