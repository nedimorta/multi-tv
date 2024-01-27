package handlers

import (
	"multi-tv/businesslogic"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleMainPage processes the main page request.
func HandleMainPage(context *gin.Context) {
	channelNames := context.QueryArray("channelNames")
	channelSources := context.QueryArray("channelSources")

	channelsMap := businesslogic.PrepareChannels(channelNames, channelSources)
	_, autoplayParam, channelCount := businesslogic.GetUserPreferences(context)
	rowClassCSS := businesslogic.DetermineRowClassCSS(channelCount)

	for name, url := range channelsMap {
		channelsMap[name] = "https://www.youtube-nocookie.com/embed/" + url + "?autoplay=" + autoplayParam + "&mute=1"
	}

	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"RootURL":  "http://localhost:8080/",
		"Channels": channelsMap,
		"RowClass": rowClassCSS,
	})
}
