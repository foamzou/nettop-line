package nettop

import (
	"bufio"
	"os/exec"
	"strings"
	"time"
)

const DIRTY_DATA_FLAG = "dirty_data_flag"

func Start (c chan<- string, params string) {
	shouldMarkDirty := false

	// The `for true` is keep the sub-process live
	for true {
		// The first output from nettop was not really increase. Need to drop it
		shouldMarkDirty = true
		cmdList := strings.Split("script -q /dev/null nettop "+params, " ")
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
				if shouldMarkDirty {
					c <- DIRTY_DATA_FLAG + scanner.Text()
					shouldMarkDirty = false
				} else {
					c <- scanner.Text()
				}
			}
		}()

		cmd.Start()
		cmd.Wait()
	}
}