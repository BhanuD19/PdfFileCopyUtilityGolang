package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"seehuhn.de/go/pdf"
	"strconv"
	"sync"
)

//func main() {
//	wg := new(sync.WaitGroup)
//	srcFolder := "/Users/vaibhav/IdeaProjects/FileUploaderDemo NativeImageBuilds copy-2/bills-pdf/"
//	srcFile := "/Users/vaibhav/IdeaProjects/file-uploader-demo-go/temp/2519~Customer 123~AWSTEMP1000~2023-05-05~10000.pdf"
//	for i := 1000; i < 2000; i++ {
//		go func(srcFile, srcFolder string, i int, wg *sync.WaitGroup) {
//			var src, err = os.Open(srcFile)
//			if err != nil {
//				log.Fatal("Couldn't open file to copy")
//			}
//			var dest, err2 = os.Create(srcFolder + "2519~Customer 123~AWS-NEW" + strconv.Itoa(i) + "~2023-05-05~" + strconv.Itoa(i*10) + ".pdf")
//			if err2 != nil {
//				log.Fatal("Couldn't copy file")
//			}
//			size, err := io.Copy(dest, src)
//			if err != nil {
//				log.Fatal("Couldn't copy file")
//			}
//			fmt.Println("\n Copied bytes: ", size)
//			err = dest.Sync()
//			if err != nil {
//				log.Fatal("Couldn't copy file because", err)
//			}
//			err = dest.Close()
//			err = src.Close()
//			if err != nil {
//				log.Fatal("Couldn't close file")
//			}
//			wg.Done()
//		}(srcFile, srcFolder, i, wg)
//		wg.Add(1)
//	}
//	wg.Wait()
//}

func main() {
	wg := new(sync.WaitGroup)
	srcFolder := "/Users/vaibhav/IdeaProjects/FileUploaderDemo NativeImageBuilds copy/ledger-pdf/"
	srcFile := "/Users/vaibhav/IdeaProjects/file-uploader-demo-go/temp/2519~Customer 123.pdf"

	for i := 1000; i < 2000; i++ {
		go func(srcFile, srcFolder string, i int, wg *sync.WaitGroup) {
			var src, err = os.Open(srcFile)
			if err != nil {
				log.Fatal("Couldn't open file to copy")
			}
			var dest, err2 = os.Create(srcFolder + "2519~Customer 123321" + strconv.Itoa(i) + ".pdf")
			if err2 != nil {
				log.Fatal("Couldn't copy file")
			}
			size, err := io.Copy(dest, src)
			if err != nil {
				log.Fatal("Couldn't copy file")
			}
			fmt.Println("\n Copied bytes: ", size)
			err = dest.Sync()
			if err != nil {
				log.Fatal("Couldn't copy file because", err)
			}
			err = dest.Close()
			err = src.Close()
			if err != nil {
				log.Fatal("Couldn't close file")
			}
			wg.Done()
		}(srcFile, srcFolder, i, wg)
		wg.Add(1)
	}
	wg.Wait()
}

//func main() {
//	wg := new(sync.WaitGroup)
//	srcFolder := "/Users/vaibhav/IdeaProjects/file-uploader-demo-go/temp/bills-pdf/"
//	srcFile := "/Users/vaibhav/IdeaProjects/file-uploader-demo-go/temp/bills-pdf/2519~Customer 123~AWSTEMP1122~2023-05-05~10000.pdf"
//	buf := make([]byte, 204800)
//	src, err := os.Open(srcFile)
//	if err != nil {
//		log.Fatal("Couldn't open file to copy")
//	}
//	defer src.Close()
//	for i := 1000; i < 2000; i++ {
//		fName := "2519~Customer 123~AWSTEMP" + strconv.Itoa(i) + "~2023-05-05~10000.pdf"
//		destString := srcFolder + fName
//		dest, err := os.Create(destString)
//		if err != nil {
//			log.Fatal("Couldn't copy file")
//		}
//		defer dest.Close()
//		go copyFile(buf, src, dest, wg)
//		wg.Add(1)
//	}
//	wg.Wait()
//}

