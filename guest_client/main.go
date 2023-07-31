package main

import (
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
	"log"
)

func main() {
	mainthread.Init(fn)
}

func fn() {
	//hotkey_game := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyS)
	hotkeyGame := hotkey.New([]hotkey.Modifier{}, hotkey.KeyF5)
	err := hotkeyGame.Register()
	if err != nil {
		log.Fatalf("hotkey: failed to register hotkey: %v", err)
		return
	} else {
		log.Printf("hotkey: %v is registered\n", hotkeyGame)
	}

	i := 0
	for i < 5 {
		<-hotkeyGame.Keydown()
		log.Printf("hotkey: %v is down\n", hotkeyGame)
		<-hotkeyGame.Keyup()
		log.Printf("hotkey: %v is up\n", hotkeyGame)
		i += 1
	}

	hotkeyGame.Unregister()
	log.Printf("hotkey: %v is unregistered\n", hotkeyGame)
}
