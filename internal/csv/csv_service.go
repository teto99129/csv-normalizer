package csv

import (
	"encoding/csv"
	"io"
	"log"
	"mime/multipart"
)

type CSVService struct {}

func (s *CSVService) CSVToJson(file multipart.File) ([]map[string]string) {
  r := csv.NewReader(file)

  headerMap := map[int]string{}
  jsonMap := map[string]string{}
  array := []map[string]string{}

  header, err := r.Read()
  if err == io.EOF {
    log.Fatal("1行もないCSV")
  }
  if err != nil {
    log.Fatal("何らかのエラー")
  }

  for i, v := range header {
    headerMap[i] = v
  }
  
  for {
    record, err := r.Read()
    if err == io.EOF {
    	break
    }
    if err != nil {
    	log.Fatal(err)
    }

    for ri, c := range record {
      jsonMap[headerMap[ri]] = c
    }
    
    array = append(array, jsonMap)
  }

  return array

}
