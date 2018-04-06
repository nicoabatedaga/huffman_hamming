package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

/*func TestTiempos(t *testing.T) {
	var archivoPrueba = struct {
		archivoEntrada      string
		archivoProtegido    string
		archivoDesprotegido string
	}{
		"./libros/biblia.txt",
		"./libros/biblia.ham",
		"./libros/bibliaDesprotegido.txt",
	}
	path := "tiemposEjecucion"
	crearArchivo(path + "P.csv")
	appendFile(path+"P.csv", "operacion;codificacion;tiempo\n")
	crearArchivo(path + "C.csv")
	appendFile(path+"C.csv", "operacion;codificacion;tiempo\n")
	crearArchivo(path + "D.csv")
	appendFile(path+"D.csv", "operacion;codificacion;tiempo\n")
	var cod []int
	cod = []int{2048}
	for _, c := range cod {
		for i := 0; i < 10; i++ {
			ahora := time.Now()
			error := hamming.ProtegerB(archivoPrueba.archivoEntrada, archivoPrueba.archivoProtegido, c)
			manejoError(error)
			appendFile(path+"P.csv", fmt.Sprintf("progeter;%v;%v\n", c, tiempoStr(time.Now().Sub(ahora))))

			ahora = time.Now()
			hamming.TieneErroresB(archivoPrueba.archivoProtegido)
			appendFile(path+"C.csv", fmt.Sprintf("comprobar;%v;%v\n", c, tiempoStr(time.Now().Sub(ahora))))

			ahora = time.Now()
			error = hamming.DesprotegerB(archivoPrueba.archivoProtegido, archivoPrueba.archivoDesprotegido)
			manejoError(error)
			appendFile(path+"D.csv", fmt.Sprintf("desproteger;%v;%v\n", c, tiempoStr(time.Now().Sub(ahora))))
		}
	}

}
*/
/*
func TestProteccionDesproteccionArchivo(t *testing.T) {
	ahora := time.Now()
	fmt.Println("Test Proteccion-Desproteccion ", ahora)
	var archivosPrueba = []struct {
		archivoEntrada      string
		archivoProtegido    string
		archivoDesprotegido string
	}{
		{"./libros/prueba.txt", "./libros/prueba.ham", "./libros/pruebaDesprotegido.txt"},
		{"./libros/alicia.txt", "./libros/alicia.ham", "./libros/aliciaDesprotegido.txt"},
		{"./libros/biblia.txt", "./libros/biblia.ham", "./libros/bibliaDesprotegido.txt"},
		{"./libros/moby.txt", "./libros/moby.ham", "./libros/mobyDesprotegido.txt"},
		{"./libros/meta.txt", "./libros/metamorfosis.ham", "./libros/metamorfosisDesprotegido.txt"},
		{"./libros/hyde.txt", "./libros/hyde.ham", "./libros/hydeDesprotegido.txt"},
	}
	var codificacionesPosibles = []int{512, 1024, 2048} //, 1024, 2048
	for _, tuplaArchivos := range archivosPrueba {
		for _, codificacion := range codificacionesPosibles {
			ahora := time.Now()
			fmt.Printf("%s\t%s\tCodificacion:%v \n", tiempo(), tuplaArchivos.archivoEntrada, codificacion)
			error := hamming.ProtegerB(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoProtegido+string(codificacion), codificacion)
			manejoError(error)
			fmt.Printf("%s\t\tDesprotejo\n", tiempo())
			error = hamming.DesprotegerB(tuplaArchivos.archivoProtegido+string(codificacion), tuplaArchivos.archivoDesprotegido)
			manejoError(error)
			if !compararArchivos(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoDesprotegido) {
				fmt.Printf("\tError en el %s para la codificacion %d\tduracion:%s\n", tuplaArchivos.archivoEntrada, codificacion, tiempoStr(time.Now().Sub(ahora)))
				t.Errorf("\tError en el %s para la codificacion %d\t\n", tuplaArchivos.archivoEntrada, codificacion)
			} else {
				fmt.Printf("\tExito en el %s para la codificacion %d\tduracion:%s\n", tuplaArchivos.archivoEntrada, codificacion, tiempoStr(time.Now().Sub(ahora)))
			}
		}
		fmt.Println("")
	}
}*/
func tiempo() string {
	ahora := time.Now()
	return fmt.Sprintf("%v' %v'' %v-", ahora.Minute(), ahora.Second(), ahora.Nanosecond())
}
func tiempoStr(ahora time.Duration) string {

	return fmt.Sprintf("%v", ahora.Seconds())

}

