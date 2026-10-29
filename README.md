# 🎶 gowayback-beats

> A command-line musical time machine built in Go. Enter any date in history, scrape the top Billboard charts for that exact week, and automatically generate a nostalgic playlist straight in your Spotify account.

---

## ✨ Features

* **Temporal Music Discovery:** Input any historical date (YYYY-MM-DD) to fetch the leading tracks of that era.
* **Billboard Scraping:** Dynamically parses historical chart data without relying on paid or restrictive third-party music history APIs.
* **Seamless Spotify Integration:** Automatically authenticates and pushes the scraped tracks directly into a new, custom-named playlist in your Spotify library.
* **Lightweight & Fast:** Written entirely in Go for instant execution and minimal resource overhead.

---

## 🛠️ Prerequisites

Before running the application, make sure you have:
* **Go (v1.22 or higher)** installed on your machine.
* A **Spotify Developer Account** to generate API credentials.
    1. Go to the [Spotify Developer Dashboard](https://developer.spotify.com/dashboard).
    2. Create a new application.
    3. Note down your `SPOTIFY_CLIENT_ID` and `SPOTIFY_CLIENT_SECRET`.
    4. Set your Redirect URI to `http://localhost:8080/callback` (or your preferred local port).

---

## ⚙️ Configuration

1. Clone the repository:
   ```bash
   git clone [https://github.com/your-username/gowayback-beats.git](https://github.com/your-username/gowayback-beats.git)
   cd gowayback-beats
