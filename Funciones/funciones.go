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
	fmt.Println("Titulo de la imagen: ", imagen.Titulo)
	fmt.Println("Formato de la imagen: ", imagen.Formato)
	fmt.Println("Canales de la imagen: ", imagen.Canales)
}

func (audio *Audio) Mostrar() {
	fmt.Println("Titulo del audio: ", audio.Titulo)
	fmt.Println("Formato del audio: ", audio.Formato)
	fmt.Println("Duracion del audio: ", audio.Duracion)
}

func (video *Video) Mostrar() {
	fmt.Println("Titulo del video: ", video.Titulo)
	fmt.Println("Formato del video: ", video.Formato)
	fmt.Println("Frames del video: ", video.Frames)
}
