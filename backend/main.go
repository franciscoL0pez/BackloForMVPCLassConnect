// print hello word

package main

import (

	"github.com/franciscoL0pez/MVPCLassConnect/backend/dbConfig/"
)

func main() {
	
	dbConfig.ConnectDB()
	defer dbConfig.CloseDB()

}