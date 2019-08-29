package main

import (
	"net/http"

	"github.com/byuoitav/common"
	"github.com/byuoitav/common/v2/auth"
	"github.com/byuoitav/epson-projector-ms/handlers"
)

func main() {
	port := ":8069"
	router := common.NewRouter()

	read := router.Group("", auth.AuthorizeRequest("read-state", "room", auth.LookupResourceFromAddress))
	read.GET("/:address/power", handlers.GetPower)
	read.GET("/:address/blanked", handlers.GetBlanked)
	read.GET("/:address/input", handlers.GetInput)
	read.GET("/:address/muted", handlers.GetMuted)
	read.GET("/:address/volume", handlers.GetVolume)
	//	read.GET("/:address/hardware", handlers.GetHardwareInfo)
	//	read.GET("/:address/activesignal", handlers.GetActiveSignal)

	write := router.Group("", auth.AuthorizeRequest("write-state", "room", auth.LookupResourceFromAddress))
	write.GET("/:address/power/:state", handlers.SetPower)     // 'on' or 'standby'
	write.GET("/:address/blanked/:state", handlers.SetBlanked) // 'blank' or 'unblank'
	write.GET("/:address/input/:port", handlers.SetInput)      // one of the epson ports
	write.GET("/:address/muted/:state", handlers.SetMuted)     // 'mute' or 'unmute'
	write.GET("/:address/volume/:level", handlers.SetVolume)   // 1-100

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
