package exercises

import (
	"fmt"
	"os"
)

const (
	max int = 100000
)

func main() {
	var n int = 1
	var k int = 1

	var arrayNumber = [max]int{0}
	var arrayPrimeNumber = [max]string{}

	// fill array with One
	for i := 1; i < max; i++ {
		arrayPrimeNumber[i] = "1"
	}

	// calculate prime numbers
	for i := 2; i < max; i++ {
		if arrayPrimeNumber[i] == "1" {
			arrayNumber[n] = i
			n++
			for j := 2 * i; j < max; j += i {
				arrayPrimeNumber[j] = "0"
			}
		}
	}

	// wait for input from user
	for {
		n, err := fmt.Scanf("%d\n", &k)
		if err != nil {
			// fmt.Println(n, err)
		}
		if n > 0 {
			fmt.Fprintf(os.Stdout, "%d\n", arrayNumber[k])
		} else {
			break
		}
	}

}
