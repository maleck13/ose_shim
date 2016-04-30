package api

import (
	"net/http"
	"os/exec"
	"encoding/json"
	"io"
	"os"
	"github.com/Sirupsen/logrus"
	"fmt"
)


type flushWriter struct {
	f http.Flusher
	w io.Writer
}

//override the write method so that we flush immediately
func (fw *flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	if fw.f != nil {
		fw.f.Flush()
	}
	return
}

func isAuthed(headers http.Header)bool{
	auth := os.Getenv("auth")
	if "" == auth {
		logrus.Info("no auth set no requests allowed. Set auth in the env")
		return false
	}
	hAuth := headers.Get("x-auth")
	return auth == hAuth
}

//Pull docker images stream result back to client
func ImagePull(rw http.ResponseWriter, req *http.Request)HttpError{
	if ! isAuthed(req.Header){
		return NewHttpError(fmt.Errorf("not authed"),http.StatusUnauthorized)
	}
	fw := flushWriter{w: rw}
	if f, ok := rw.(http.Flusher); ok {
		fw.f = f
	}
	decoder := json.NewDecoder(req.Body)
	images := make([]string,0)
	if err := decoder.Decode(&images); err != nil{
		return NewHttpError(err,http.StatusBadRequest)
	}
	for _,img := range images{
		cmd := exec.Command("docker","pull", img)
		cmd.Stdout = &fw
		cmd.Stderr = &fw
		if err := cmd.Run(); err != nil{
			return NewHttpError(err,http.StatusInternalServerError)
		}
	}
	return nil
}
