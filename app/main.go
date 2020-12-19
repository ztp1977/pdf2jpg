package main

// export CGO_CFLAGS_ALLOW='-Xpreprocessor'
import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, file.Name())
	}

	return paths
}

const Format = "jpg"

func convert(in string, pdf string, out string) error {
	readFile := filepath.Join(in, pdf)
	if !strings.HasSuffix(readFile, "pdf") {
		return nil
	}
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	err := mw.SetResolution(150, 150)
	if err != nil {
		log.Fatal("failed at SetResolution", err)
	}

	err = mw.ReadImage(readFile)
	if err != nil {
		log.Fatal("failed at ReadImage", err)
	}

	n := mw.GetNumberImages()
	log.Println("number image: ", n)

	err = mw.SetImageFormat(Format)
	if err != nil {
		log.Fatal("failed at SetImageFormat")
	}

	for i := 0; i < int(n); i++ {
		if ret := mw.SetIteratorIndex(i); !ret {
			break
		}

		writeFile := filepath.Join(out, fmt.Sprintf("%s_%d.%s", strings.Trim(pdf, ".pdf"), i, Format))
		err = mw.WriteImage(writeFile)
		if err != nil {
			log.Fatal("failed at WriteImage, " + writeFile)
		}
	}
	return err
}

func main() {
	var inDir = flag.String("in", "data", "input pdf dir")
	var outDir = flag.String("out", "out", "output image dir")
	flag.Parse()

	imagick.Initialize()
	defer imagick.Terminate()

	pdfs := dirwalk(*inDir)
	for _, pdf := range pdfs {
		err := convert(*inDir, pdf, *outDir)
		if err != nil {
			panic(err)
		}
	}
}
