// Stefan Nilsson 2013-02-27
// Farhan Syed 2022-04-05

// Original runtime: 17.409s
// Improved runtime: 4.096s

// This program creates pictures of Julia sets (en.wikipedia.org/wiki/Julia_set).
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"runtime"
	"strconv"
	"sync"
)

type ComplexFunc func(complex128) complex128

var Funcs []ComplexFunc = []ComplexFunc{
	func(z complex128) complex128 { return z*z - 0.61803398875 },
	func(z complex128) complex128 { return z*z + complex(0, 1) },
	func(z complex128) complex128 { return z*z + complex(-0.835, -0.2321) },
	func(z complex128) complex128 { return z*z + complex(0.45, 0.1428) },
	func(z complex128) complex128 { return z*z*z + 0.400 },
	func(z complex128) complex128 { return cmplx.Exp(z*z*z) - 0.621 },
	func(z complex128) complex128 { return (z*z+z)/cmplx.Log(z) + complex(0.268, 0.060) },
	func(z complex128) complex128 { return cmplx.Sqrt(cmplx.Sinh(z*z)) + complex(0.065, 0.122) },
}

func init() {
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu) // Use all available CPUs.
}

func main() {
	var wg sync.WaitGroup
	wg.Add(len(Funcs))
	for n, fn := range Funcs {
		//A new go routine is created for each image
		go CreatePng("picture-"+strconv.Itoa(n)+".png", fn, 1024, &wg)
	}
	wg.Wait() // waits for all images to be created
}

// CreatePng creates a PNG picture file with a Julia image of size n x n.
func CreatePng(filename string, f ComplexFunc, n int, wg *sync.WaitGroup) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()
	err = png.Encode(file, Julia(f, n))
	if err != nil {
		log.Fatal(err)
	}
	wg.Done()
	return
}

// Julia returns an image of size n x n of the Julia set for f.
func Julia(f ComplexFunc, n int) image.Image {
	bounds := image.Rect(-n/2, -n/2, n/2, n/2)

	// Creating a waitGroup and adding n elements (i.e. number of pixels in one row)
	var wg2 sync.WaitGroup
	wg2.Add(n)

	img := image.NewRGBA(bounds)
	s := float64(n / 4)

	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		// A go routine that creates all the pixels below the top pixel (that is in the top row)
		go func(i int, s float64) {
			for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
				n := Iterate(f, complex(float64(i)/s, float64(j)/s), 256)
				r := uint8(0)
				g := uint8(n % 32 * 8)
				b := uint8(n % 32 * 8)
				img.Set(i, j, color.RGBA{r, g, b, 255})
			}
			wg2.Done()
		}(i, s)
	}
	wg2.Wait() //waits for all pixels to be created
	return img
}

// Iterate sets z_0 = z, and repeatedly computes z_n = f(z_{n-1}), n â‰¥ 1,
// until |z_n| > 2  or n = max and returns this n.
func Iterate(f ComplexFunc, z complex128, max int) (n int) {
	for ; n < max; n++ {
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			break
		}
		z = f(z)
	}
	return
}
