package component

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	cnf "snuggie12/eida/config"
)

type Input struct {
	conf *cnf.Config
}

func NewInput(conf *cnf.Config) *Input {
	return newInput(conf)
}

func newInput(conf *cnf.Config) *Input {
	 i := &Input {
		conf: conf,
	 }

	 return i
}

func (i *Input) handleWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("headers: %v\n", r.Header)

	_, err := io.Copy(os.Stdout, r.Body)
	if err != nil {
		log.Println(err)
		return
	}
}

func (i *Input) StartInput() {
	inputConf := i.conf.InputConfig
	serverAddress := fmt.Sprintf("%v:%v", inputConf.Address, inputConf.Port)
  log.Println("starting server")
	http.HandleFunc("/webhook", i.handleWebhook)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}
