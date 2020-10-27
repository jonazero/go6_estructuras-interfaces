package main

import (
	"./Funciones"
)

func main() {

	imagen := Funciones.Imagen{"Vacaciones", "jpg", 2}
	audio := Funciones.Audio{"Playa", "flac", 30}
	video := Funciones.Video{"The tree of life", "mkv", 60}
	

	/*
		imagen.Mostrar()
		audio.Mostrar()
		video.Mostrar()
	*/
	fm := Funciones.ContenidoWeb{
		Funciones: []Funciones.Multimedia{
			&imagen,
			&audio,
			&video,
		},
	}
	fm.Mostrar()

	/*
		var opcion uint
		for {
			fmt.Println("\n\nMenu de opciones")
			fmt.Println("1.Agregar imagen")
			fmt.Println("2.Agregar audio")
			fmt.Println("3.Agregar video")
			fmt.Println("4.Mostrar")
			fmt.Println("Salir")
			fmt.Println("Ingresa una opcion")
			fmt.Scanf("%d", &opcion)
			switch opcion {
			case 1:

			case 2:

			}
			if opcion == 5 {
				break
			}
		}
	*/

}
