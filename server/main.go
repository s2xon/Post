package main

import (
	"log"
	"net/http"
	"path"

	"root/server/api/facebook"
)

func Hanlder(w http.ResponseWriter, r *http.Request) {
	root := "./app"
	http.ServeFile(w, r, path.Join(root, r.URL.Path))
}



func FBHanlder(w http.ResponseWriter, r *http.Request) {
	root := "./app"

	

	acc_tkn := fb.ACC_TKN(w, r)
	app_tkn := fb.APP_TKN()

	// log.Println(acc_tkn)
	
	// log.Println(app_tkn)

	log.Println(fb.USER(acc_tkn, app_tkn))

	http.ServeFile(w, r, path.Join(root, "index.html"))
 
}

func main() {

	http.HandleFunc("/", Hanlder)
	http.HandleFunc("/facebook", FBHanlder)
	log.Println("https://localhost:3000/")
	log.Fatal(http.ListenAndServeTLS(":3000", "./server.crt", "./server.key", nil))
}
