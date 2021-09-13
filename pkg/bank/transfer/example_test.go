package transfer

import (
	"fmt"
)

func ExampleTotal() {
	fmt.Println(Total(0))
	fmt.Println(Total(5000_00))
	fmt.Println(Total(10000_00))

	//Output:
	//0
	//502500
	//1005000

}
