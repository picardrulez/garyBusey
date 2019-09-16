package main

import (
	"io"
	"net/http"
	"text/template"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		t := template.New("createUser Template")
		content := `
		<h1>Create User</h1>
		<form method="post" action="/usercreation">
			<label for="username">User</label>
			<input type="text" id="username" name="username">
			<label for="pass">Password</label>
			<input type="password" id="pass" name="pass">
			<button type="submit">Create User</button>
		</form>
		`
		createUserPage := Page{Username: userName, Content: content}
		pageTemplate := pageBuilder()
		t, _ = t.Parse(pageTemplate)
		t.Execute(w, createUserPage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func usercreation(w http.ResponseWriter, r *http.Request) {
	newuser := r.FormValue("username")
	password := r.FormValue("pass")
	hashedpass := encryptPassword(password)
	passres := dbPost(newuser, hashedpass, "pass")
	if passres != 0 {
		io.WriteString(w, "an error occured creating user password")
		return
	}
	http.Redirect(w, r, "/createUser", 302)
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		t := template.New("listUser Template")
		userList := listBucket("pass")
		sortList := alphabetizeList(userList)
		content := `
		<h1>User List</h1>
		<br/>
		<table <tr><td>USER</td></tr>
		`
		for _, k := range sortList {
			content = content + "<tr><td>" + k + "</td></tr>"
		}
		content = content + "</table>"
		listUserPage := Page{Username: userName, Content: content}
		pageTemplate := pageBuilder()
		t, _ = t.Parse(pageTemplate)
		t.Execute(w, listUserPage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func changePassword(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		userList := listBucket("pass")
		sortList := alphabetizeList(userList)
		t := template.New("changePassword Template")
		content := `
		<h1>Change Passord</h1>
		<br/>
		<form method="post" action="/passEntry">
			<select name="user">
		`
		for _, k := range sortList {
			content = content + `<option value="` + k + `">` + k + `</option>`
		}
		content = content + `
		</select>
		<input type="password" name="newpass" id="newpass">
		<button type="submit">Submit</button>
		</form>
		`
		changePasswordPage := Page{Username: userName, Content: content}
		pageTemplate := pageBuilder()
		t, _ = t.Parse(pageTemplate)
		t.Execute(w, changePasswordPage)
	} else {
		http.Redirect(w, r, "/admin", 302)
	}
}

func passEntry(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	user := r.FormValue("user")
	password := r.FormValue("newpass")
	hashedpass := encryptPassword(password)
	if userName != "" {
		postres := dbPost(user, hashedpass, "pass")
		if postres != 0 {
			io.WriteString(w, "an error occured updating password")
			return
		}
		http.Redirect(w, r, "/changePassword", 302)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		t := template.New("deleteUser Template")
		userList := listBucket("pass")
		sortList := alphabetizeList(userList)
		content := `
		<h1>Delete User</h1>
		<br/>
		<form method="post" action="/userdel">
			<select name="user">
		`
		for _, k := range sortList {
			content = content + `<option value="` + k + `">` + k + `</option>`
		}
		content = content + `
		</select>
		<button type="submit">Delete User</button>
		</form>
		`
		deleteUserPage := Page{Username: userName, Content: content}
		pageTemplate := pageBuilder()
		t, _ = t.Parse(pageTemplate)
		t.Execute(w, deleteUserPage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func userdel(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	passdelres := dbRemove("pass", user)
	if passdelres != 0 {
		io.WriteString(w, "an error occured deleting user")
		return
	}
	http.Redirect(w, r, "/deleteUser", 302)
}
