package materialParsing

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-gl/gl/v2.1/gl"
)

func isPath(s string) bool {
	return strings.Contains(s, "/") ||
		strings.Contains(s, "\\") ||
		strings.Contains(s, string(filepath.Separator))
}

// texturePath resolves and returns the absolute path to the texture file.
// In map_* lines, options like "-blendu on" may precede the filename — the
// filename is always the last argument, so args is args[1:] from the handler.
func resolveTexturePath(sourcePath string, args []string) string {
	raw := args[len(args)-1]
	if isPath(raw) {
		return raw
	}
	return filepath.Join(filepath.Dir(sourcePath), raw)
}

func loadTexture(path string) (uint32, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("failed to open texture '%s': %w", path, err)
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(path))
	var img image.Image
	switch ext {
	case ".png":
		img, err = png.Decode(file)
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	default:
		return 0, fmt.Errorf("unsupported texture format '%s'", ext)
	}
	if err != nil {
		return 0, fmt.Errorf("failed to decode texture '%s': %w", path, err)
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var textureID uint32
	gl.GenTextures(1, &textureID)
	gl.BindTexture(gl.TEXTURE_2D, textureID)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexImage2D(
		gl.TEXTURE_2D, 0, gl.RGBA,
		int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y),
		0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix),
	)

	return textureID, nil
}
