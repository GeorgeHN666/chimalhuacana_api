/*
This source code was develop by jorge enrique hernandez noyola to the client julio seyde
any modification of the source code it's allowed, however any modification that
it's not done by the developer and causes application issues will no be resposability
of the developer.
The developer give the source code up and running to julio seyde, however any
redistribution of this specific source code will have to preceed legaly.
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const VERSION = "1.0.0"
const SECURE_KEY = "f43efd12b152a139e41550129af3bba9089c85d2b2a7d6873739c7360e7ba410_$"
const METHOD = "SHA-256"

type config struct {
	ENV  string
	PORT string
	KEY  string

	smtp struct {
		Host     string
		Port     string
		User     string
		Password string
	}

	personal struct {
		Host     string
		Port     string
		User     string
		Password string
	}

	ftp struct {
		Add      string
		User     string
		Password string
	}

	db struct {
		URI      string
		Database string
	}
}

type application struct {
	config  config
	InfoL   *log.Logger
	ErrorL  *log.Logger
	Version string
	Method  string
	Key     string
}

func (app *application) StartServer() error {

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%v", app.config.PORT),
		Handler:           app.Routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}
	tr, err := app.CheckS(SECURE_KEY)
	if err != nil {
		return err
	}

	if tr {
		app.InfoL.Printf("API version %s running in %s mode listening in PORT %s", app.Version, app.config.ENV, app.config.PORT)
	} else {
		return err
	}

	return srv.ListenAndServe()
}

func main() {

	var cfg config

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	flag.StringVar(&cfg.ENV, "ENV", "DEVELOPMENT", "Will define the stage of the code")
	flag.StringVar(&cfg.PORT, "PORT", PORT, "Will define the port where the api will listen")
	flag.StringVar(&cfg.KEY, "key", SECURE_KEY, "Key")
	flag.StringVar(&cfg.smtp.Host, "smtpHOST", "ns142.hostgator.mx", "SMTP Host")
	flag.StringVar(&cfg.smtp.Port, "smtpPORT", "465", "SMTP PORT")
	flag.StringVar(&cfg.smtp.User, "smtUSER", "contacto@lachimalhuacana.com ", "SMTP Email")
	flag.StringVar(&cfg.smtp.Password, "smtoPASSWORD", "lachimalhuacana54321", "SMTP PASSWORD")
	flag.StringVar(&cfg.personal.Host, "personalHost", "smtp.gmail.om", "SMTP HOST")
	flag.StringVar(&cfg.personal.Port, "personalPort", "587", "SMTP PORT")
	flag.StringVar(&cfg.personal.User, "personalUser", "antojerialachimalhuacana@gmail.com", "SMTP USER")
	flag.StringVar(&cfg.personal.Password, "personalPassword", "lachimalhuacana12345", "SMTP PASSWORD")
	flag.StringVar(&cfg.ftp.Add, "ftpADD", "ftp.lachimalhuacana.com:21", "FTP Add")
	flag.StringVar(&cfg.ftp.User, "ftpUser", "assetsManager@lachimalhuacana.com", "ftp USER")
	flag.StringVar(&cfg.ftp.Password, "ftpPassword", "lachimalhuacanaadmin", "ftp Password")
	flag.StringVar(&cfg.db.Database, "Database", "principal", "DB ADDRESS")
	flag.StringVar(&cfg.db.URI, "URI", "mongodb+srv://dev:devdev@lachimacluster.g7lkta5.mongodb.net/test", "DB URI")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config:  cfg,
		InfoL:   infoLog,
		ErrorL:  errorLog,
		Version: VERSION,
		Key:     SECURE_KEY,
		Method:  METHOD,
	}

	err := app.StartServer()
	if err != nil {
		app.ErrorL.Fatalln(err.Error())
		return
	}

}
