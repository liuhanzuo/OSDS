package main

import (
	"os"
	"fmt"
	"log"
	"image/jpeg"
	"time"
	"github.com/nfnt/resize"
)

func main(){
	start:=time.Now()
	for i:=0;i<10000;i++ {
		filename := fmt.Sprintf("./gitclone/tiny-imagenet-200/test/images/test_%d.JPEG",i);
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Load picture failed", filename)
			log.Fatal(err)
		}
		img, err := jpeg.Decode(file)
		m := resize.Resize(128, 128, img, resize.Lanczos3)
		if err != nil {
			fmt.Println("Decode picture failed", filename)
			log.Fatal(err)
		}
		file.Close()
		filename = fmt.Sprintf("./gitclone/tiny-imagenet-200/test/images/resize_demo_resized_%d.JPEG",i);
		out, err := os.Create(filename)
		if err != nil {
			fmt.Println("Write picture failed", filename)
			log.Fatal(err)
		}
		defer out.Close()

		jpeg.Encode(out, m, nil)
	}
	fmt.Println("time used baseline %v",time.Since(start))
}
