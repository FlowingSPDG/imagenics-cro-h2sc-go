package h2sc

var (
	OUTPUT_FORMAT_1080i5994 = []byte("00000")
	OUTPUT_FORMAT_480i      = []byte("+0001")
	OUTPUT_FORMAT_720p5994  = []byte("+0004")
	OUTPUT_FORMAT_1080p5994 = []byte("+0005")
	OUTPUT_FORMAT_1080p60   = []byte("+0006")
	OUTPUT_FORMAT_1080p30   = []byte("+0008")
	// OUTPUT_FORMAT_1080i5994 = []byte("+0016") // Same as 000000
	// OUTPUT_FORMAT_480i = []byte("+0017") // same as +0001
)
