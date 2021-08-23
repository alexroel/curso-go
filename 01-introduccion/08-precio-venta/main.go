package main

import "fmt"

//Función que calcula el igv
func calcularIgv(vv, igv float64) float64 {
	return igv * vv
}

//Funcion que calcula percio de venta
func precioVenta(vv, igv float64) float64 {
	return vv + igv
}

//Función principal
func main() {

	//Variables
	var vv, igv, pv float64

	//Entrada  de datos
	fmt.Print("Ingrese Valor de Venta: ")
	fmt.Scanln(&vv)

	//Calcular igv
	igv = calcularIgv(vv, 0.18)

	//Calcular precio de venta
	pv = precioVenta(vv, igv)

	//Salida de datos
	fmt.Println("==========================")
	fmt.Println("Valor de Venta:", vv)
	fmt.Println("IGV:", igv)
	fmt.Println("Precio de Venta:", pv)
	fmt.Println("==========================")
}
