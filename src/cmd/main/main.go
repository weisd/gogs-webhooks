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
	"crypto/hmac"
	"crypto/sha256"
	"io"
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

	//log.Println(Rep)

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var v GogsHookRequest

	err = json.Unmarshal(data, &v)
	if err != nil {
		return
	}

	//fmt.Println(v)

	pushUser := v.Pusher.Username
	pushRef := v.Ref
	pushSecret := r.Header.Get("X-Gogs-Signature")

	//https://gogs.io/docs/features/webhook.html
	if pushSecret != getHmacCode(string(data), Rep.Secret) {
		err = fmt.Errorf("secret auth failed")
		return
	}


	if pushRef != Rep.Ref {
		err = fmt.Errorf("Ref auth failed")
		return
	}

	authUserOk := false

	println(pushUser)


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

	fmt.Println(result)
	fmt.Printf("%s pull ok", key)

	fmt.Fprintf(w, result)
}

func gitPullSrc(path string) (string, error) {
	os.Chdir(path)
	out, err := exec.Command("git", "pull").Output()
	return string(out), err
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi man !")
}

var tomlPath string

func init() {
	defaultToml, _ := os.Getwd()
	defaultToml = filepath.Join(defaultToml, "./src/cmd/main/app.toml")
	flag.StringVar(&tomlPath, "c", defaultToml, "-c /path/to/app.toml config gile")
}

//密钥文本将被用于计算推送内容的 SHA256 HMAC 哈希值，并设置为 X-Gogs-Signature 请求头的值。
func getHmacCode(s string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
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
