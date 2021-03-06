package status

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// Get downloads the package and updates its CanGet status.
func (pkg *Package) Get() (err error) {
	pkg.CanGet = false
	// go get flags:
	//    -d (only download)
	//    -u (update)
	cmd := exec.Command("go", "get", "-d", "-u", pkg.Path)
	cmd.Env = append([]string{"GIT_ASKPASS=/usr/bin/echo"}, os.Environ()...)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Get failed for: %s", pkg.Path)
	}
	pkg.CanGet = true
	return nil
}

// init initializes a custom $GOPATH.
func init() {
	gopath, err := ioutil.TempDir("", "pkg_")
	if err != nil {
		log.Fatalln(err)
	}
	os.Setenv("GOPATH", gopath)
	fmt.Printf("Using the following GOPATH: %q\n", gopath)
}
