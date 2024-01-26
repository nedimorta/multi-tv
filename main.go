package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(context *gin.Context) {
		rootURL := "http://localhost:8080/"
		var channelsMap map[string]string
		defaultChannelsMap := map[string]string{
			"NTV":          "XEJM4Hcgd3M",
			"Habertürk":    "SqHIO2zhxbA",
			"Haber Global": "UVPejgEw21c",
			"TRT Haber":    "Rc5qrxlJZzc",
		}

		// Retrieve query parameters as slices
		channelNames := context.QueryArray("channelNames")
		channelSources := context.QueryArray("channelSources")

		// Check if the slices are empty to determine if the parameters were provided
		channelNamesExist := len(channelNames) > 0
		channelSourcesExist := len(channelSources) > 0

		if channelNamesExist && channelSourcesExist {
			channelsMap = make(map[string]string)
			for i, name := range channelNames {
				if i < len(channelSources) {
					channelsMap[name] = channelSources[i]
				}
			}
		} else {
			channelsMap = defaultChannelsMap
		}

		channelCount, _ := strconv.Atoi(context.DefaultQuery("channelCount", "9"))
		var rowClassCSS string
		switch {
		case channelCount <= 4:
			rowClassCSS = "row row-cols-1 row-cols-sm-2 justify-content-center align-items-center m-0"
		case channelCount <= 9:
			rowClassCSS = "row row-cols-1 row-cols-sm-2 row-cols-md-3 justify-content-center align-items-center m-0"
		case channelCount <= 16:
			rowClassCSS = "row row-cols-1 row-cols-sm-2 row-cols-xl-4 justify-content-center align-items-center m-0"
		default:
			rowClassCSS = "row row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-xl-4 row-cols-xxl-5 justify-content-center align-items-center m-0"
		}

		autoplayEnabled, _ := strconv.ParseBool(context.DefaultQuery("autoplayEnabled", "true"))

		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Root":     rootURL,
			"Channels": channelsMap,
			"RowClass": rowClassCSS,
			"Autoplay": autoplayEnabled,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
