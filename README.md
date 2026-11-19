# ASCII Art Web

ASCII Art Web is a simple web application that allows users to generate ASCII art based on different art styles and input text. Users can select an art style and provide text, and the application will generate and display ASCII art corresponding to the chosen style.

# Features
- You can select from different art styles: 
    1. Standard
    2. Shadow
    3. TinkerToy
- You can input your own text to generate ASCII Art
- There is a user-friendly web interface for easy interaction

# Project Structure 

ascii-art-web/
│
├── banners/
│   ├── shadow.txt
│   ├── standard.txt
│   └── thinkertoy.txt
│
├── static/
│   └── style.css
│
├── templates/
│   └── index.html
│
├── ascii.go
├── go.mod
├── main.go
└── README.md

# Directory and File description

- banners/ - Contains ASCII art banner templates
- shadow.txt - Shadow ASCII art font
- standard.txt - Standard ASCII art font
- thinkertoy.txt - Thinkertoy ASCII art font
- static/ - Contains static assets
- style.css - Styles the web interface
- templates/ - Contains HTML templates
- index.html - HTML page for the application
- ascii.go - Handles ASCII art conversion logic
- main.go - Application entry point and web server
- go.mod - Defines the Go module
- README.md - Project documentation and instructions

# Technologies

- Golang
- HTML
- CSS

# How to run

1. Clone the repository to your local machine
2. Install Go if you haven't already
3. Navigate to the project directory in your terminal
4. Run `go run main.go` to start the web server
5. Open your web browser and go to `http://localhost:8080` to access the application

# Usage

1. Open the web application in your browser
2. Select an art style from the provided options
3. Enter the text you want to convert to ASCII art
4. Click the "Generate ASCII Art" button to see the generated ASCII art

# Implementation Details: Algorithm

The application uses Go to receive user input, generate ASCII art, and display the result.

- User Input: The user enters text and selects a banner style (Standard, Shadow, or Thinkertoy)
- Input Validation: The application checks that the text is not empty, contains valid characters, and that the selected banner is valid
- Banner Loading: The selected banner file is read from the banners/ directory
- ASCII Art Generation: The printAscii() function processes each input character, finds its corresponding rows in the banner, and combines them to create the ASCII art
- Newline Handling: The convertNewLines() and splitLines() functions handle line breaks in the input
- Displaying the Result: The generated ASCII art is passed to the HTML template and displayed on the web page

# Authors

- Abbygail Collazo
- Aglaia Balandina