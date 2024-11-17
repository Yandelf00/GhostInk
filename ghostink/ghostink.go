package ghostink

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	c "github.com/Yandelf00/GhostInk/constants"
	sh "github.com/Yandelf00/GhostInk/shades"
)

func Haunt(message ...string) {
	// Prints the file name, line number, function name, and timestamp of where this method is called.
	//
	// Parameters :
	// - message (str) : Optional message to print before the file information.
	//
	// Prints the file information along with the message if provided, including the file name, line number, function name, and timestamp.

	// runtime.Caller returns program counter, file name and line
	// the ok is a boolean indicating that the information was successfully retrieved
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("unknown file")
	}

	// filepath.Base(filename) gives the base of the pathfile
	// example :
	// if filname = Desktop/repositories/main.go
	// then basefile = main.go
	basefile := filepath.Base(file)
	//converts the line to string
	line_str := strconv.Itoa(line)
	//gives the name of the function where haunt is called
	func_name := runtime.FuncForPC(pc).Name()
	curr_time := time.Now().Format("15:04:05")

	output_text := c.BRIGHT_YELLOW + basefile + c.RESET
	output_text += ":"
	output_text += c.MAGENTA + line_str + c.RESET
	output_text += " in " + c.RED + func_name + "()" + c.RESET
	output_text += " at " + curr_time

	if len(message) > 0 {
		fmt.Println(message[0])
		fmt.Println(c.BRIGHT_YELLOW + "└──" + output_text)
	} else {
		fmt.Println(output_text)
	}
}

func Get_Shades() sh.Shades {
	shades := sh.Shades{TODO: "TODO",
		INFO:  "INFO",
		DEBUG: "DEBUG",
		WARN:  "WARN",
		ERROR: "ERROR"}

	return shades
}

func Drop() {

}
