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
	"math/big"
	"strconv"
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

func concludeContract(w http.ResponseWriter, r *http.Request) {
	var proposal types.Proposal
	r.ParseForm()
	proposal.To = r.Form.Get("neighaddress")
	proposal.Price = new(big.Int)
	proposal.Price.SetString(r.Form.Get("price"), 10)
	proposal.RelError = strconv.ParseUint(r.Form.Get("relerror"), 10, 16)
	proposal.AbsError = new(big.Int)
	proposal.AbsError.SetString(r.Form.Get("abserror"), 10)
	proposal.TTL = strconv.ParseUint(r.Form.Get("durability"), 10, 64)
	proposal.TotalAmount = new(big.Int)
	proposal.TotalAmount.SetString(r.Form.Get("amount"), 10)
	username = r.Cookie("username")
	id := make([]byte, 8)
	binary.LittleEndian.PutUint64(id, crc64.Checksum([]byte(username), crc64.MakeTable(crc64.ECMA)))
	if db.Add(&proposal, id) != nil { // здесь я не разобрался пока
		log.Println("Failed to add propose")
	}
	resp, err := http.Post("localhost:8000", &proposal)
	if err != nil {
		log.Println("Can't send propose")
	}
	http.Redirect(w, r, "main", 307)
}

func contractPage(w http.ResponseWriter, r *http.Request) {
	serveFile("./web/static/contract.html")(w, r)
}

func Serve() {
	http.HandleFunc("/", loginPage)
	http.HandleFunc("/shooter", serveFile("./web/static/shooter.html"))
	http.HandleFunc("/register", registerPage)
	http.HandleFunc("/register/impl", registerUser)
	http.HandleFunc("/main", serveFile("./web/static/main.html"))
	http.HandleFunc("/style.css", serveFile("./web/static/style.css"))
	http.HandleFunc("/contract/impl", concludeContract)
	http.HandleFunc("/contract", contractPage)
	http.ListenAndServe(":80", nil)
}