//func main() {
//	wg := new(sync.WaitGroup)
//	srcFolder := "/Users/vaibhav/IdeaProjects/file-uploader-demo-go/ledger-pdf/"
//	srcFile := "/Users/vaibhav/IdeaProjects/file-uploader-demo-go/temp/ledger-pdf/2519~Customer 123.pdf"
//	buf := make([]byte, 204800)
//	src, err := os.Open(srcFile)
//	if err != nil {
//		log.Fatal("Couldn't open file to copy")
//	}
//	defer src.Close()
//	for i := 1000; i < 2000; i++ {
//		fName := "2519~Customer " + strconv.Itoa(i) + ".pdf"
//		destString := srcFolder + fName
//		dest, err := os.Create(destString)
//		if err != nil {
//			log.Fatal("Couldn't copy file")
//		}
//		defer dest.Close()
//		go copyFile(buf, src, dest, wg)
//		wg.Add(1)
//	}
//	wg.Wait()
//}

//func main() {
//	wg := new(sync.WaitGroup)
//	dstFolder := "/Users/vaibhav/IdeaProjects/tempFileCopyUtility/"
//	srcFile := "/Users/vaibhav/IdeaProjects/file-uploader-demo-go/temp/2519~Customer 123.pdf"
//
//	src, err := pdf.Open(srcFile, nil)
//	if err != nil {
//		log.Fatal("Couldn't open file to copy")
//	}
//	defer src.Close()
//	for i := 1; i < 2; i++ {
//		fName := "481~Customer " + strconv.Itoa(i) + ".pdf"
//		destString := dstFolder + fName
//		dest, err := os.Create(destString)
//		if err != nil {
//			log.Fatal("Couldn't copy file")
//		}
//		defer dest.Close()
//		//go copyFile(buf, src, dest, wg)
//		go makePdf(src, dest, wg)
//		wg.Add(1)
//	}
//	wg.Wait()
//}

type walker struct {
	trans map[pdf.Reference]pdf.Reference
	r     *pdf.Reader
	w     *pdf.Writer
}

func makePdf(r *pdf.Reader, out *os.File, wg *sync.WaitGroup) {
	w, err := pdf.NewWriter(out, &pdf.WriterOptions{
		Version: r.Version,
	})
	if err != nil {
		log.Fatal(err)
	}

	catalog := r.Catalog

	trans := &walker{
		trans: map[pdf.Reference]pdf.Reference{},
		r:     r,
		w:     w,
	}
	catDict := pdf.AsDict(catalog)
	newCatDict := pdf.Dict{}
	for key, val := range catDict {
		obj, err := trans.Transfer(val)
		if err != nil {
			log.Fatal(err)
		}
		newCatDict[key] = obj
	}
	err = r.DecodeDict(catalog, newCatDict)
	if err != nil {
		log.Fatal(err)
	}

	trans.w.SetInfo(r.Info)

	trans.w.Catalog = catalog

	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}
	wg.Done()
}

func copyFile(buf []byte, src *os.File, dest *os.File, wg *sync.WaitGroup) {
	for {
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		if _, err := dest.Write(buf[:n]); err != nil {
			log.Fatal(err)
		}
	}
	wg.Done()
}

func (w *walker) Transfer(obj pdf.Object) (pdf.Object, error) {
	switch x := obj.(type) {
	case pdf.Dict:
		res := pdf.Dict{}
		for key, val := range x {
			repl, err := w.Transfer(val)
			if err != nil {
				return nil, err
			}
			res[key] = repl
		}
		return res, nil
	case pdf.Array:
		var res pdf.Array
		for _, val := range x {
			repl, err := w.Transfer(val)
			if err != nil {
				return nil, err
			}
			res = append(res, repl)
		}
		return res, nil
	case *pdf.Stream:
		res := &pdf.Stream{
			Dict: make(pdf.Dict),
			R:    x.R,
		}
		for key, val := range x.Dict {
			repl, err := w.Transfer(val)
			if err != nil {
				return nil, err
			}
			res.Dict[key] = repl
		}
		return res, nil
	case pdf.Reference:
		other, ok := w.trans[x]
		if ok {
			return other, nil
		}
		other = w.w.Alloc()
		w.trans[x] = other

		val, err := pdf.Resolve(w.r, x)
		if err != nil {
			return nil, err
		}
		trans, err := w.Transfer(val)
		if err != nil {
			return nil, err
		}
		err = w.w.Put(other, trans)
		if err != nil {
			return nil, err
		}
		return other, nil
	}
	return obj, nil
}
