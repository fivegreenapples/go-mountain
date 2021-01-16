package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"unicode/utf8"
)

func INT(val float64) int {
	return int(val)
}

func VAL(in string) int {
	val, err := strconv.Atoi(in)
	if err != nil {
		return 0
	}
	return val
}

func STR_(val int) string {
	return strconv.Itoa(val)
}

func LEFT_(in string, l int) string {
	if l > len(in) {
		return in
	}
	return in[:l]
}

func RIGHT_(in string, l int) string {
	if l > len(in) {
		return in
	}
	return in[len(in)-l:]
}

func MID_(in string, start, l int) string {
	start--
	if start < 0 {
		l += start
		start = 0
	}
	if l <= 0 {
		return ""
	}
	if start >= len(in) {
		return ""
	}
	if start+l > len(in) {
		return in[start:]
	}
	return in[start : start+l]
}

func ASC(val string) int {
	if val == "" {
		return -1
	}
	return int([]rune(val)[0])
}

func CHR_(val int) string {
	val = val & 0xFF
	return string(byte(val))
}

// Not a technically accurate implementation of TAB (it doesn't account for
// current cursor position)
func TAB(val int) {
	fmt.Print(strings.Repeat(" ", val))
}

func CLS() {
	// see https://stackoverflow.com/a/22892171
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J")
	}
}

// RND1 implements RND(1)
func RND_1() float64 {
	return rand.Float64()
}

func OPENIN(filename string) *os.File {
	fh, err := os.Open(filename)
	if err != nil {
		return nil
	}
	return fh
}
func OPENOUT(filename string) *os.File {
	fh, err := os.Create(filename)
	if err != nil {
		return nil
	}
	return fh
}

func CLOSE(fh *os.File) {
	fh.Close()
}

func INPUT_DISC(r io.Reader) string {
	var val string
	buf := make([]byte, 1)
	for {
		n, err := r.Read(buf)
		if n != 1 || err != nil {
			if err != nil {
				log.Printf("Error receiving input: %s", err.Error())
			} else if n != 1 {
				log.Printf("Failed to read single byte from input")
			}
			break
		}

		if buf[0] == '\r' {
			continue
		}
		if buf[0] == '\n' {
			break
		}
		val += string(buf[0])
	}
	return val
}

func INPUT_DISC_NUMERIC(r io.Reader) int {
	return VAL(INPUT_DISC(r))
}

func INPUT() string {
	fmt.Print("?")
	return INPUT_DISC(os.Stdin)
}

func INPUT_NUMERIC() int {
	return VAL(INPUT())
}

func LEN(val string) int {
	return utf8.RuneCountInString(val)
}

func STOP(line int) {
	fmt.Println("STOP at line", line)
	os.Exit(0)
}

func PRINT_DISC(w io.Writer, val string) {
	fmt.Fprintln(w, val)
}

func PRINT_DISC_NUMERIC(w io.Writer, val int) {
	PRINT_DISC(w, strconv.Itoa(val))
}

func PRINT(val string) {
	fmt.Println(val)
}

func PRINT_(val string) {
	fmt.Print(val)
}
