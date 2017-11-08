package hamming

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//Hamming es el programa de prueba
func Hamming() {
	fmt.Println("-Hamming-")

	codificacion := 32
	fmt.Println(fmt.Sprintf("Codificacion: %v, Bits Paridad: %v, Bits Información: %v", 
		codificacion,
		bitsParidad(codificacion),
		bitsInformacion(codificacion)))
	pruebaHGR(32)
	return
}

func pruebaMatriz() {
	k := [][]bool{{false, true, true}, {true, true, false}}
	ma := Matriz{datos: k}
	r := [][]bool{{false, true}, {true, true}, {false, true}}
	mb := Matriz{datos: r}
	error,m := ma.Multiplicar(&mb)
	if error{
		fmt.Println("Error al multiplicar")
	}
	fmt.Println("Matriz Resultante:")
	fmt.Println(m.ToString())
}

func pruebaHGR(codificacion int){
	fmt.Println("H:")
	h:=h(codificacion)
	strinH:=h.ToString()
	fmt.Println(strinH)

	fmt.Println("G:")
	g:=g(codificacion)
	strinG:=g.ToString()
	fmt.Println(strinG)
	
	fmt.Println("R:")
	r:=r(codificacion)
	strinR:=r.ToString()
	fmt.Println(strinR)

	
	fmt.Println("Entrada:")
	c:=matrizEntradaPrueba(codificacion)
	strinC:=c.ToString()
	fmt.Println(strinC)

	error,codificada := g.Multiplicar(c)

	AgregarError(&codificada)	
	error,posicion:=TieneError(&codificada)
	fmt.Println(fmt.Sprintf("Tienen errores:%v en %v",error,posicion))

	CorregirError(&codificada)
	error,posicion=TieneError(&codificada)
	fmt.Println(fmt.Sprintf("Tienen errores:%v en %v",error,posicion))

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
	alto := codificacion
	ancho:=bitsParidad(codificacion)
	aux := NuevaMatriz(ancho, alto )
	for i:=0; i<ancho ;i++{
		b:=1
		uno:=false
		for j:=0;j<alto;j++{
			f:=float64(i)
			if int(math.Pow(2,f))==b{
				b=0
				uno = !uno
			}
			b++
			if(uno){ 
				aux.datos[i][j]=true
			}       
		}
	}
	return *aux
}

//g Funcion que crea la matriz generadora 
func g(codificacion int) *Matriz {
        n:=codificacion
        m:=bitsInformacion(codificacion)
        aux:=NuevaMatriz(n,m)
        k:=0
        p:=-1
        for i:=0;i<n;i++{
            if(!esPotenciaDeDos(i+1)){
                aux.datos[i][k]=true
                k++
            }else{
                    p++
                    r:=1
                    uno:=false
                    b:=1;
                    for j:=0;j<m; j++{
                        for ;esPotenciaDeDos(r);{
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
                            aux.datos[i][j]=true
                        }
                        r++                    
                }
            }
        }
        return aux;
}
//r Funcion que crea la matriz decodificadora 
func r(codificacion int) *Matriz {
	n:=bitsInformacion(codificacion)
	m:=codificacion
	aux:=NuevaMatriz(n,m)
	k:=0
	for i:=0;i<m;i++{
		if(!esPotenciaDeDos(i+1)){
			aux.datos[k][i]=true
			k++
		}
	}
	return aux;
}


//Codificar  Cifra el archivo de entrada
func Codificar(operando *Matriz) (bool,Matriz){
	cod := len(operando.datos)
	h:=h(cod)
	error,codificada :=h.Multiplicar(operando)
	if error{
		return true,*operando
	}
	return false,codificada
	
}

//TieneError verifica si tiene error una matriz, y devuelve la posicion del mismo.
func TieneError(operando *Matriz) (bool, int){
	codificacion:= len(operando.datos)
	h := h(codificacion)
	error,sindrome := h.Multiplicar(operando)
	if error {
		fmt.Println("Error al multiplicar")
		return true,-1
	}
	resultB:=sindrome.TieneUnos()
	resultI := -1
	for i,fila:=range sindrome.datos{
		for _,b := range fila{
			if b{
				f := float64(i)
				resultI=resultI+int(math.Pow(2,f))
			}
		}
	}

 return resultB ,resultI
}

//Decodificar descrifra el archivo de entrada
func Decodificar() {

}

//AgregarError a un archivo, si devuelve true es porque agrego error, sino ya habia antes un error
func AgregarError(operando *Matriz) bool{
	tieneError,_ :=TieneError(operando)
	if tieneError{
		return false
	}
	n:= len(operando.datos)
	m:= len(operando.datos[0])
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(n)
	j := r.Intn(m)
	operando.datos[i][j] = !operando.datos[i][j]
	return true
}

//CorregirError toma como entrada una matriz, si la corrige retorna verdadero, sino retorna false
func CorregirError(operando *Matriz)bool {
	tieneError, posicion := TieneError(operando)
	if tieneError{
		if posicion !=-1{
			operando.datos[posicion][0] = !operando.datos[posicion][0]	
			return true
		}
	}
	return false
}

//matrizEntradaPrueba genera una matriz de entrada aleatoria del tamaño de la codificacion
func matrizEntradaPrueba (codificacion int) *Matriz{
	aux := make([]bool ,bitsInformacion(codificacion))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range aux{
		j := r.Intn(100)
		aux[i]= (j%2 == 1)
	}
	m:=MatrizColumna(aux)
	return  m
}