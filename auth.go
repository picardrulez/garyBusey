package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"text/template"
)

const indexPage = `
<form method="post" action="/login">
	<input type="text" id="name" name="name" autofocus="autofocus" placeholder="username">
	<input type="password" id="password" name="password" placeholder="password">
	<button type="submit">Login</button>
</form>
<br/>
`

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		http.Redirect(w, r, "/admin", 302)
	} else {
		fmt.Fprintf(w, "<html><head></head><body bgcolor='#33FF33'>")
		fmt.Fprintf(w, "<center>")
		fmt.Fprintf(w, "<br/><br/><br/>")
		fmt.Fprintf(w, "<img src='/resources/garyBuseyLogo.png'>")
		fmt.Fprintf(w, indexPage)
		fmt.Fprintf(w, "My dark side, my shadow, my lower companion is now in the back room blowing up balloons for kids' parties.")
	}
}

func getUserName(r *http.Request) (userName string) {
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func adminPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		t := template.New("admin Template")
		content := `<img src="/resources/garyBuseyLogo.png">`
		adminPage := Page{Username: userName, Content: content}
		pageTemplate := pageBuilder()
		t, _ = t.Parse(pageTemplate)
		t.Execute(w, adminPage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pass := r.FormValue("password")
	redirectTarget := "/"
	if name != "" && pass != "" {
		realPass := getPassword(name)
		hashcheck := checkPass(pass, realPass)
		if hashcheck {
			setSession(name, w)
			redirectTarget = "/admin"
		}
	}
	http.Redirect(w, r, redirectTarget, 302)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", 302)
}

func setSession(userName string, w http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func getPassword(username string) string {
	user := username
	userPassword := dbGet(user, "pass")
	return userPassword
}

func checkPass(password string, hash string) bool {
	bytePass := []byte(password)
	byteHash := []byte(hash)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePass)
	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}

func encryptPassword(password string) string {
	bytePassword := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("an error occured hashing password")
	}
	return string(hashedPassword)
}
func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
