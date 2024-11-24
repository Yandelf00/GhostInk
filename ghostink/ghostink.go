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

type InkDropConfig struct {
	entry_input string
	shade       *string
	tags        []string
}

type SetType struct {
	entry_input string
	shade       *string
	tags        []string
	line_no     string
	func_name   string
	basefile    string
}

var s map[string]SetType

type Options func(*InkDropConfig)

// functional Options Pattern
func InkDrop(opts ...Options) {
	var init_shade = "TODO"
	var myset SetType
	config := InkDropConfig{
		entry_input: "Note : this is only the default message",
		tags:        nil,
		shade:       &init_shade,
	}
	for _, opt := range opts {
		opt(&config)
	}
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Unknown file")
	}
	basefile := filepath.Base(file)
	line_str := strconv.Itoa(line)
	func_name := runtime.FuncForPC(pc).Name()
	key := basefile + line_str + func_name
	key += config.entry_input + *config.shade
	myset = SetType{
		entry_input: config.entry_input,
		shade:       config.shade,
		tags:        config.tags,
		line_no:     line_str,
		func_name:   func_name,
		basefile:    basefile,
	}
	s = make(map[string]SetType)

	s[key] = myset
	fmt.Println(s)
}

func WithEntryInput(entry_input string) Options {
	return func(cfg *InkDropConfig) {
		cfg.entry_input = entry_input
	}
}

func WithTags(tags []string) Options {
	return func(cfg *InkDropConfig) {
		cfg.tags = tags
	}
}

func WithShade(shade *string) Options {
	return func(cfg *InkDropConfig) {
		cfg.shade = shade
	}
}

func DropFile(filename string) {

}

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
