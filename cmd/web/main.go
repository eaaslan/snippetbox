package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
	"html/template"
	"log"
	"net/http"
	"os"
	"tekrar/pkg/models/mysql"
	"time"
)

type application struct {
	infoLog *log.Logger
	errorLog *log.Logger
	snippets *mysql.SnippetModel
	users *mysql.UserModel
	templateCache map[string]*template.Template
	session *sessions.Session
}

func main()  {
	addr := flag.String("address",":4000","HTTP network address")
	dsn:=flag.String("dsn","web:pass@/snippetbox?parseTime=true","MySQL data source name")
	secret := flag.String("secret","s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge","Secret key")
	flag.Parse()// if you do not parse this flag you always use port 4000
	infoLog := log.New(os.Stdout,"INFO\t",log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr,"ERROR\t",log.Lshortfile|log.Ldate|log.Ltime)




	db,err := openDB(*dsn)

	if err!=nil{
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache,err := newTemplateCache("./ui/html/")
	if err!=nil{
		log.Fatalln(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime=12*time.Hour
	session.Secure = true

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		snippets: &mysql.SnippetModel{DB: db},
		users:&mysql.UserModel{DB: db},
	templateCache: templateCache,
	session: session}

	tlsConfig :=&tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519,tls.CurveP256},
	}

	infoLog.Printf("Server starting on %s",*addr)

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 5*time.Second,
		WriteTimeout: 10*time.Second,
	    TLSConfig: tlsConfig}

	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB,error) {

		db,err := sql.Open("mysql",dsn)

		if err!=nil{
			return nil, err
		}

		if err=db.Ping(); err!=nil{
			return nil, err
	}
return db,nil
}



