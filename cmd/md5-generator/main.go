package main

import md5_generator "github.com/samarec1812/md5-generator/internal/app/md5-generator"

// точка входа. запуск приложения


func main() {
	 var i md5_generator.Service = md5_generator.Services{PORT: ":8081"}
	// algo.Hashing()
	 i.Run()

}