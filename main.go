package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	black   = "\033[30m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"
	reset   = "\033[0m"
)

type Option struct {
	color  string
	target string
}

func (opt *Option) Colorize(s string) (res string) {
	for {
		idx := strings.Index(s, opt.target)
		if idx == -1 {
			if s != "" {
				res += s
			}
			break
		}
		res += s[:idx] + opt.color + opt.target + reset
		if len(s) < idx+len(opt.target)+1 {
			break
		}
		s = s[idx+len(opt.target):]
	}
	return res
}

var flagColor string
var flagTarget string

func init() {
	flag.StringVar(&flagColor, "color", "red", "string color")
	flag.StringVar(&flagColor, "c", "red", "string color (shorthand)")
	flag.StringVar(&flagTarget, "target", "", "target string")
	flag.StringVar(&flagTarget, "t", "", "target string (shorthand)")
}

func getColor(c string) string {
	switch c {
	case "black":
		return black
	case "red":
		return red
	case "green":
		return green
	case "yellow":
		return yellow
	case "blue":
		return blue
	case "magenta":
		return magenta
	case "cyan":
		return cyan
	case "white":
		return white
	default:
		break
	}
	return ""
}

func NewOption(c, target string) (*Option, error) {
	color := getColor(c)

	if target == "" {
		return nil, errors.New("Error: target is none")
	}

	return &Option{color: color,
		target: target,
	}, nil
}

func main() {
	flag.Parse()

	opt, err := NewOption(flagColor, flagTarget)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Printf("%s", opt.Colorize(str))
	}
}
