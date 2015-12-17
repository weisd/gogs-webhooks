package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func Webhook(w http.ResponseWriter, r *http.Request) {
	ret := "ok"
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
			ret = err.Error()
		}
		fmt.Fprintf(w, ret)
	}()

	params := r.URL.Query()
	key := params.Get("k")
	Rep, ok := Setting.Reps[key]
	if !ok {
		err = fmt.Errorf("rep not found")
		return
	}

	log.Println(Rep)

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	fmt.Println(string(data))
	var v GogsHookRequest

	err = json.Unmarshal(data, &v)
	if err != nil {
		return
	}

	fmt.Println(v)

	pushUser := v.Pusher.Name
	pushRef := v.Ref
	pushSecret := v.Secret

	if pushSecret != Rep.Secret {
		err = fmt.Errorf("secret auth failed")
		return
	}

	if pushRef != Rep.Ref {
		err = fmt.Errorf("Ref auth failed")
		return
	}

	authUserOk := false

	for _, u := range Rep.AllowUser {
		if pushUser == u {
			authUserOk = true
		}
	}

	if !authUserOk {
		err = fmt.Errorf("user auth failed")
		return
	}

	result, err := gitPullSrc(Rep.SrcPath)
	if err != nil {
		return
	}

	fmt.Fprintf(w, result)
}

func gitPullSrc(path string) (string, error) {
	os.Chdir(path)
	out, err := exec.Command("git", "pull").Output()
	return string(out), err
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}

var tomlPath string

func init() {
	defaultToml, _ := os.Getwd()
	defaultToml = filepath.Join(defaultToml, "./src/cmd/main/app.toml")
	flag.StringVar(&tomlPath, "c", defaultToml, "-c /path/to/app.toml config gile")
}

func main() {
	flag.Parse()

	if err := InitConfig(tomlPath); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hooks", Webhook)
	mux.HandleFunc("/", Hello)

	log.Fatal(http.ListenAndServe(Setting.Listen, mux))
}
