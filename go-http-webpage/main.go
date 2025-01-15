package main

import (
	"fmt"
	"net/http"
)

// Handler for the root endpoint
func handler(w http.ResponseWriter, r *http.Request) {
	// Serve the HTML page without using format specifiers
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Enhanced Go Webpage</title>
			<style>
				body {
					font-family: 'Arial', sans-serif;
					background-color: #f0f8ff;
					text-align: center;
					margin: 0;
					padding: 0;
				}
				.container {
					margin-top: 50px;
					padding: 20px;
					background-color: #ffffff;
					box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
					border-radius: 8px;
					width: 80%;
					max-width: 600px;
					margin-left: auto;
					margin-right: auto;
				}
				h1 {
					color: #4CAF50;
				}
				p {
					color: #555;
					font-size: 18px;
				}
				button {
					padding: 10px 20px;
					background-color: #4CAF50;
					color: white;
					border: none;
					border-radius: 5px;
					cursor: pointer;
					font-size: 16px;
				}
				button:hover {
					background-color: #45a049;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Welcome to the Enhanced Go Webpage</h1>
				<p>This is a simple Go server serving a webpage with enhanced UI.</p>
				<button onclick="changeText()">Click Me!</button>
				<p id="dynamic-text">Hello, click the button to change this text.</p>
			</div>

			<script>
				function changeText() {
					document.getElementById("dynamic-text").innerText = "You clicked the button! Nice job!";
				}
			</script>
		</body>
		</html>
	`)
}

func main() {
	// Set up the HTTP server with a route and handler
	http.HandleFunc("/", handler)

	// Start the server on port 8080
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
