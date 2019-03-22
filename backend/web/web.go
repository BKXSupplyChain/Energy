package web

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"github.com/BKXSupplyChain/Energy/db"
	"github.com/BKXSupplyChain/Energy/types"
	"github.com/BKXSupplyChain/Energy/process"
	"hash/crc64"
	"log"
	"net/http"
	"time"
	"math/big"
	"strconv"
	"fmt"
	"os/exec"
	"net/url"
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
	i, err := strconv.ParseUint(r.Form.Get("price"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	proposal.Price = uint64(i)
	i, err = strconv.ParseUint(r.Form.Get("relerror"), 10, 16)
	if err != nil {
		log.Fatal(err)
	}
	proposal.RelError = uint16(i)
	bigI := new(big.Int)
	bigI, errs := bigI.SetString(r.Form.Get("abserror"), 10)
	proposal.AbsError = *bigI
	if errs {
		log.Fatal(errs)
	}
	i, err = strconv.ParseUint(r.Form.Get("durability"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	proposal.TTL = uint64(i)
	bigI = new(big.Int)
	bigI, errs = bigI.SetString(r.Form.Get("amount"), 10)
	proposal.TotalAmount = *bigI
	if errs {
		log.Fatal(err)
	}
	user, err := getUser(r)
	if err != nil {
		log.Fatal(err)
	}
	id := make([]byte, 8)
	binary.LittleEndian.PutUint64(id, crc64.Checksum([]byte(user.Username), crc64.MakeTable(crc64.ECMA)))
	if db.Add(&proposal, string(id)) != nil { // здесь я не разобрался пока
		log.Println("Failed to add propose")
	}
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	proposal.ID = string(id)
	//proposal.ID = fmt.Sprintf("%x.%x.%x.%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:]);
	sendProposal(proposal, r.Form.Get("socketId")/*, здесб идет публичный ключ*/)
	http.Redirect(w, r, "/main", 307)
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
