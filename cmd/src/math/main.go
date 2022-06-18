package mainMath

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

func userInput() {

}
func add() {
	var add int
}

func main1() { // Boilerplate code so VSCode doesn't delete my "unused imports"
	const name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())

}

func main2() { // Boilerplate code so VSCode doesn't delete my "unused imports"
	x := math.Abs(-2)
	fmt.Printf("%.1f\n", x)

	y := math.Abs(2)
	fmt.Printf("%.1f\n", y)
}

func main3() { // Boilerplate code so VSCode doesn't delete my "unused imports"
	if _, err := io.WriteString(os.Stdout, "Hello World"); err != nil {
		log.Fatal(err)
	}

}
