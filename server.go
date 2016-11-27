package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	rc "github.com/girish946/gorcontrol/cmdhandler"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/templates/{path}", pathHandler)
	router.HandleFunc("/handle/{command}", handleCommand)
	fmt.Printf("server starting at http://127.0.0.1:8080\n")
	http.ListenAndServe(":8080", router)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("request from ", r.RemoteAddr, " ", r.URL)
	data, _ := rc.Read("templates/index.html")
	fmt.Fprintf(w, data)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file_path := vars["path"]
	data, _ := rc.Read("templates/" + file_path)
	fmt.Fprintf(w, data)

}

func handleCommand(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	log.Info("request from ", r.RemoteAddr, " ", r.URL)
	commands := map[string]func(){
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
	}
	execute, ok := commands[vars["command"]]
	if ok {
		execute()
		fmt.Fprintf(w, "ok")
	} else {
		log.Error("command not found")
		fmt.Fprintf(w, "command not found")
	}
}
