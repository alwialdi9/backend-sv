package main

import (
	backendsv "github.com/alwialdi9/backend-sv"
	_ "github.com/alwialdi9/backend-sv/docs"
)

// @title backend SV
// @version 1.0
// @description This is a doc API for backend SV
// @termsOfService http://swagger.io/terms/
// @contact.name Alwi Aldiansyach
// @contact.email alwi.aldisyach@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath localhost:3000/api
func main() {
	backendsv.App()
}
