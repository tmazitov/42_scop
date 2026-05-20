# 42 Scop

A 3D model viewer built with Go and OpenGL. Loads `.obj` files with `.mtl` materials, renders them with lighting and textures, and lets you navigate the scene freely.

This project has been created as part
of the 42 curriculum by [tmazitov](https://github.com/tmazitov).

![preview](resources/Rock-Texture-Surface.jpg)

## Features

- Parse and render `.obj` / `.mtl` files with multiple objects and materials
- Free camera navigation (WASD + mouse look)
- Toggle between fill and vertex (wireframe) mode
- Auto-rotation of the model
- Automatic camera positioning based on model size

## Controls

| Input | Action |
|---|---|
| `W / A / S / D` | Move camera forward / left / backward / right |
| `Space` | Move camera up |
| `Left Shift` | Move camera down |
| Mouse drag | Look around |
| `Escape` | Exit |
| **Vertex Mode** button | Toggle wireframe / fill rendering |
| **Enable Rotation** button | Toggle auto-rotation |

## Dependencies

### System

- **OpenGL 2.1+** — comes with most GPU drivers
- **GLFW** — windowing and input library

On Ubuntu / Debian:
```bash
sudo apt-get install libgl1-mesa-dev libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev xorg-dev
```

On Fedora / RHEL:
```bash
sudo dnf install mesa-libGL-devel libXi-devel libXcursor-devel libXrandr-devel libXinerama-devel
```

### Go modules

| Package | Purpose |
|---|---|
| `github.com/go-gl/gl` | OpenGL bindings |
| `github.com/go-gl/glfw` | Window and input handling |
| `github.com/joho/godotenv` | `.env` config loading |
| `golang.org/x/image` | Texture image decoding |

## Getting started

### 1. Install Go

Requires **Go 1.18+**. Download from [go.dev/dl](https://go.dev/dl/) or via your package manager.

### 2. Clone the repository

```bash
git clone https://github.com/tmazitov/42_scop.git
cd 42_scop
```

### 3. Install Go dependencies

```bash
go mod download
```

### 4. Build

```bash
make
```

### 5. Run

Pass the path to any `.obj` file as the first argument:

```bash
./scop resources/airplane.obj
```

## Configuration (optional)

Window settings can be customised via a `.env` file in the project root.
If the file is absent, the program starts with the defaults shown below.

```env
WINDOW_TITLE=SCOP
WINDOW_HEIGHT=720
WINDOW_WIDTH=1080
ROTATION_SPEED=0.1
```

## Sample models

The `resources/` directory includes several `.obj` files you can use right away:

| File | Description |
|---|---|
| `resources/42.obj` | The 42 logo mesh with material |
| `resources/Rock1.obj` | Rock model with texture |
| `resources/cat.obj` | Fancy cat model with texture |
| `resources/airplane.obj` | Airplane model with two materials |
| `resources/teapot.obj` | Classic Utah teapot |
| `resources/teapot2.obj` | Utah teapot with material |

```bash
./scop resources/42.obj
./scop resources/airplane.obj
./scop resources/teapot.obj
```

## Project structure

```
.
├── cmd/                  # Entry point, config loading, render loop
├── internal/
│   ├── appx/             # Application, camera, controller, UI setup
│   ├── geom/             # Vertex, position, Vec3, Mat4 types
│   ├── clr/              # Color utilities
│   ├── parsing/
│   │   ├── object/       # .obj file parser
│   │   └── material/     # .mtl file parser
│   ├── rende/            # OpenGL rendering (VAO, objects, materials)
│   ├── ui/               # Buttons and text overlay
│   └── window/           # GLFW window wrapper
└── resources/            # Sample .obj / .mtl / texture files
```

## Resources

- [OBJ file format specification](http://www.martinreddy.net/gfx/3d/OBJ.spec) — original Wavefront OBJ format reference
- [MTL material format specification](http://www.paulbourke.net/dataformats/mtl/) — full list of MTL parameters and their meaning
- [OpenGL 2.1 Reference Pages](https://registry.khronos.org/OpenGL-Refpages/gl2.1/) — fixed-function pipeline API reference
- [Learn OpenGL](https://learnopengl.com/) — practical tutorials covering camera, lighting, and textures
- [Scratchapixel — Rasterization](https://www.scratchapixel.com/lessons/3d-basic-rendering/rasterization-practical-implementation) — from-scratch explanation of the rendering pipeline
- [Paul Bourke — 3D geometry](http://paulbourke.net/geometry/) — articles on normals, triangulation, and mesh processing
- [Real-Time Rendering (book)](https://www.realtimerendering.com/) — comprehensive reference for graphics algorithms
