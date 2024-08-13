package uiavatars

import (
	"bytes"
	"math/rand"
	"strings"
	"time"
	"unicode"

	svg "github.com/ajstarks/svgo"
)

// GenerateImage creates an SVG image and returns its content as a base64 string.
func GenerateImage(name string) []byte {
	width := 64
	height := 64

	// Generate SVG content as a string
	svgContent := generateSVG(name, width, height)
	svgData := []byte(svgContent)
	return svgData
}

// GetRandomColor returns a random color from a predefined list.
func getRandomColor() (string, string) {
	colors := []struct {
		b string // background color
		t string // text color
	}{
		{"5e35b1", "FFFFFF"},
		{"512da8", "FFFFFF"},
		{"4527a0", "FFFFFF"},
		{"311b92", "FFFFFF"},
		{"8e24aa", "FFFFFF"},
		{"7b1fa2", "FFFFFF"},
		{"6a1b9a", "FFFFFF"},
		{"4a148c", "FFFFFF"},
		{"3949ab", "FFFFFF"},
		{"303f9f", "FFFFFF"},
		{"283593", "FFFFFF"},
		{"1a237e", "FFFFFF"},
		{"1e88e5", "FFFFFF"},
		{"1976d2", "FFFFFF"},
		{"1565c0", "FFFFFF"},
		{"0d47a1", "FFFFFF"},
		{"039be5", "FFFFFF"},
		{"0288d1", "FFFFFF"},
		{"0277bd", "FFFFFF"},
		{"01579b", "FFFFFF"},
		{"00acc1", "FFFFFF"},
		{"0097a7", "FFFFFF"},
		{"00838f", "FFFFFF"},
		{"006064", "FFFFFF"},
		{"00897b", "FFFFFF"},
		{"00796b", "FFFFFF"},
		{"00695c", "FFFFFF"},
		{"004d40", "FFFFFF"},
		{"43a047", "FFFFFF"},
		{"388e3c", "FFFFFF"},
		{"2e7d32", "FFFFFF"},
		{"1b5e20", "FFFFFF"},
		{"7cb342", "FFFFFF"},
		{"689f38", "FFFFFF"},
		{"558b2f", "FFFFFF"},
		{"33691e", "FFFFFF"},
		{"c0ca33", "FFFFFF"},
		{"afb42b", "FFFFFF"},
		{"9e9d24", "FFFFFF"},
		{"827717", "FFFFFF"},
		{"fdd835", "FFFFFF"},
		{"fbc02d", "FFFFFF"},
		{"f9a825", "FFFFFF"},
		{"f57f17", "FFFFFF"},
		{"ffb300", "FFFFFF"},
		{"ffa000", "FFFFFF"},
		{"ff8f00", "FFFFFF"},
		{"ff6f00", "FFFFFF"},
		{"fb8c00", "FFFFFF"},
		{"f57c00", "FFFFFF"},
		{"ef6c00", "FFFFFF"},
		{"e65100", "FFFFFF"},
		{"f4511e", "FFFFFF"},
		{"e64a19", "FFFFFF"},
		{"d84315", "FFFFFF"},
		{"bf360c", "FFFFFF"},
		{"6d4c41", "FFFFFF"},
		{"5d4037", "FFFFFF"},
		{"4e342e", "FFFFFF"},
		{"3e2723", "FFFFFF"},
		{"546e7a", "FFFFFF"},
		{"455a64", "FFFFFF"},
		{"37474f", "FFFFFF"},
		{"263238", "FFFFFF"},
		{"F44336", "FFFFFF"},
		{"E53935", "FFFFFF"},
		{"D32F2F", "FFFFFF"},
		{"C62828", "FFFFFF"},
		{"B71C1C", "FFFFFF"},
		{"FFEBEE", "000000"},
		{"FFCDD2", "000000"},
		{"EF9A9A", "000000"},
		{"E57373", "000000"},
		{"EF5350", "000000"},
		{"FF8A80", "000000"},
		{"FF5252", "000000"},
		{"FF1744", "000000"},
		{"D50000", "000000"},
		{"FCE4EC", "000000"},
		{"F8BBD0", "000000"},
		{"F48FB1", "000000"},
		{"F06292", "000000"},
		{"EC407A", "000000"},
		{"FF80AB", "000000"},
		{"FF4081", "000000"},
		{"F50057", "000000"},
		{"C51162", "000000"},
		{"D81B60", "FFFFFF"},
		{"C2185B", "FFFFFF"},
		{"AD1457", "FFFFFF"},
		{"880E4F", "FFFFFF"},
		{"9c27b0", "000000"},
		{"f3e5f5", "000000"},
		{"e1bee7", "000000"},
		{"ce93d8", "000000"},
		{"ba68c8", "000000"},
		{"ab47bc", "000000"},
		{"ea80fc", "000000"},
		{"e040fb", "000000"},
		{"d500f9", "000000"},
		{"aa00ff", "000000"},
		{"673ab7", "000000"},
		{"ede7f6", "000000"},
		{"d1c4e9", "000000"},
		{"b39ddb", "000000"},
		{"9575cd", "000000"},
		{"7e57c2", "000000"},
		{"b388ff", "000000"},
		{"7c4dff", "000000"},
		{"651fff", "000000"},
		{"6200ea", "000000"},
		{"3f51b5", "000000"},
		{"e8eaf6", "000000"},
		{"c5cae9", "000000"},
		{"9fa8da", "000000"},
		{"7986cb", "000000"},
		{"5c6bc0", "000000"},
		{"8c9eff", "000000"},
		{"536dfe", "000000"},
		{"3d5afe", "000000"},
		{"304ffe", "000000"},
		{"2196f3", "000000"},
		{"e3f2fd", "000000"},
		{"bbdefb", "000000"},
		{"90caf9", "000000"},
		{"64b5f6", "000000"},
		{"42a5f5", "000000"},
		{"82b1ff", "000000"},
		{"448aff", "000000"},
		{"2979ff", "000000"},
		{"2962ff", "000000"},
		{"03a9f4", "000000"},
		{"e1f5fe", "000000"},
		{"b3e5fc", "000000"},
		{"81d4fa", "000000"},
		{"4fc3f7", "000000"},
		{"29b6f6", "000000"},
		{"80d8ff", "000000"},
		{"40c4ff", "000000"},
		{"00b0ff", "000000"},
		{"0091ea", "000000"},
		{"00bcd4", "000000"},
		{"e0f7fa", "000000"},
		{"b2ebf2", "000000"},
		{"80deea", "000000"},
		{"4dd0e1", "000000"},
		{"26c6da", "000000"},
		{"84ffff", "000000"},
		{"18ffff", "000000"},
		{"00e5ff", "000000"},
		{"00b8d4", "000000"},
		{"009688", "000000"},
		{"e0f2f1", "000000"},
		{"b2dfdb", "000000"},
		{"80cbc4", "000000"},
		{"4db6ac", "000000"},
		{"26a69a", "000000"},
		{"a7ffeb", "000000"},
		{"64ffda", "000000"},
		{"1de9b6", "000000"},
		{"00bfa5", "000000"},
		{"4caf50", "000000"},
		{"e8f5e9", "000000"},
		{"c8e6c9", "000000"},
		{"a5d6a7", "000000"},
		{"81c784", "000000"},
		{"66bb6a", "000000"},
		{"b9f6ca", "000000"},
		{"69f0ae", "000000"},
		{"00e676", "000000"},
		{"00c853", "000000"},
		{"8bc34a", "000000"},
		{"f1f8e9", "000000"},
		{"dcedc8", "000000"},
		{"c5e1a5", "000000"},
		{"aed581", "000000"},
		{"9ccc65", "000000"},
		{"ccff90", "000000"},
		{"b2ff59", "000000"},
		{"76ff03", "000000"},
		{"64dd17", "000000"},
		{"cddc39", "000000"},
		{"f9fbe7", "000000"},
		{"f0f4c3", "000000"},
		{"e6ee9c", "000000"},
		{"dce775", "000000"},
		{"d4e157", "000000"},
		{"f4ff81", "000000"},
		{"eeff41", "000000"},
		{"c6ff00", "000000"},
		{"aeea00", "000000"},
		{"ffeb3b", "000000"},
		{"fffde7", "000000"},
		{"fff9c4", "000000"},
		{"fff59d", "000000"},
		{"fff176", "000000"},
		{"ffee58", "000000"},
		{"ffff8d", "000000"},
		{"ffff00", "000000"},
		{"ffea00", "000000"},
		{"ffd600", "000000"},
		{"ffc107", "000000"},
		{"fff8e1", "000000"},
		{"ffecb3", "000000"},
		{"ffe082", "000000"},
		{"ffd54f", "000000"},
		{"ffca28", "000000"},
		{"ffe57f", "000000"},
		{"ffd740", "000000"},
		{"ffc400", "000000"},
		{"ffab00", "000000"},
		{"ff9800", "000000"},
		{"fff3e0", "000000"},
		{"ffe0b2", "000000"},
		{"ffcc80", "000000"},
		{"ffb74d", "000000"},
		{"ffa726", "000000"},
		{"ffd180", "000000"},
		{"ffab40", "000000"},
		{"ff9100", "000000"},
		{"ff6d00", "000000"},
		{"ff5722", "000000"},
		{"fbe9e7", "000000"},
		{"ffccbc", "000000"},
		{"ffab91", "000000"},
		{"ff8a65", "000000"},
		{"ff7043", "000000"},
		{"ff9e80", "000000"},
		{"ff6e40", "000000"},
		{"ff3d00", "000000"},
		{"dd2c00", "000000"},
		{"795548", "000000"},
		{"efebe9", "000000"},
		{"d7ccc8", "000000"},
		{"bcaaa4", "000000"},
		{"a1887f", "000000"},
		{"8d6e63", "000000"},
		{"607d8b", "000000"},
		{"eceff1", "000000"},
		{"cfd8dc", "000000"},
		{"b0bec5", "000000"},
		{"90a4ae", "000000"},
		{"78909c", "000000"},
	}

	rand.Seed(time.Now().UnixNano())
	color := colors[rand.Intn(len(colors))]
	return color.b, color.t
}

