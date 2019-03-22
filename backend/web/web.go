package web

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/BKXSupplyChain/Energy/backend/conf"
	"github.com/BKXSupplyChain/Energy/db"
	"github.com/BKXSupplyChain/Energy/types"
	"hash/crc64"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"time"
)

func serveFile(path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	}
}

func UsernameToId(name string) string {
	id := make([]byte, 8)
	binary.LittleEndian.PutUint64(id, crc64.Checksum([]byte(name), crc64.MakeTable(crc64.ECMA)))
	return string(id)
}

func getUser(r *http.Request) (types.UserData, error) {
	name, err := r.Cookie("username")
	password, err := r.Cookie("password")
	if err != nil {
		return types.UserData{}, errors.New("Auth error")
	}
	log.Println("DATA GOT ", name, password)
	id := UsernameToId(name.Value)
	var user types.UserData
	if db.Get(&user, string(id)) != nil {
		return types.UserData{}, errors.New("No such user")
	}
	if user.Username != name.Value || user.PasswordHash != sha256.Sum256([]byte(password.Value)) {
		return types.UserData{}, errors.New("Wrong password")
	}
	return user, nil
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	http.SetCookie(w, &http.Cookie{
		Name:    "username",
		Value:   r.Form.Get("username"),
		Expires: time.Now().AddDate(0, 0, 1),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "password",
		Value:   r.Form.Get("password"),
		Expires: time.Now().AddDate(0, 0, 1),
	})
	http.Redirect(w, r, "/main", 307)
}

func formatFloat(a float64) string {
	pow := int(math.Floor(math.Log10(a) / 3))
	if -3 <= pow && pow <= 4 {
		return fmt.Sprintf("%f", a*math.Pow10(-pow*3))[:4] + []string{" n", " μ", " m", " ", " k", " M", " G", " T"}[pow+3]
	} else {
		return fmt.Sprintf("%g ", a)
	}
}

func mainData(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r)
	if err != nil {
		w.Header().Add("err", "Auth error")
		http.Redirect(w, r, "/", 307)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("["))
	for _, socketID := range user.Sockets {
		var soc types.SocketInfo
		db.Get(&soc, socketID)
		w.Write([]byte(fmt.Sprintf("[\"%s\"", soc.Alias)))
		w.Write([]byte(fmt.Sprintf(", \"%sJ/s\"", formatFloat(float64(db.TokenGetPower(UsernameToId(user.Username), time.Now().Unix()-10))))))
		if soc.ActiveProposal != "" {
			var prop types.Proposal
			db.Get(&prop, soc.ActiveProposal)
			w.Write([]byte(fmt.Sprintf(", \"%sGwei/J\"", formatFloat(float64(prop.Price)/1e9))))
		} else {
			w.Write([]byte(", \"NC\""))
		}
		w.Write([]byte(fmt.Sprintf("[\"%s\"", socketID)))
	}
	w.Write([]byte("]"))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	_, err := getUser(r)
	if err != nil {
		w.Header().Add("err", "Auth error")
		http.Redirect(w, r, "/", 307)
		return
	}
	pattern, _ := ioutil.ReadFile("./web/static/main.html")
	w.Write([]byte(fmt.Sprintf(string(pattern), conf.GetSelfAddress())))
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	var user types.UserData
	r.ParseForm()
	user.Username = r.Form.Get("username")
	id := UsernameToId(user.Username)
	user.PasswordHash = sha256.Sum256([]byte(r.Form.Get("password")))
	user.PrivateKey = r.Form.Get("privkey")
	log.Println(r.Header)
	if db.Add(&user, string(id)) != nil {
		w.Header().Add("err", "Username is reserved")
		http.Redirect(w, r, "/register", 307)
		return
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
	if r.Header.Get("err") != "" {
		w.Write([]byte("<div class=\"err\">" + r.Header.Get("err") + "</div>"))
	}
	file, _ := os.Open("./web/static/register.html")
	io.Copy(w, file)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("err") != "" {
		w.Write([]byte("<div class=\"err\">" + r.Header.Get("err") + "</div>"))
	}
	file, _ := os.Open("./web/static/login.html")
	io.Copy(w, file)
}

func Serve() {
	http.HandleFunc("/", loginPage)
	http.HandleFunc("/login/impl", loginUser)
	http.HandleFunc("/shooter", serveFile("./web/static/shooter.html"))
	http.HandleFunc("/register", registerPage)
	http.HandleFunc("/register/impl", registerUser)
	http.HandleFunc("/main", mainPage)
	http.HandleFunc("/mainData", mainData)
	http.HandleFunc("/style.css", serveFile("./web/static/style.css"))
	http.ListenAndServe(":80", nil)
}
