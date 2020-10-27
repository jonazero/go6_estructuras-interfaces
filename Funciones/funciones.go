package Funciones

import (
	"fmt"
)

type ContenidoWeb struct {
	Funciones []Multimedia
}

func (fm *ContenidoWeb) Mostrar() {
	for _, f := range fm.Funciones {
		f.Mostrar()
	}
}

type Multimedia interface {
	Mostrar()
}

type Imagen struct {
	Titulo  string
	Formato string
	Canales uint
}
type Audio struct {
	Titulo   string
	Formato  string
	Duracion uint
}
type Video struct {
	Titulo  string
	Formato string
	Frames  uint
}

func (imagen *Imagen) Mostrar() {
	fmt.Println(imagen.Titulo)
	fmt.Println(imagen.Formato)
	fmt.Println(imagen.Canales)
}

func (audio *Audio) Mostrar() {
	fmt.Println(audio.Titulo)
	fmt.Println(audio.Formato)
	fmt.Println(audio.Duracion)
}

func (video *Video) Mostrar() {
	fmt.Println(video.Titulo)
	fmt.Println(video.Formato)
	fmt.Println(video.Frames)
}
