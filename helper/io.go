package helper

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"strings"
)

func WriteFileWithDirectory(path string, data []byte, perm os.FileMode) error {

	s := strings.Split(path, "/")

	// fmt.Println("io.goのs", s)

	var dir string
	if len(s) > 1 {
		dir = strings.Join(s[0:len(s)-1], "/")
		// fmt.Println("dir", dir)
	} else {
		dir = path
	}

	if err := os.MkdirAll(dir, 0777); err != nil {
		return errors.Wrapf(err, "create directory is failed. [%s]", dir)
	}
	// fmt.Println("path", path)
	if err := ioutil.WriteFile(path, data, perm); err != nil {
		return errors.Wrapf(err, "write data to file is failed. [%s]", path)
	}
	return nil
}
