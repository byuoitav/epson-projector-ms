package main

import (
	"github.com/byuoitav/common"
	"github.com/byuoitav/common/v2/auth"
)

func main() {
	port := ":8069"
	router := common.NewRouter()

	write := router.Group("", auth.AuthorizeRequest("write-state", "room", auth.LookupResourceFromAddress))

	read := router.Group("", auth.AuthorizeRequest("read-state", "room", auth.LookupResourceFromAddress))

}
