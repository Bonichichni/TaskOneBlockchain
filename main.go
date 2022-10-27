package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	fmt.Println("Задача 1")
	fmt.Println()
	for i := 8; i <= 4096; i *= 2 {
		resultKeysVariant := new(big.Int)
		resultKeysVariant.Exp(big.NewInt(2), big.NewInt(int64(i)), nil)
		fmt.Println(resultKeysVariant)
		fmt.Println()
	}
	fmt.Println("Задача 2")
	fmt.Println()
	for i := 8; i <= 4096; i *= 2 {
		maxNum := new(big.Int)
		maxNum.Exp(big.NewInt(2), big.NewInt(int64(i)-1), nil)
		randomNum, _ := rand.Int(rand.Reader, maxNum)
		fmt.Printf("%x\n", randomNum)
		fmt.Println()
	}

	fmt.Println("Задача 3")
	fmt.Println()
	testValue := big.NewInt(12341212)
	fmt.Println(bruteforceNum(testValue).String())
}

func bruteforceNum(num *big.Int) *big.Int {
	start := time.Now()
	//numString := fmt.Sprintf("%d", num)
	n := big.NewInt(0)
	oneBigInt := big.NewInt(1)
	for {
		// nString := fmt.Sprintf("%d", n)
		if n.Cmp(num) != 0 {
			n.Add(n, oneBigInt)
		} else {
			break
		}
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed.Milliseconds(), "Milliseconds")
	return n
}
