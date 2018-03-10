package jsonwriter

import (
	"encoding/json"
	"fmt"
	"github.com/haisum/tcbuildmonitor/pkg/tc/build"
	"io"
	"time"
)

type jsonwriter struct {
	w io.Writer
}

func New(w io.Writer) *jsonwriter {
	return &jsonwriter{w}
}

func (j *jsonwriter) Alert(b build.Build) error {
	o := struct {
		Timestamp string
		Build     build.Build
	}{
		time.Now().String(),
		b,
	}
	v, err := json.Marshal(o)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(j.w, "%s\n", v)
	return err
}
