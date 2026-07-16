# Security News Scraper

> **A blazing-fast, keyless security news and CVE intelligence engine built with Go.**
> Created and maintained by **Harsh Sonker**.

<p align="center">
  <img src="assets/dashboard.png?v=2"
       alt="Security News Scraper Terminal UI"
       width="100%">
</p>

---

## Overview

Every day, thousands of cybersecurity articles are published. The challenge isn't collecting them—it's identifying the few stories that actually require immediate attention.

**Security News Scraper** aggregates security news from multiple sources, clusters duplicate stories, enriches them with vulnerability intelligence, and ranks them using deterministic scoring based on:

- 🕒 Recency
- 📈 News velocity
- 🚨 CVSS severity
- 🎯 FIRST EPSS probability
- 🔥 CISA Known Exploited Vulnerabilities (KEV)

Everything works **without requiring API keys**.

---

## Features

- ⚡ **Blazing-Fast Ingestion (36+ Built-in Sources)**
  - Out-of-the-box support for CISA, Project Zero, Talos, Unit 42, CrowdStrike, and more
  - Concurrent RSS/Atom feed scraping
  - Per-host rate limiting
  - HTTP caching with `304 Not Modified`

- 🔍 **Keyless CVE Intelligence**
  - CVEList v5
  - CISA KEV
  - FIRST EPSS

- 📰 **Smart Story Clustering**
  - Union-Find clustering
  - Jaccard similarity matching
  - Automatic duplicate detection

- 🎨 **Premium Terminal Dashboard**
  - Custom boxed layout with clean borders, header pills, and persistent status bar
  - **Dynamic Sorting** (`s` key): Instantly reorganize feed by Score, Time, Outlets, or CVE Severity
  - Color-coded severity
  - Story drill-down
  - Open articles directly in the browser

- 🤖 **Local AI Summaries (Optional)**
  - Ollama
  - Qwen
  - Cloud LLM providers

- ⏰ **Watch Mode**
  - Background monitoring
  - Webhook notifications
  - Alerts for newly emerging high-signal stories

---

# Installation

## Prerequisites

- Go 1.25+
- Git

---

## Clone

```bash
git clone https://github.com/Harsh-Sonker/security-news-scraper.git

cd security-news-scraper
```

---

## Build

### Install globally

```bash
go install ./cmd/security-news-scraper
```

### Or build a local binary

```bash
go build -o security-news-scraper ./cmd/security-news-scraper
```

---

# Quick Start

### Scrape feeds

```bash
go run ./cmd/security-news-scraper scrape
```

### Launch the Terminal UI

```bash
go run ./cmd/security-news-scraper tui
```

#### TUI Keyboard Shortcuts
- `↑ / k` : Move cursor up
- `↓ / j` : Move cursor down
- `s` : Cycle dynamic sorting mode (Score -> Time -> Outlets -> CVE)
- `Enter` : View full story details
- `o` : Open original article in default web browser
- `q` : Quit dashboard

---

# CLI Examples

Generate a Markdown report:

```bash
security-news-scraper digest --top 20
```

Lookup a CVE:

```bash
security-news-scraper cve CVE-2021-44228
```

Run continuous monitoring:

```bash
security-news-scraper watch --interval 1h
```

---

# Architecture

```
             Security Feeds
(Krebs, CISA, BleepingComputer, etc.)
                  │
                  ▼
       Concurrent HTTP Fetcher
                  │
                  ▼
        Parse & Deduplicate
                  │
                  ▼
        Keyless CVE Enrichment
                  │
                  ▼
       Deterministic Ranking
                  │
                  ▼
         SQLite (Embedded)
                  │
                  ▼
         Bubble Tea Terminal UI
```

---

# Configuration

The Terminal UI uses a built-in **Dracula-inspired** color palette.

To customize it, edit:

```
internal/tui/theme.go
```

and rebuild the application.

---

# Contributing

Contributions are welcome.

Ideas include:

- New RSS feed sources
- Ranking improvements
- Additional enrichers
- UI enhancements
- Performance optimizations
- Bug fixes

Feel free to open an issue or submit a Pull Request.

---

# License

Licensed under the **GNU AGPL v3.0**.

See the [LICENSE](LICENSE) file for details.

---

<div align="center">

Built with ❤️ in Go by **Harsh Sonker**

</div>