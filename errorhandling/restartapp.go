package errorhandling

import (
	"github.com/kardianos/osext"
	"log"
	"os"
	"syscall"
)

func Restart() {
	log.Println("An error has occured, restarting the app")
	file, _ := osext.Executable()
	syscall.Exec(file, os.Args, os.Environ())
}
