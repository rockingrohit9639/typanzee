package typanzee

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func TerminalSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}

	height, err := strconv.Atoi(sArr[0])
	if err != nil {
		log.Fatal(err)
	}

	return width, height

}
