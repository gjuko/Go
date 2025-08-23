package main

import (
	"encoding/base64"

	"github.com/jchv/go-webview2"
)

func main() {
	// HTML with JavaScript to update time every second
	html := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Live Clock</title>
        <style>
            body {
                font-family: Arial, sans-serif;
                text-align: center;
                padding-top: 50px;
                background-color: #f0f0f0;
            }
            h1 {
                color: #333;
            }
            #clock {
                font-size: 24px;
                margin-top: 20px;
                color: #555;
            }
        </style>
    </head>
    <body>
        <h1>Current Date & Time</h1>
        <div id="clock">Loading...</div>

        <script>
            function updateClock() {
                const now = new Date();
                const options = {
                    weekday: 'long',
                    year: 'numeric',
                    month: 'long',
                    day: 'numeric',
                    hour: '2-digit',
                    minute: '2-digit',
                    second: '2-digit',
                    timeZoneName: 'short'
                };
                document.getElementById('clock').textContent = now.toLocaleString('en-AU', options);
            }
            setInterval(updateClock, 1000);
            updateClock(); // initial call
        </script>
    </body>
    </html>
    `

	// Encode HTML to base64
	encoded := base64.StdEncoding.EncodeToString([]byte(html))
	uri := "data:text/html;base64," + encoded

	// Launch WebView2
	w := webview2.New(true)
	defer w.Destroy()
	w.SetTitle("Live Clock — Go WebView2")
	w.SetSize(600, 400, webview2.HintNone)
	w.Navigate(uri)
	w.Run()
}
