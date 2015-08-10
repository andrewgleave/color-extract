package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"log"
	"os"
	"sync"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/soniakeys/quant"
	"github.com/soniakeys/quant/mean"
	"github.com/soniakeys/quant/median"
)

func main() {
	mode := flag.String("mode", "median", "'median' or 'mean' (default 'median')")
	size := flag.Int("size", 16, "Palette size (default 16)")
	flag.Parse()

	paths := flag.Args()
	if len(paths) == 0 {
		log.Fatal("No images specified")
	}

	var wg sync.WaitGroup
	wg.Add(len(paths))

	for _, path := range paths {
		go func(path string) {
			defer wg.Done()

			var buff bytes.Buffer
			buff.WriteString(fmt.Sprintf("\n%s\n", path))

			reader, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer reader.Close()

			img, _, err := image.Decode(reader)
			if err != nil {
				log.Fatal(err)
			}

			var q quant.Quantizer
			if *mode == "median" {
				q = median.Quantizer(*size)
			} else {
				q = mean.Quantizer(*size)
			}

			pal := q.Palette(img)
			for _, c := range pal.ColorPalette() {
				r, g, b, _ := c.RGBA()
				buff.WriteString(fmt.Sprintf("#%02x%02x%02x\n", uint8(r>>8), uint8(g>>8), uint8(b>>8)))
			}

			fmt.Print(buff.String())
		}(path)
	}

	wg.Wait()
}
