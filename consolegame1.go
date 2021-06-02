package main

import (
	"encoding/csv"
	"os"
	"fmt"
)

const rooms = "rooms.csv"

func findText(n string,a [][]string) string {
//for a given 2d array where [i][j] = [csv line number][csv line contents]
//match the input string to the first column and output the second
//this formats the given array and returns a string which is like description+exits
	description := ""
	for i := 0; i < len(a); i++ {
		if a[i][0] == n {
			description = a[i][1]
			for j := 2; j < len(a[i]); j++ {
				description = description + a[i][j]
			}
			return description
		}
	}
	return "-"
}
func openTexts(fpath string) [][]string{
//create a 2d array with the texts arranged like string[i][j] with i being the csv line and j being the comma deliminated contents of that line 
	file, err := os.Open(fpath)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, _ := reader.ReadAll()
	return records
}
func oFP(name string) {
//ofp=openfindprint from the rooms list, the description
//	open the file for reading, find the wanted text by searching the first collumn of the file  for name and print the second column 
	texts := openTexts(rooms)
	fmt.Println(findText(name, texts))
}
func main(){
	var currentroom = "introscrawl"
	var input string
	oFP(currentroom)
	for {
		fmt.Scanf("%s", &input)
		if input == "l"{
			oFP(currentroom)
		}
		input = ""
	}
}