/*
func TestCorregirError(t *testing.T) {
	var archivosPrueba = []struct {
		archivoEntrada      string
		archivoProtegido    string
		archivoConError     string
		archivoCorregido    string
		archivoDesprotegido string
	}{
		{"./libros/prueba.txt", "./libros/prueba.ham", "./libros/pruebaConError.ham", "./libros/pruebaCorregido.ham", "./libros/pruebaDesprotegido.txt"},
		{"./libros/alicia.txt", "./libros/alicia.ham", "./libros/aliciaConError.ham", "./libros/aliciaCorregido.ham", "./libros/aliciaDesprotegido.txt"},
		{"./libros/biblia.txt", "./libros/biblia.ham", "./libros/bibliaConError.ham", "./libros/bibliaCorregido.ham", "./libros/bibliaDesprotegido.txt"},
		{"./libros/moby.txt", "./libros/moby.ham", "./libros/mobyConError.ham", "./libros/mobyCorregido.ham", "./libros/mobyDesprotegido.txt"},
		{"./libros/meta.txt", "./libros/meta.ham", "./libros/metaConError.ham", "./libros/metaCorregido.ham", "./libros/metaDesprotegido.txt"},
		{"./libros/hyde.txt", "./libros/hyde.ham", "./libros/hydeConError.ham", "./libros/hydeCorregido.ham", "./libros/hydeDesprotegido.txt"},
	}
	var codificacionesPosibles = []int{512, 1024, 2048} //512, 1024, 2048
	for _, tuplaArchivos := range archivosPrueba {
		for _, codificacion := range codificacionesPosibles {
			error := hamming.ProtegerB(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoProtegido, codificacion)
			manejoError(error)
			hamming.IntroducirError(tuplaArchivos.archivoProtegido, tuplaArchivos.archivoConError)

			hamming.CorregirError(tuplaArchivos.archivoConError, tuplaArchivos.archivoCorregido)
			manejoError(error)
			error = hamming.DesprotegerB(tuplaArchivos.archivoCorregido, tuplaArchivos.archivoDesprotegido)
			manejoError(error)
			if !compararArchivos(tuplaArchivos.archivoEntrada, tuplaArchivos.archivoDesprotegido) {
				t.Errorf("-Error en el %s para la codificacion %d", tuplaArchivos.archivoEntrada, codificacion)
			} else {
				t.Logf("Exito en el %s para la codificacion %d", tuplaArchivos.archivoEntrada, codificacion)
			}
		}
	}
}
*/ /*
func TestProteccionArchivosTieneErrores(t *testing.T) {
	ahora := time.Now()
	fmt.Println("Test TieneErrores ", ahora)
	var archivosPrueba = []struct {
		archivoEntrada string
		archivoSalida  string
	}{
		{"./libros/prueba.txt", "./libros/prueba.ham"},
		{"./libros/alicia.txt", "./libros/alicia.ham"},
		{"./libros/biblia.txt", "./libros/biblia.ham"},
		{"./libros/moby.txt", "./libros/moby.ham"},
		{"./libros/meta.txt", "./libros/meta.ham"},
		{"./libros/hyde.txt", "./libros/hyde.ham"},
	}
	var codificacionesPosibles = []int{512, 1024, 2048} //512, 1024, 2048
	for _, parArchivos := range archivosPrueba {
		for _, codificacion := range codificacionesPosibles {
			fmt.Printf("%s\t%s\tCodificacion:%v \n", tiempo(), parArchivos.archivoEntrada, codificacion)
			erro := hamming.ProtegerB(parArchivos.archivoEntrada, parArchivos.archivoSalida, codificacion)
			manejoError(erro)
			error, bloque, posicion := hamming.TieneErroresB(parArchivos.archivoSalida)
			if error {
				t.Errorf("Error en %s(%d):en el bloque %d en la posición %d", parArchivos.archivoEntrada, codificacion, bloque, posicion)
			} else {
				t.Logf("Exito en el %s para la codificacion %d", parArchivos.archivoEntrada, codificacion)
			}
		}
		fmt.Println("")
	}
}*/
/*
func TestProteccionArchivosAgregarErrores(t *testing.T) {
	ahora := time.Now()
	fmt.Println("Test Agregar Errores ", ahora)
	var archivosPrueba = []struct {
		archivoEntrada  string
		archivoSalida   string
		archivoConError string
	}{
		{"./libros/prueba.txt", "./libros/prueba.ham", "./libros/pruebaConError.ham"},
		{"./libros/alicia.txt", "./libros/alicia.ham", "./libros/aliciaConError.ham"},
		{"./libros/biblia.txt", "./libros/biblia.ham", "./libros/bibliaConError.ham"},
	}
	var codificacionesPosibles = []int{512, 1024, 2048} //512,, 2048
	for _, parArchivos := range archivosPrueba {
		for _, codificacion := range codificacionesPosibles {
			fmt.Printf("%s\t%s\tCodificacion:%v \n", tiempo(), parArchivos.archivoEntrada, codificacion)
			erro := hamming.ProtegerB(parArchivos.archivoEntrada, parArchivos.archivoSalida, codificacion)
			manejoError(erro)
			error, bloque, posicion := hamming.TieneErroresB(parArchivos.archivoSalida)
			if error {
				t.Errorf("-\tError:%s(%d)-(%d,%d)", parArchivos.archivoEntrada, codificacion, bloque, posicion)
			} else {
				t.Logf("+\tExito:%s(%d)", parArchivos.archivoEntrada, codificacion)
			}
			hamming.IntroducirError(parArchivos.archivoSalida, parArchivos.archivoConError)

			error, bloque, posicion = hamming.TieneErroresB(parArchivos.archivoConError)
			if !error {
				t.Errorf("-\t(IE)Error:%s(%d)", parArchivos.archivoEntrada, codificacion)
			} else {
				t.Logf("+\t(IE)Exito:%s(%d)-(%d,%d)", parArchivos.archivoEntrada, codificacion, bloque, posicion)
			}
		}
		fmt.Println("")
	}
}*/
func manejoError(err error) {
	if err != nil {
		panic(fmt.Sprintf("hubo un error, %v", err))
	}
}

