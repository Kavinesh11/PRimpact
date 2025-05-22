package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pr-impact-analyzer/internal/analyzer"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nChoose a function to test:")
		fmt.Println("1. Test identifyModule")
		fmt.Println("2. Test isRiskyChange")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice (1-3): ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter a file path (e.g., internal/analyzer/analyzer.go): ")
			scanner.Scan()
			path := scanner.Text()
			module := analyzer.IdentifyModule(path)
			fmt.Printf("Module identified: %q\n", module)

		case "2":
			fmt.Print("Enter a filename (e.g., go.mod): ")
			scanner.Scan()
			filename := scanner.Text()
			fmt.Print("Enter status (modified/removed): ")
			scanner.Scan()
			status := scanner.Text()
			isRisky := analyzer.IsRiskyChange(filename, status)
			fmt.Printf("Is risky change: %v\n", isRisky)

		case "3":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
