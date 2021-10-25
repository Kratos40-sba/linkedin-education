package main

import (
	"encoding/csv"
	"os"
)

func main()  {

}

type Person struct {
	Name string
	Age int
}

func WriteToCSV(fileName string , persons []Person) error  {
	file , err := os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err != nil {
		return err
	}
	defer file.Close()
	row := []string{"Name","Age"}
	csvWriter := csv.NewWriter(file)
	// data in the buffer will flushed
	defer csvWriter.Flush()
	if err := csvWriter.Write(row);err != nil {
		return err
	}
	for _ , person := range persons {
		row[0] = person.Name
		row[1] = string(rune(person.Age))
		if err := csvWriter.Write(row);err != nil {
			return err
		}
	}

	return csvWriter.Error()
}