package ghostink

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	c "github.com/Yandelf00/GhostInk/constants"
)

func Haunt(message ...string) {
	// Prints the file name, line number, function name, and timestamp of where this method is called.
	//
	// Parameters :
	// - message (str) : Optional message to print before the file information.
	//
	// Prints the file information along with the message if provided, including the file name, line number, function name, and timestamp.

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("unknown file")
	}
	basefile := filepath.Base(file)
	line_str := strconv.Itoa(line)
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
