package api

import (
	"net/http"
	"os/exec"
	"encoding/json"
	"io"
	"os"
	"github.com/Sirupsen/logrus"
	"fmt"
	"github.com/maleck13/ose_shim/config"
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
	creds := newDockerCredentials(req.Header)
	if err := dockerLogin(creds); err != nil{
		return NewHttpError(fmt.Errorf("not authed. Docker login failed " + err.Error()),http.StatusUnauthorized)
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

type dockerCredentials struct {
	User string
	Pass string
}

func newDockerCredentials(headers http.Header)dockerCredentials{
	var (
		user,pass string
	)
	if nil != headers{
		user = headers.Get("x-docker-user")
		pass = headers.Get("x-docker-pass")
	}
	if "" != user && "" != pass{
		return dockerCredentials{User:user,Pass:pass}
	}
	conf := config.Conf
	user = conf.GetDockerUser()
	pass = conf.GetDockerPass()
	return dockerCredentials{User:user,Pass:pass}
}

func dockerLogin(credentials dockerCredentials)error{
	cmd := exec.Command("docker", "login", "-p ", credentials.Pass, "-u ",credentials.User)
	var out []byte
	var err error
	if out,err = cmd.Output(); err != nil{
		return err
	}
	fmt.Println(out)
	return nil
}