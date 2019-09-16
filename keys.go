package main

import (
	"io"
	"net/http"
	"sort"
	"text/template"
)

func addKey(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		t := template.New("Enter Key Tempate")
		content := `
		<h1>Add Key</h1>
		<br/>
		<form method="post" action="/createkey" enctype="multipart/form-data">
			<input type="text" name="key" id="username">
			<input type="text" name="value" id="value">
			<input type="hidden" name="isnew" value="true"> 
			<button type="submit">Submit</button>
		</form>
		`
		enterKeyPage := Page{Username: userName, Content: content}
		pageTemplate := pageBuilder()
		t, _ = t.Parse(pageTemplate)
		t.Execute(w, enterKeyPage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func createkey(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		newkey := r.FormValue("key")
		newvalue := r.FormValue("value")
		encryptedValue := encrypt(cryptKey, newvalue)
		isnew := r.FormValue("isnew")
		if isnew == "true" {
			checkkey := dbGet(newkey, "dataVault")
			if checkkey != "" {
				io.WriteString(w, "Key already exists")
				return
			}
		}
		passres := dbPost(newkey, encryptedValue, "dataVault")
		if passres != 0 {
			io.WriteString(w, "An error occured posting key")
			return
		}
		http.Redirect(w, r, "/addKey", 302)
	} else {
		http.Redirect(w, r, "/admin", 302)
	}
}

func listKeys(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		t := template.New("List Keys Template")
		keyList := listBucket("dataVault")
		sortList := alphabetizeList(keyList)
		content := `
		<h1>List Keys</h1>
		<br/>
		<table><tr><td>KEY</td><td>VALUE</td></tr>
		`
		for _, k := range sortList {
			value := dbGet(k, "dataVault")
			decryptedValue := decrypt(cryptKey, value)
			content = content + "<tr><td>" + k + "</td><td>" + decryptedValue + "</td><td><form method='post' action='/editKey' enctype='multipart/form-data'><input type='hidden' name='key' value='" + k + "'><button type='submit'>Edit</button></form></td></tr>"
		}
		content = content + "</table>"
		listKeysPage := Page{Username: userName, Content: content}
		pageTemplate := pageBuilder()
		t, _ = t.Parse(pageTemplate)
		t.Execute(w, listKeysPage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func editKey(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		key := r.FormValue("key")
		value := dbGet(key, "dataVault")
		decryptedValue := decrypt(cryptKey, value)
		t := template.New("Edit Key Template")
		content := `
		<h1>Edit Key</h1>
		<br/>
		<form method="post" action="/createkey" enctype="multipart/form-data">
			<input type="text" id="key" name="key" value=` + key + ` readonly>
			<input type="text" id="value" name="value" value="` + decryptedValue + `">
			<button type="submit">Submit</button>
		</form>
		`
		editKeyPage := Page{Username: userName, Content: content}
		pageTemplate := pageBuilder()
		t, _ = t.Parse(pageTemplate)
		t.Execute(w, editKeyPage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func removeKey(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		t := template.New("Remove Key Template")
		keyList := listBucket("dataVault")
		sortList := alphabetizeList(keyList)
		content := `
		<h1>Remove Key</h1>
		<br/>
		<form method="post" action="/dropkey" enctype="multipart/form-data">
			<select name="key">
		`
		for _, k := range sortList {
			content = content + `<option value="` + k + `">` + k + `</option>`
		}
		content = content + `
		</select>
		<button type="submit">Submit</button>
		</form>
		`
		removeKeyPage := Page{Username: userName, Content: content}
		pageTemplate := pageBuilder()
		t, _ = t.Parse(pageTemplate)
		t.Execute(w, removeKeyPage)
	} else {
		http.Redirect(w, r, "/admin", 302)
	}
}

func dropkey(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		key := r.FormValue("key")
		keydelres := dbRemove("dataVault", key)
		if keydelres != 0 {
			io.WriteString(w, "")
		}
		http.Redirect(w, r, "/removeKey", 302)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func alphabetizeList(keyList map[string]string) []string {
	s := make([]string, 0)
	for k, _ := range keyList {
		s = append(s, k)
	}
	sort.Strings(s)
	return s
}