// getInitials extracts the first and second letter of the given name or email.
func getInitials(name string) string {
	// Remove leading and trailing spaces
	name = strings.TrimSpace(name)

	// Split the name into parts based on spaces
	nameParts := strings.Fields(name)
	// Handle cases with fewer than 2 parts
	if len(nameParts) == 0 {
		return ""
	}

	// Extract initials using runes to handle Unicode
	var initials []rune

	if len(nameParts) == 1 {
		// If there is only one part, use the first two characters
		partRunes := []rune(nameParts[0])
		if len(partRunes) >= 2 {
			initials = append(initials, unicode.ToUpper(partRunes[0]), unicode.ToUpper(partRunes[1]))
		} else if len(partRunes) == 1 {
			initials = append(initials, unicode.ToUpper(partRunes[0]))
		}
	} else {
		// For multiple parts, use the first character of the first two parts
		for _, part := range nameParts {
			if len(part) > 0 {
				runes := []rune(part)
				initials = append(initials, unicode.ToUpper(runes[0]))

				// Stop after first two initials
				if len(initials) == 2 {
					break
				}
			}
		}
	}

	return string(initials)
}

// generateSVG creates an SVG with the given name, width, and height.
func generateSVG(name string, width, height int) string {
	var buffer bytes.Buffer
	canvas := svg.New(&buffer)
	canvas.Start(width, height)

	// Get random colors for background and text
	b, t := getRandomColor()

	// Draw a rectangle with the random background color
	canvas.Rect(0, 0, width, height, "fill:#"+b)

	shortTermsName := getInitials(name)

	canvas.Text(width/2, height/2, shortTermsName,
		"fill:#"+t+"; text-anchor:middle; dominant-baseline:middle; font-size:28; font-weight:400; dy:.1em; font-family:-apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif;")
	canvas.End()
	return buffer.String()
}
