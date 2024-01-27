package businesslogic

// DefaultChannels provides the default set of channels.
func DefaultChannels() map[string]string {
	return map[string]string{
		"NTV":               "XEJM4Hcgd3M",
		"Habertürk":         "SqHIO2zhxbA",
		"Haber Global":      "UVPejgEw21c",
		"TRT Haber":         "Rc5qrxlJZzc",
		"TV 100":            "sd94keSra6A",
		"Halk TV":           "_FUlFVJo77I",
		"24 TV":             "V5mBTSql74Q",
		"TGRT Haber":        "8YPC2IV7ve0",
		"KRT TV":            "3QDiWPZ2D_k",
		"TELE 1":            "mRK3wXGdsLk",
		"Bloomberg HT":      "hHSmBJk6w0c",
		"Ulusal Kanal":      "SdCJquYL-CQ",
		"Artı TV":           "xpoetRCJKqY",
		"TVNET":             "SR396EBvGUk",
		"Ülke TV":           "dAknS8uSPW8",
		"Flash Haber TV":    "oGHfM6AC7QU",
		"Bengü Türk":        "Lx_fVB6FQ5Y",
		"Kanal D":           "ubWBmjt4x7U",
		"Show TV":           "eVVlKAUlSNQ",
		"Fox TV":            "6W1ePcaAUFY",
		"360 TV":            "bOlsH34--PU",
		"TV5":               "-zNkYBJ0wKI",
		"Ekotürk TV":        "zJYTpkv7aAQ",
		"Cadde TV":          "BSM8MIH7yAE",
		"beIN Sports Haber": "MZ5D6NK5Qqo",
	}
}

// PrepareChannels prepares the channel map based on input or defaults.
func PrepareChannels(channelNames, channelSources []string) map[string]string {
	if len(channelNames) > 0 && len(channelSources) > 0 {
		channelsMap := make(map[string]string)
		for i, name := range channelNames {
			if i < len(channelSources) {
				channelsMap[name] = channelSources[i]
			}
		}
		return channelsMap
	}
	return DefaultChannels()
}
