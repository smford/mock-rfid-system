package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	_ "sort"
	"strings"
)

var allusers []eehuser

type eehuser struct {
	Name   string `json:"Name"`
	RFID   string `json:"RFID"`
	Status string `json:"Status"`
}

func init() {
	flag.Bool("help", false, "Display help")
	flag.String("indexfile", "index.html", "Default file to present")
	flag.String("listenip", "", "IP address for webservice to bind to")
	flag.String("listenport", "56000", "Port for webservice to listen upon, default 56000")
	flag.Bool("listusers", false, "List users")
	flag.Bool("listdevices", false, "List devices")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	allusers = append(allusers, eehuser{"user1", "aa", "user"})
	allusers = append(allusers, eehuser{"user2", "bb", "user"})
	allusers = append(allusers, eehuser{"admin1", "cc", "admin"})
}

func main() {
	if viper.GetBool("help") {
		displayHelp()
	}

	startWeb(viper.GetString("ListenIP"), viper.GetString("ListenPort"), viper.GetBool("EnableTLS"))
	os.Exit(0)
}

func displayHelp() {
	helpmessage := `
Options:
      --help            Help
      --listenip        IP to listen on
      --listenport      Port to listen on
      --listusers       List users
      --listdevices     List devices
`
	fmt.Printf("%s", helpmessage)

	os.Exit(0)
}

func ValidIP(ip string) bool {
	if net.ParseIP(ip) != nil {
		return true
	}
	showerror("ip is not valid", errors.New(ip), "warn")
	return false
}

func showerror(message string, e error, reaction string) bool {
	if e != nil {
		if strings.ToLower(reaction) == "fatal" {
			log.Fatalf("ERROR: %s:%s", message, e)
		} else {
			log.Printf("%s: %s:%s", strings.ToUpper(reaction), message, e)
		}
		// return true if an error was shown
		return true
	}
	// return false as no error shown
	return false
}

func printFile(filename string, webprint http.ResponseWriter) {
	fmt.Println("Starting printFile")
	texttoprint, err := ioutil.ReadFile(filename)
	if err != nil {
		showerror("cannot open file", errors.New(filename), "warn")
		if webprint != nil {
			http.Error(webprint, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	if webprint != nil {
		fmt.Fprintf(webprint, "%s", string(texttoprint))
	} else {
		fmt.Print(string(texttoprint))
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting handlerIndex")
	printFile(viper.GetString("IndexFile"), w)
}

func handlerUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting handlerUsers")
	givejson := true
	queries := r.URL.Query()
	if strings.ToLower(queries.Get("json")) == "n" {
		givejson = false
	}
	listUsers(w, "someuser", givejson)
}

func listUsers(webprint http.ResponseWriter, username string, printjson bool) {
	if printjson {
		c, err := json.Marshal(allusers)
		showerror("cannot marshal json", err, "warn")
		webprint.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(webprint, "%s", c)
	} else {
		for _, user := range allusers {
			fmt.Fprintf(webprint, "%-17s  %-15s    %s\n", user.Name, user.RFID, user.Status)
		}
	}
}

func handlerGetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting handlerGetUser")
	givejson := true
	queries := r.URL.Query()
	if strings.ToLower(queries.Get("json")) == "n" {
		givejson = false
	}
	getUser(w, strings.ToLower(queries.Get("rfid")), givejson)
}

func getUser(webprint http.ResponseWriter, rfid string, printjson bool) {

	var founduser eehuser

	for i := range allusers {
		if allusers[i].RFID == rfid {
			// Found!
			founduser = allusers[i]
			break
		}
	}

	if printjson {
		c, err := json.Marshal(founduser)
		showerror("cannot marshal json", err, "warn")
		webprint.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(webprint, "%s", c)
	} else {
		fmt.Fprintf(webprint, "%-17s  %-15s    %s\n", founduser.Name, founduser.RFID, founduser.Status)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println("MIDDLEWARE: ", r.RemoteAddr, " ", r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func startWeb(listenip string, listenport string, usetls bool) {
	r := mux.NewRouter()

	if viper.GetString("IndexFile") != "" {
		r.HandleFunc("/", handlerIndex)
	}

	usersRouter := r.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("", handlerUsers)
	usersRouter.Use(loggingMiddleware)

	getuserRouter := r.PathPrefix("/getuser").Subrouter()
	getuserRouter.HandleFunc("", handlerGetUser)
	getuserRouter.Use(loggingMiddleware)

	showerror("Starting HTTP Webserver", errors.New(listenip+":"+listenport), "info")
	err := http.ListenAndServe(listenip+":"+listenport, r)
	showerror("cannot start http server", err, "fatal")

}
