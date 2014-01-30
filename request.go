package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/mattn/go-scan"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"sort"
)

const debug = false
const api_key = "c442e1b9dc9b5cfb0351a966a7729ac6"
const shared_secret = "e1090daccc20bb2c"
const authUrl = "http://www.rememberthemilk.com/services/auth/"
const restUrl = "https://api.rememberthemilk.com/services/rest/"

type api_param map[string]string

func auth() (string, error) {
	res, err := api(api_param{
		"method": "rtm.auth.getFrob",
	})
	if err != nil {
		return "", err
	}
	var frob string
	err = scan.ScanTree(res, "/rsp/frob", &frob)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("api_key", api_key)
	params.Add("perms", "delete")
	params.Add("frob", frob)
	u := authUrl + sign(params)
	fmt.Println(u)
	openBrowser(u)
	fmt.Println("Hit Enter!")

	var b [1]byte
	os.Stdin.Read(b[:])

	res, err = api(api_param{
		"method": "rtm.auth.getToken",
		"frob":   frob,
	})
	if err != nil {
		log.Fatal(err)
	}
	var token string
	err = scan.ScanTree(res, "/rsp/auth/token", &token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func openBrowser(u string) error {
	cmd := "xdg-open"
	args := []string{cmd, u}
	if runtime.GOOS == "windows" {
		cmd = "rundll32.exe"
		args = []string{cmd, "url.dll,FileProtocolHandler", u}
	} else if runtime.GOOS == "darwin" {
		cmd = "open"
		args = []string{cmd, u}
	}
	cmd, err := exec.LookPath(cmd)
	if err != nil {
		return err
	}
	p, err := os.StartProcess(cmd, args, &os.ProcAttr{Dir: "", Files: []*os.File{nil, nil, os.Stderr}})
	if err != nil {
		return err
	}
	defer p.Release()
	return err
}

func sign(params url.Values) string {
	keys := []string{}
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	s := shared_secret
	query := ""
	for _, key := range keys {
		value := params.Get(key)
		s += key + value
		if query == "" {
			query += "?"
		} else {
			query += "&"
		}
		query += key + "=" + url.QueryEscape(value)
	}
	return query + "&api_sig=" + fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func api(kv map[string]string) (interface{}, error) {
	params := url.Values{}
	for k, v := range kv {
		params.Add(k, v)
	}
	if params.Get("api_key") == "" {
		params.Add("api_key", api_key)
	}
	params.Add("format", "json")
	r, err := http.Get("https://api.rememberthemilk.com/services/rest/" + sign(params))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	var v interface{}
	if debug {
		err = json.NewDecoder(io.TeeReader(r.Body, os.Stdout)).Decode(&v)
		os.Stdout.Write([]byte{'\n'})
	} else {
		err = json.NewDecoder(r.Body).Decode(&v)
	}
	if err != nil {
		return nil, err
	}
	return v, nil
}

func timeline(cfg config) (string, error) {
	res, err := api(api_param{
		"method":     "rtm.timelines.create",
		"auth_token": cfg["auth_token"],
	})
	if err != nil {
		return "", err
	}
	var tl string
	err = scan.ScanTree(res, "/rsp/timeline", &tl)
	if err != nil {
		return "", err
	}
	return tl, nil
}
