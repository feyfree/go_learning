package test

import "testing"

import "github.com/headfirstgo/gadget"

func PlayList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}

}

func TestTapePlayer(t *testing.T) {
	player := gadget.TapePlayer{}
	mixTape := []string{"Jessie is a girl", "Whip it"}
	PlayList(player, mixTape)
}
