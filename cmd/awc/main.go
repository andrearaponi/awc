package main

import (
	"fmt"
	"os"

	"github.com/andrearaponi/awc/pkg/count"
)

func printHelp() {
	fmt.Println("Usage: awc <file_path> [option]")
	fmt.Println("Options:")
	fmt.Println("-c, --bytes    Print the byte count of the specified file.")
	fmt.Println("-m, --chars    Print the character count of the file. Treats each byte as a character.")
	fmt.Println("-w, --words    Print the word count of the file.")
	fmt.Println("-l, --lines    Print the line count of the file.")
	fmt.Println("--help         Show this usage information.")
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: awc <file_path> [option]")
		fmt.Println("For more information, try 'awc --help'")
		return
	}

	if os.Args[1] == "--help" {
		printHelp()
		return
	}

	filePath := os.Args[len(os.Args)-1]

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var option string
	if len(os.Args) > 2 {
		option = os.Args[1]
	}

	switch option {
	case "-c", "--bytes":
		fmt.Printf("%d %s\n", count.Bytes(data), filePath)
	case "-m", "--chars":
		fmt.Printf("%d %s\n", count.Characters(data), filePath)
	case "-w", "--words":
		fmt.Printf("%d %s\n", count.Words(data), filePath)
	case "-l", "--lines":
		fmt.Printf("%d %s\n", count.Lines(data), filePath)
	case "":
		fmt.Printf("Lines: %d, Words: %d, Characters: %d in %s\n", count.Lines(data), count.Words(data), count.Characters(data), filePath)
	default:
		fmt.Println("Unrecognized option. For help, try 'awc --help'")
	}
}
