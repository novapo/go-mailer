package main

import (
	"flag"
	"strings"
)

var (
	port       int
	username   string
	password   string
	smtpHost   string
	smtpPort   int
	recipients Recipients
)

type Recipients []string

func (r *Recipients) Set(rcpt string) error {
	*r = strings.Split(rcpt, ",")
	return nil
}

func (r *Recipients) String() string {
	return strings.Join(*r, ",")
}

func usage() {
	var usage string

	usage = "listening port"
	flag.IntVar(&port, "port", 8080, usage)
	flag.IntVar(&port, "p", 8080, usage+" (shorthand)")

	usage = "SMTP username"
	flag.StringVar(&username, "username", "", usage)
	flag.StringVar(&username, "u", "", usage+" (shorthand)")

	usage = "SMTP password"
	flag.StringVar(&password, "password", "", usage)

	usage = "SMTP host"
	flag.StringVar(&smtpHost, "host", "smtp.googlemail.com", usage)

	usage = "SMTP port"
	flag.IntVar(&smtpPort, "smtp-port", 587, usage)

	usage = "Comma-separated list of mail recipients"
	flag.Var(&recipients, "recipients", usage)

	flag.Parse()
}
