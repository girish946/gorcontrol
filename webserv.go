package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	rc "github.com/girish946/gorcontrol/cmdhandler"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	commands = map[string]func(){
		"show_desktop":     rc.ShowDesktop,
		"show_dashboard":   rc.ShowDashBoard,
		"alt_tab":          rc.AltTab,
		"up":               rc.Up,
		"down":             rc.Down,
		"left":             rc.Left,
		"right":            rc.Right,
		"show_filemanager": rc.ShowFilemanager,
		"close_Window":     rc.AltF4,
		"holdWindows":      rc.HoldWindows,
		"enter":            rc.Enter,
		"backspace":        rc.Backspace,
		"tab":              rc.Tab,
		"pageup":           rc.PageUp,
		"pagedown":         rc.PageDown,
		"f5":               rc.F5,
		"home":             rc.Home,
		"end":              rc.End,
		"power":            rc.Power,
		"escape":           rc.Escape,
	}
)

func main() {
	var dir string

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/"))))
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/handle/{command}", handleCommand)
	http.Handle("/", router)
	fmt.Printf("server starting at http://0.0.0.0:8100\n")
	http.ListenAndServe(":8100", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("request from ", r.RemoteAddr, " ", r.URL)
	data, _ := rc.Read("templates/index.html")
	fmt.Fprintf(w, data)
}

func handleCommand(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	log.Info("request from ", r.RemoteAddr, " ", r.URL)
	execute, ok := commands[vars["command"]]
	if ok {
		execute()
		http.Redirect(w, r, "/", 302)
	} else {
		log.Error("command not found")
		fmt.Fprintf(w, "command not found")
	}
}
