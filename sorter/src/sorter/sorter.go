package main

import (
	"flag"
	"time"
)
import "fmt"
import "os"
import "bufio"
import "strconv"
import "io"
import "src/algorithm/bubblesort"
import "src/algorithm/qsort"

var infile *string = flag.String("i", "infile", "file contains values for sorting")
var outfile *string = flag.String("o", "outfile", "file to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort Algorithm")

func readValues(infile string) (values [] int, err error) {
	file, err := os.Open(infile)
	if err != nil{
		fmt.Println("Failed to open input file ", infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)

	values = make([] int, 0)

	for{
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil{
			if err1 != io.EOF{
				err = err1
			}
			break
		}
		
		if isPrefix{
			fmt.Println("A too long line, seems unexpected")
			return
		}

		str := string(line)

		value, err1 := strconv.Atoi(str)

		if err1 != nil{
			err = err1
			return
		}
		
		values = append(values, value)
	}
	return
}


func writeFile(outfile string, values [] int) error {
	file, err := os.Create(outfile)
	if err != nil{
		fmt.Println("create file error: ", err)
		return err
	}
	defer file.Close()

	for _, value := range values{
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}



func main() {
        flag.Parse()

        if infile != nil {
                fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algorithm", *algorithm)
        }

		values, err := readValues(*infile)
		if err == nil{
			t1 := time.Now()
			switch *algorithm {
			case "qsort":
				qsort.QuickSort(values)
			case "bubblesort":
				bubblesort.BubbleSort(values)
			default:
				fmt.Println("Sorting algorithm", * algorithm, "is either unknown or unsupported")
			}
			t2 := time.Now()
			fmt.Println("the sorting process costs ", t2.Sub(t1), "to complete")
		}else {
			fmt.Println(err)
			return
		}
		if outfile == nil{
			fmt.Println("not outfile path, display result to screen")
			for i := range values{
				fmt.Println(i)
			}
		}
		err = writeFile(*outfile, values)
		if err != nil{
			fmt.Println("write error: ", err)
		}

}