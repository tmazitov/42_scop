package materialParsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
	"github.com/go-gl/gl/v2.1/gl"
	"strings"
	"fmt"
    "image"
    "errors"
    "image/png"
    "image/draw"
    "image/jpeg"
    "os"
    "path/filepath"
)

func isPath(s string) bool {
    // Check for common path separators
    return strings.Contains(s, "/") || 
           strings.Contains(s, "\\") ||
           strings.Contains(s, string(filepath.Separator))
}

func diffuseTextureMapHandler(material *rende.Material, args []string) error {

	if len(args) != 2 {
		return ErrInvalidDiffuseTextureMapLine
	}
	
	texturePath := args[1]

	if !isPath(texturePath) {
		texturePath = filepath.Join(filepath.Dir(material.SourcePath), texturePath)
	}

	file, err := os.Open(texturePath)
    if err != nil {
        return fmt.Errorf("failed to open texture %s: %w", texturePath, err)
    }
    defer file.Close()

    // Decode PNG
    var bounds image.Rectangle
    var img image.Image

    if strings.Contains(texturePath, ".png") {
        img, err = png.Decode(file)
        if err != nil {
            return fmt.Errorf("failed to decode PNG %s: %w", texturePath, err)
        }
        bounds = img.Bounds()
    } else if strings.Contains(texturePath, ".jpeg") || strings.Contains(texturePath, ".jpg") {
        img, err = jpeg.Decode(file)
        if err != nil {
            return fmt.Errorf("failed to decode JPG/JPEG %s: %w", texturePath, err)
        }
        bounds = img.Bounds()
    } else {
        return errors.New("failed to decode image: invalid file type")
    }

    
    // Convert to RGBA
    rgba := image.NewRGBA(bounds)
    draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
    
    // Generate OpenGL texture
    var textureID uint32
    gl.GenTextures(1, &textureID)
    gl.BindTexture(gl.TEXTURE_2D, textureID)
    
    // Set texture parameters
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
    
    // Upload texture data to GPU
    gl.TexImage2D(
        gl.TEXTURE_2D,
        0,
        gl.RGBA,
        int32(rgba.Rect.Size().X),
        int32(rgba.Rect.Size().Y),
        0,
        gl.RGBA,
        gl.UNSIGNED_BYTE,
        gl.Ptr(rgba.Pix),
    )
    
    material.SetTextureId(textureID)

	return nil
}


