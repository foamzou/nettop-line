package nettop

import (
	"bufio"
	"os/exec"
	"strings"
	"time"
)


func Start (c chan<- string, params string) {
	cmdList := strings.Split("script -q /dev/null nettop " + params, " ")
	cmd := exec.Command(cmdList[0], cmdList[1:]...)

	// Block the stdin, otherwise the nettop process would take cpu high usage
	// Maybe a bug with nettop?
	stdin, _ := cmd.StdinPipe()
	go func() {
		for {
			time.Sleep(time.Hour)
		}
		// Yap, the stdin never be close
		stdin.Close()
	}()


	outPipe, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(outPipe)

	go func() {
		for scanner.Scan() {
			c <- scanner.Text()
		}
	}()

	cmd.Start()
	cmd.Wait()
}