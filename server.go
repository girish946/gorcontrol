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
	router.HandleFunc("/templates/static/css/{path}", cssHandler)
	router.HandleFunc("/templates/static/js/{path}", jsHandler)
	router.HandleFunc("/templates/static/fonts/{path}", fontsHandler)
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
	//log.Info(file_path)
	log.Info("request from ", r.RemoteAddr, " ", r.URL)
	data, _ := rc.Read("templates/" + file_path)
	fmt.Fprintf(w, data)

}

func jsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file_path := vars["path"]
	//log.Info(file_path)
	log.Info("request from ", r.RemoteAddr, " ", r.URL)
	data, _ := rc.Read("templates/static/js/" + file_path)
	fmt.Fprintf(w, data)

}
func cssHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file_path := vars["path"]
	//log.Info(file_path)
	log.Info("request from ", r.RemoteAddr, " ", r.URL)
	data, _ := rc.Read("templates/static/css/" + file_path)
	fmt.Fprintf(w, data)

}
func fontsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file_path := vars["path"]
	//log.Info(file_path)
	log.Info("request from ", r.RemoteAddr, " ", r.URL)
	data, _ := rc.Read("templates/static/fonts/" + file_path)
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
		//fmt.Fprintf(w, "ok")
		http.Redirect(w, r, "/", 200)
	} else {
		log.Error("command not found")
		fmt.Fprintf(w, "command not found")
	}
}
