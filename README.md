# Drawing Application

A simple drawing application built with Ebiten UI in Go. This application is designed for creating lightweight drawings on both desktop and WebGL platforms. It features an intuitive interface, real-time drawing, and basic project management capabilities.

## Features

### Core Features
1. **Drawing Canvas**
   - A resizable area for freehand drawing.
   - Supports mouse input for desktop and touch input for WebGL.

2. **Brush Tools**
   - Adjustable brush sizes.
   - Multiple brush shapes (e.g., round, square).

3. **Color Picker**
   - Predefined color palette.
   - Advanced RGB/HSV color picker (stretch goal).

4. **Eraser Tool**
   - Removes strokes or parts of the drawing by replacing them with the background color.

5. **Clear Canvas**
   - Reset the entire canvas with a single click.

6. **Save to Image**
   - Save the current drawing as a PNG file.

7. **Undo/Redo (Stretch Goal)**
   - Allows users to revert or reapply changes to the canvas.

8. **Layers (Stretch Goal)**
   - Create, hide/show, and reorder layers for more advanced editing.

### UI Layout
- **Project Management**: A top menu bar with options for creating, loading, and saving projects.
- **Canvas**: Central drawing area that occupies most of the screen.
- **Tools Panel**: A vertical panel on the side for selecting brushes, colors, and other tools.
- **Layer Management Panel** (Stretch Goal): A collapsible panel to manage layers.

### Planned Stretch Goals
1. **Custom Shapes**: Stamp predefined shapes onto the canvas.
2. **Zoom and Pan**: Allow users to zoom into the canvas and move around for detailed work.
3. **Advanced Export Options**: Support for additional file formats or copy-to-clipboard functionality.

## Getting Started

### Prerequisites
- Go 1.20 or later.
- Ebiten and Ebiten UI libraries.

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/drawing-app.git
   cd drawing-app
   ```
2. Install dependencies:
   ```bash
   go get -u github.com/hajimehoshi/ebiten
   go get -u github.com/hajimehoshi/ebitenui
   ```

### Running the Application
Run the application locally with:
```bash
go run main.go
```

### Usage
- Use the **Tools Panel** to select brushes, colors, and other tools.
- Draw directly on the canvas using mouse or touch input.
- Save your drawing via the top menu bar.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License
This project is licensed under the MIT License.
