# 42 Scop

A 3D model viewer built with Go and OpenGL. Loads `.obj` files with `.mtl` materials, renders them with lighting and textures, and lets you navigate the scene freely.

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

Requires **Go 1.24+**. Download from [go.dev/dl](https://go.dev/dl/) or via your package manager.

### 2. Clone the repository

```bash
git clone https://github.com/tmazitov/42_scop.git
cd 42_scop
```

### 3. Install Go dependencies

```bash
go mod download
```

### 4. Create a `.env` file

```bash
cp .env.example .env
```

Or create `.env` manually:

```env
OBJ_FILE_PATH=resources/42.obj
WINDOW_TITLE=SCOP
WINDOW_HEIGHT=800
WINDOW_WIDTH=800
```

`OBJ_FILE_PATH` is required. `WINDOW_TITLE`, `WINDOW_HEIGHT`, and `WINDOW_WIDTH` are optional (defaults: `500x500`).

### 5. Build and run

```bash
go build -o scop ./cmd && ./scop
```

Or use the provided script:

```bash
bash run.sh
```

## Sample models

The `resources/` directory includes several `.obj` files you can use right away:

| File | Description |
|---|---|
| `resources/42.obj` | The 42 logo mesh with material |
| `resources/Rock1.obj` | Rock model with texture |
| `resources/airplane.obj` | Airplane model with material |
| `resources/teapot.obj` | Classic Utah teapot |
| `resources/teapot2.obj` | Utah teapot with material |

To load any of them, set `OBJ_FILE_PATH` in your `.env`:

```env
OBJ_FILE_PATH=resources/airplane.obj
```

## Project structure

```
.
├── cmd/                  # Entry point, config loading, render loop
├── internal/
│   ├── appx/             # Application, camera, controller, UI setup
│   ├── geom/             # Vertex, position, normal types
│   ├── clr/              # Color utilities
│   ├── parsing/
│   │   ├── object/       # .obj file parser
│   │   └── material/     # .mtl file parser
│   ├── rende/            # OpenGL rendering (VAO, objects, materials)
│   ├── ui/               # Buttons and text overlay
│   └── window/           # GLFW window wrapper
└── resources/            # Sample .obj / .mtl / texture files
```