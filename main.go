package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/lyubenblagoev/email"
	"github.com/lyubenblagoev/go-sendmail/config"
)

var (
	attachment = flag.String("a", "", "attachments (comma-separated)")
	bcc        = flag.String("b", "", "bcc-addresses (comma-separated)")
	cc         = flag.String("c", "", "cc-addresses (comma-separated)")
	from       = flag.String("r", "", "from-address")
	bodyFile   = flag.String("q", "", "file containing the message body")
	subject    = flag.String("s", "", "subject")
	server     = flag.String("srv", "", "email server (host:port)")
	configFile = flag.String("cfg", filepath.Join(filepath.Dir(os.Args[0]), "config.json"), "configuration file")
)

func main() {
	flag.Usage = func() {
		command := filepath.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage: %s [options] to-address\n\n", command)
		fmt.Fprintf(os.Stderr, "If -b <body_file> is not provided the input is taken from stdin.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("to-address is required")
	}

	body := readMessageBody(*bodyFile)
	message := email.NewMessage(*subject, body)

	message.To = args
	if *cc != "" {
		message.Cc = toSlice(*cc)
	}
	if *bcc != "" {
		message.Bcc = toSlice(*bcc)
	}

	if *attachment != "" {
		attachments := toSlice(*attachment)
		for _, att := range attachments {
			err := message.Attach(att)
			if err != nil {
				panic(err)
			}
		}
	}

	conf, err := config.Parse(*configFile)
	if err != nil {
		panic(err)
	}

	if *from != "" {
		message.From = *from
	} else {
		message.From = conf.Sender
	}

	if *server != "" {
		err = email.Send(*server, nil, message)
	} else {
		err = email.Send(conf.Server, nil, message)
	}
	if err != nil {
		panic(err)
	}
}

func toSlice(str string) []string {
	return strings.Split(str, ",")
}

func readMessageBody(filename string) string {
	var bytes []byte
	var err error
	if filename != "" {
		bytes, err = ioutil.ReadFile(filename)
	} else {
		bytes, err = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(bytes)
}
