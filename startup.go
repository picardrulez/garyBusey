package main

import (
	"log"
)

func startup() {
	adminpass := getPassword("admin")
	passres := checkPass(DEFAULTAUTH, adminpass)
	if passres != true {
		log.Println("admin pass does not match, resetting pass")
		hashedpass := encryptPassword(DEFAULTAUTH)
		passres := dbPost("admin", hashedpass, "pass")
		if passres != 0 {
			log.Println("an error occured creating password")
			return
		}

	}
}
