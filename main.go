package main

import (
	"fmt"

	"./Funciones"
)

func main() {
	var opcion uint8
	var timagen string
	var fimagen string
	var cimagen uint
	var taudio string
	var faudio string
	var daudio uint
	var tvideo string
	var fvideo string
	var frvideo uint
	for {
		fmt.Println("\n\nMenu de opciones")
		fmt.Println("1.Agregar contenido multimedia")
		fmt.Println("2.Mostrar")
		fmt.Println("3.Salir")
		fmt.Print("Ingresa una opcion: ")
		fmt.Scan(&opcion)
		switch opcion {
		case 1:
			fmt.Print("Ingrese el titulo de la imagen: ")
			fmt.Scan(&timagen)
			fmt.Print("Ingrese el formato de la imagen: ")
			fmt.Scan(&fimagen)
			fmt.Print("Ingrese los canales de la imagen: ")
			fmt.Scan(&cimagen)
			fmt.Print("Ingrese el titulo del audio: ")
			fmt.Scan(&taudio)
			fmt.Print("Ingrese el formato del audio : ")
			fmt.Scan(&faudio)
			fmt.Print("Ingrese la duracion del audio: ")
			fmt.Scan(&daudio)
			fmt.Print("Ingrese el titulo del video: ")
			fmt.Scan(&tvideo)
			fmt.Print("Ingrese el formato del video : ")
			fmt.Scan(&fvideo)
			fmt.Print("Ingrese los frames del video ")
			fmt.Scan(&frvideo)
			break
		case 2:
			imagen := Funciones.Imagen{timagen, fimagen, cimagen}
			audio := Funciones.Audio{taudio, faudio, daudio}
			video := Funciones.Video{tvideo, fvideo, frvideo}
			fm := Funciones.ContenidoWeb{
				Funciones: []Funciones.Multimedia{
					&imagen,
					&audio,
					&video,
				},
			}
			fm.Mostrar()
		case 3:
			break
		}
		if opcion == 3 {
			break
		}
	}
}