func compararArchivos(path1, path2 string) bool {
	archivo1, error := os.Open(path1)
	manejoError(error)
	archivo2, error := os.Open(path2)
	manejoError(error)
	for {
		buffer1 := make([]byte, 64000)
		_, error1 := archivo1.Read(buffer1)
		buffer2 := make([]byte, 64000)
		_, error2 := archivo2.Read(buffer2)
		if error1 != nil || error2 != nil {
			if error1 == io.EOF && error2 == io.EOF {
				return true
			} else if error1 == io.EOF || error2 == io.EOF {
				return false
			} else {
				manejoError(error1)
				manejoError(error2)
			}
		}
		if !bytes.Equal(buffer1, buffer2) {
			fmt.Println("Comparar ")
			for i := 0; i < len(buffer1); i++ {
				if buffer1[i] != buffer2[i] {
					fmt.Printf("\t%b\t \t%v \n\t%b \t%v\n", buffer1[i], buffer1[i], buffer2[i], buffer2[i])
				}
			}
			return false
		}
	}
}
func appendFile(url, cadena string) {
	f, err := os.OpenFile(url, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "%s", cadena)
}

func crearArchivo(url string) {
	_, err := os.Stat(url)
	if os.IsNotExist(err) {
		os.Create(url)
	}
}