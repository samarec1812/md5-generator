package md5_generator

import (
	"fmt"
	"github.com/samarec1812/md5-generator/internal/app/algo"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

var (
	Salt               string = "sflpr9fhi2" // произвольная соль
	SaveInTxt          bool   = false        // флаг сохранения в файле
	CheckFile          bool   = false
	DataKey            string = ""
	ValidResult        bool   = true
	CalculateValidHash        = false
)

type HashStruct struct {
	Hash               string
	Files              bool
	CheckFileValid     bool
	ValidResult        bool
	HashValid          string
	CalculateValidHash bool
}

func (a Services) Run() error {
	port := os.Getenv("PORT")
	// port := "8081"
	fs := http.FileServer(http.Dir("asserts"))
	http.Handle("/", fs)
	data := ""
	dataValid := ""
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(http.StatusOK)
	//
	//})

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		fileContents, err := ioutil.ReadFile("asserts/index.html")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if r.Method == "GET" {
			SaveInTxt = false
			CheckFile = false
			w.WriteHeader(http.StatusOK)
			w.Write(fileContents)
		} else if r.Method == "POST" {
			if r.FormValue("send-text") != "" {
				dataKey := ""
				text := r.FormValue("send-text")
				if r.FormValue("hash-empty-string") == "on" {
					text = ""
				}
				if r.FormValue("check-password") == "on" {
					fmt.Println(r.FormValue("inputPassword"), reflect.TypeOf(r.FormValue("inputPassword")))
					password := r.FormValue("inputPassword")
					dataKey = algo.CalcMD5([]byte(password + Salt))
				}
				dataALl := append([]byte(dataKey), []byte(text)...)
				data = algo.CalcMD5(dataALl)
				//data = algo.CalcMD5([]byte(text))
				if r.FormValue("check-save-file-text") == "on" {
					_, err = os.Create("hash.txt")
					if err != nil {
						log.Println(err)
						w.WriteHeader(http.StatusNotFound)
						return
					}
					err = ioutil.WriteFile("hash.txt", []byte(data), 0)
					if err != nil {
						log.Println(err)
						w.WriteHeader(http.StatusNotFound)
						return
					}
					SaveInTxt = true
					//	fmt.Println(SaveInTxt)
				}

				http.Redirect(w, r, "hash-result", http.StatusSeeOther)

			} else {
				dataKey := ""
				if r.FormValue("check-password") == "on" {
					fmt.Println(r.FormValue("inputPassword"), reflect.TypeOf(r.FormValue("inputPassword")))
					password := r.FormValue("inputPassword")
					dataKey = algo.CalcMD5([]byte(password + Salt))
				}
				file, _, err := r.FormFile("file")
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusNotFound)
					return
				}
				defer file.Close()
				filesBytes, err := ioutil.ReadAll(file)
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusNotFound)
					return
				}
				DataKey = dataKey
				dataALl := append([]byte(dataKey), filesBytes...)
				data = algo.CalcMD5(dataALl)

				if r.FormValue("check-save-file") == "on" {
					fmt.Println("ЗДЕСЬ ЗАПИСЬ В ФАЙЛ")

					_, err = os.Create("hash.txt")
					if err != nil {
						log.Println(err)
						w.WriteHeader(http.StatusNotFound)
						return
					}
					err = ioutil.WriteFile("hash.txt", []byte(data), 0)
					if err != nil {
						log.Println(err)
						w.WriteHeader(http.StatusNotFound)
						return
					}
					SaveInTxt = true
				}
				if r.FormValue("check-validation") == "on" {
					CheckFile = true
				}

				http.Redirect(w, r, "hash-result", http.StatusSeeOther)
				// fmt.Fprintln(w, r.Form)
			}
		}
	})

	http.HandleFunc("/hash-result", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			if r.FormValue("download") == "on" {
				w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote("hash.txt"))
				w.Header().Set("Content-Type", "application/octet-stream")
				http.ServeFile(w, r, "hash.txt")
			} else {
				tmpl, err := template.ParseFiles("asserts/hash-result.html")
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusNotFound)
					return
				}
				fmt.Println(SaveInTxt)
				dataTemplate := HashStruct{Hash: data, Files: SaveInTxt, CheckFileValid: CheckFile, CalculateValidHash: false}
				err = tmpl.Execute(w, dataTemplate)
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusNotFound)
					return
				}
			}
		} else if r.Method == "POST" {
			if r.FormValue("check-valid") == "on" {
				file, _, err := r.FormFile("file-valid")
				if err != nil {
					log.Println(err)
					fmt.Println("TUT1")
					w.WriteHeader(http.StatusNotFound)
					return
				}
				defer file.Close()
				filesBytes, err := ioutil.ReadAll(file)
				if err != nil {
					fmt.Println("NEN1")
					log.Println(err)
					w.WriteHeader(http.StatusNotFound)
					return
				}

				dataALl := append([]byte(DataKey), filesBytes...)
				dataValid = algo.CalcMD5(dataALl)

				if dataValid != data {
					fmt.Println(dataValid, data)
					ValidResult = false
				}
				dataTemplate := HashStruct{Hash: data, Files: SaveInTxt, CheckFileValid: CheckFile, ValidResult: ValidResult, HashValid: dataValid, CalculateValidHash: true}
				tmpl, err := template.ParseFiles("asserts/hash-result.html")
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusNotFound)
					return
				}
				err = tmpl.Execute(w, dataTemplate)
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusNotFound)
					return
				}
			} else {
				http.Redirect(w, r, "/index", http.StatusSeeOther)
			}
			// defer os.Remove("asserts/hash.txt")

		}
	})
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
