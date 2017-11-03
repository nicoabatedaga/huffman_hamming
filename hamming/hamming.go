package hamming

import (
	"fmt"
	"math"
)

//Hamming es el programa de prueba
func Hamming() {
	fmt.Println("-Hamming-")

	codificacion := 512
	fmt.Println(fmt.Sprintf("Codificacion: %v, Bits Paridad: %v, Bits Información: %v", 
		codificacion,
		bitsParidad(codificacion),
		bitsInformacion(codificacion)))
		
	return
}

func pruebaMatriz() {
	k := [][]bool{{false, true, true}, {true, true, false}}
	ma := Matriz{datos: k}
	r := [][]bool{{false, true}, {true, true}, {false, true}}
	mb := Matriz{datos: r}
	var m Matriz
	m = ma.Multiplicar(mb)
	fmt.Println("Matriz Resultante:")
	fmt.Println(m.ToString())
}

func pruebaHG(codificacion int){
	fmt.Println("H:")
	h:=h(codificacion)
	strinH:=h.ToString()
	fmt.Println(strinH)

	fmt.Println("G:")
	g:=g(codificacion)
	strinG:=g.ToString()
	fmt.Println(strinG)

}

func esPotenciaDeDos( ent int)bool {
	return ((ent != 0) && ((ent & (ent-1)) == 0))
}

func bitsInformacion(ent int) int {
	return(ent - (bitsParidad(ent)))
}

//bitsParidad Devuelve que candidad de bit corresponderian a bits de paridad para 
//determinada codifciacion, si se desea para evitar iterar, podria hacerse un mapeo
//con los 5 valores de codificacion que exiten.
func bitsParidad(ent int) int{
	n:=0
	for i:=2;i<=ent;i=i*2{
		n++;
	}
	if(ent == 7){
		return 3;}
	if(ent == 31){
		return 5;
	}
	return n
}
//h metodo que devuelve la matriz que se multiplica para codificar una entrada
func h(codificacion int) Matriz {
	ancho := codificacion
	alto:=bitsParidad(codificacion)
	aux := NuevaMatriz(ancho, alto )
	for i:=0; i<alto ;i++{
		b:=1
		uno:=false
		for j:=0;j<ancho;j++{
			f:=float64(i)
			if int(math.Pow(2,f))==b{
				b=0
				uno = !uno
			}
			b++
			if(uno){ 
				aux.datos[j][i]=true
			}       
		}
	}
	return *aux
}

func g(codificacion int) *Matriz {
        n:=bitsInformacion(codificacion)
        m:=codificacion
        aux:=NuevaMatriz(n,m)
        k:=0
        p:=-1
        for i:=0;i<m;i++{
            if(!esPotenciaDeDos(i+1)){
                aux.datos[k][i]=true
                k++
            }else{
                    p++
                    r:=1
                    uno:=false
                    b:=1;
                    for j:=0;j<n; j++{
                        for cond:=true;cond;cond=esPotenciaDeDos(r){
							r++;
							f:=float64(p)
                            if b== int(math.Pow(2,f)){
                                uno=!uno
                                b=0
                            }
                            b++
						}
						f:=float64(p)						
                        if b == int( math.Pow(2,f)){
                            uno=!uno
                            b=0
                        }
                        b++
                        if(uno){
                            aux.datos[j][i]=true
                        }
                        r++
                    /*
                    if(b==Math.pow(2,(double)i)){
                        uno= !uno;
                        b=0;
                    }
                    b=b+1;
                    if(uno){ 
                        aux.estructura[j][i]=1;
                    }
                    */
                }
            }
        }
        return aux;
}

//Codificar  Cifra el archivo de entrada
func Codificar() {

}

func tieneError() {

}

//Decodificar descrifra el archivo de entrada
func Decodificar() {

}

//AgregarError a un archivo
func AgregarError() {

}
