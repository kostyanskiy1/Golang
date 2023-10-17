package main

import "fmt"

type Player struct {
	model string
}

func (pl *Player) Play(audioType, filename string) {
	if audioType == "wav" {
		fmt.Println(pl.model, "plays music type", audioType, "from", filename)
	} else {
		fmt.Printf("%s plays invalid media. %s format not supported\n", pl.model, audioType)
	}
}

type MediaPlayer interface { //есть интерфейс MediaPlayer с методом Play
	Play(audioType, filename string)
}

type VLCPlayer struct{} //мы хотим представить новый, более шикарный VLCPlayer, который может воспроизводить как аудио-, так и видеофайлы.

func (v *VLCPlayer) PlayVideo(filename string) {
	fmt.Printf("Playing video file. File: %s\n", filename)
}

func (v *VLCPlayer) PlayAudio(filename string) {
	fmt.Printf("Playing audio file. File: %s\n", filename)
}

type MediaAdapter struct { //мы теперь делаем MediaAdapter. Он реализует интерфейс MediaPlayer, но внутри использует VLCPlayer.
	advancedMusicPlayer VLCPlayer
}

func (m *MediaAdapter) Play(audioType string, fileName string) {
	if audioType == "vlc" {
		m.advancedMusicPlayer.PlayVideo(fileName)
	} else if audioType == "mp4" {
		m.advancedMusicPlayer.PlayAudio(fileName)
	} else {
		fmt.Printf("invalid media. %s format not supported\n", audioType)
	}
}

func Write(m MediaPlayer, audioType, fileName string) { //Объявляем функцию Write, которая берёт любой объект, удовлетворяющий интерфейсу
	m.Play(audioType, fileName)
}

func main() {

	a1 := Player{model: "panasonic"}

	a2 := MediaAdapter{advancedMusicPlayer: VLCPlayer{}}

	Write(&a1, "mp4", "littleFile")
	Write(&a1, "wav", "littleFile")

	Write(&a2, "mp4", "Bigfile")
	Write(&a2, "wav", "Bigfile")
}
