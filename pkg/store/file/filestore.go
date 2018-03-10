package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type file struct {
	Path string
}

func New(path string) (*file, error) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return nil, err
	}
	return &file{path}, nil
}

func (f *file) Set(k int, v int64) error {
	fd, err := os.Create(f.Path + string(os.PathSeparator) + strconv.Itoa(k))
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = fd.WriteString(strconv.FormatInt(v, 10))
	return err
}

func (f *file) Get(k int) (int64, error) {
	fd, err := os.Open(f.Path + string(os.PathSeparator) + strconv.Itoa(k))
	if err != nil {
		return 0, err
	}
	defer fd.Close()
	b, err := ioutil.ReadAll(fd)
	return strconv.ParseInt(fmt.Sprintf("%s", b), 10, 64)
}
