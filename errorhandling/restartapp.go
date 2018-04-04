package errorhandling

import (
	"github.com/kardianos/osext"
	"log"
	"os"
	"syscall"
)

// Restart restarts the application
func Restart() {
	log.Println("An error has occured, restarting the app")
	file, _ := osext.Executable()
	syscall.Exec(file, os.Args, os.Environ())
}
