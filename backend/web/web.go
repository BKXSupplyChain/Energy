package web

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"github.com/BKXSupplyChain/Energy/db"
	"github.com/BKXSupplyChain/Energy/types"
	"hash/crc64"
	"log"
	"net/http"
	"time"
)

func serveFile(path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	}
}

func getUser(name string, password string) (types.UserData, error) {
	id := make([]byte, 8)
	binary.LittleEndian.PutUint64(id, crc64.Checksum([]byte(name), crc64.MakeTable(crc64.ECMA)))
	var user types.UserData
	if db.Get(&user, string(id)) != nil {
		return types.UserData{}, errors.New("No such user")
	}
	if user.Username != name || user.PasswordHash != sha256.Sum256([]byte(password)) {
		return types.UserData{}, errors.New("Wrong password")
	}
	return user, nil
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	var user types.UserData
	r.ParseForm()
	user.Username = r.Form.Get("username")
	id := make([]byte, 8)
	binary.LittleEndian.PutUint64(id, crc64.Checksum([]byte(user.Username), crc64.MakeTable(crc64.ECMA)))
	user.PasswordHash = sha256.Sum256([]byte(r.Form.Get("password")))
	user.PrivateKey = r.Form.Get("privkey")
	log.Println(r.Header)
	if db.Add(&user, string(id)) != nil {
		w.Header().Add("err", "Username is reserved")
		http.Redirect(w, r, "/register", 307)
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "username",
		Value:   user.Username,
		Expires: time.Now().AddDate(0, 0, 1),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "password",
		Value:   r.Form.Get("password"),
		Expires: time.Now().AddDate(0, 0, 1),
	})
	http.Redirect(w, r, "/main", 307)
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	serveFile("./web/static/register.html")(w, r)
	if r.Header.Get("err") != "" {
		w.Write([]byte("<div class=\"err\">" + r.Header.Get("err") + "</div>"))
	}
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	serveFile("./web/static/login.html")(w, r)
}

func Serve() {
	http.HandleFunc("/", loginPage)
	http.HandleFunc("/shooter", serveFile("./web/static/shooter.html"))
	http.HandleFunc("/register", registerPage)
	http.HandleFunc("/register/impl", registerUser)
	http.HandleFunc("/main", serveFile("./web/static/main.html"))
	http.HandleFunc("/style.css", serveFile("./web/static/style.css"))
	http.ListenAndServe(":80", nil)
}
