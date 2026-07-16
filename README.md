# Security News Scraper

```text
███████╗ ██████╗██████╗  █████╗ ██████╗ ███████╗██████╗ 
██╔════╝██╔════╝██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗
███████╗██║     ██████╔╝███████║██████╔╝█████╗  ██████╔╝
╚════██║██║     ██╔══██╗██╔══██║██╔═══╝ ██╔══╝  ██╔══██╗
███████║╚██████╗██║  ██║██║  ██║██║     ███████╗██║  ██║
╚══════╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚══════╝╚═╝  ╚═╝
```
<p align="center">
  <img src="https://img.shields.io/badge/Cybersecurity--Projects-Threat%20Intel-red?style=flat-square&logo=github" alt="Project" />
  <img src="https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat-square&logo=go" alt="Go Version" />
  <img src="https://img.shields.io/badge/License-AGPL%20v3-purple?style=flat-square" alt="License" />
  <img src="https://img.shields.io/badge/MITRE%20ATT&CK-Reconnaissance-orange?style=flat-square" alt="MITRE" />
</p>

> **A blazing-fast, keyless cyber threat intelligence engine built with Go.**
> Created and maintained by **Harsh Sonker**.

<p align="center">
  <img src="assets/dashboard.png?v=2"
       alt="Security News Scraper Terminal UI"
       width="100%">
</p>

---

## 🚨 The Problem: The Cybersecurity Noise
Every day, thousands of security articles, vulnerability disclosures, and exploit proofs-of-concept are published. The problem isn't a lack of information—it’s an overwhelming surplus of it. 

When a critical vulnerability drops, 20 different outlets write about it within the hour. Most security dashboards simply aggregate these links, leaving you to manually sift through duplicate stories and check if the vulnerability is actually being exploited in the wild.

I built this engine to solve this exact problem: it automatically collects news, deduplicates it, enriches it with severity metrics, and ranks it mathematically.

---

## 📚 Terminology Glossary (Beginner's Guide)
If you're new to cybersecurity or Go, here are the core concepts used in this project:
- **Go (Golang):** A programming language known for being incredibly fast and handling "concurrency" (doing thousands of things at exactly the same time).
- **RSS/Atom:** Standard web formats used to publish frequently updated works, like news headlines or blog posts.
- **CVE (Common Vulnerabilities and Exposures):** A unique ID for a specific flaw (e.g., `CVE-2024-1234`).
- **CVSS (Common Vulnerability Scoring System):** A severity score from `0.0` to `10.0`. A `9.8` means the flaw is critical and easy to exploit.
- **EPSS (Exploit Prediction Scoring System):** A statistical model predicting the probability (from 0 to 100%) that a vulnerability will actually be exploited in the wild.
- **KEV (Known Exploited Vulnerabilities):** A catalog maintained by CISA listing vulnerabilities that hackers are actively exploiting.
- **TUI (Terminal User Interface):** A visual interface constructed entirely within a text-based terminal, rather than a web browser.

---

## 🛠️ How the Engine Works

### 1. Gathering the Data (36+ Built-in Sources)
The engine reaches out to an expansive network of **36 top-tier security feeds**—ranging from government advisories (CISA) to elite research teams (Google Project Zero, Cisco Talos, Unit 42) and breaking news outlets.
Because pulling 36 websites one-by-one is slow, the engine uses Go’s concurrency model to grab them simultaneously. It utilizes smart HTTP caching to honor `304 Not Modified` headers, meaning it skips feeds that haven't updated to save bandwidth.

### 2. Stopping the Echo Chamber (Deduplication)
The engine uses a mathematical algorithm called **Jaccard Similarity** to compare the text overlap of headlines. If it decides two headlines are talking about the same event, it uses a **Union-Find** data structure to group them together. Instead of seeing 10 separate rows for the same zero-day, you see a single "Cluster" that shows it was corroborated by 10 different sources.

### 3. Keyless CVE Enrichment
A headline containing a CVE tag is useless without context. Most tools require you to sign up for paid API keys. I designed a **keyless architecture** that reaches out to open public databases. It scans the article, finds the CVE tag, and automatically pulls the CVSS Score, the EPSS Score, and checks the CISA KEV catalog.

### 4. The Premium TUI and Dynamic Sorting
Finally, all this data is presented in a rich, color-coded, keyboard-driven dashboard built using the `charmbracelet/bubbletea` framework. 
By pressing the `s` key, you can dynamically sort your threat feed based on your immediate needs:
- **SCORE**: A mathematical algorithmic rank.
- **TIME**: A chronological feed for breaking news.
- **OUTLETS**: Sorts by the number of sources writing about it.
- **CVE SEVERITY**: Brings the most dangerous, highest CVSS-rated vulnerabilities to the top.

---

# 🚀 Installation

## Prerequisites
- Go 1.25+
- Git

## Clone
```bash
git clone https://github.com/Harsh-Sonker/security-news-scraper.git
cd security-news-scraper
```

## Build
```bash
# Install globally
go install ./cmd/security-news-scraper

# Or build a local binary
go build -o security-news-scraper ./cmd/security-news-scraper
```

---

# 💻 Quick Start

### 1. Scrape feeds
```bash
go run ./cmd/security-news-scraper scrape
```

### 2. Launch the Terminal UI
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

# ⚙️ CLI Examples

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

# 🏗️ Architecture

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

# 🎨 Configuration

The Terminal UI uses a built-in **Dracula-inspired** color palette.
To customize it, edit:
```
internal/tui/theme.go
```
and rebuild the application.

---

# 🤝 Contributing
Contributions are welcome. Feel free to open an issue or submit a Pull Request.

# 📄 License
Licensed under the **GNU AGPL v3.0**. See the [LICENSE](LICENSE) file for details.

---

<div align="center">

Built with ❤️ in Go by **Harsh Sonker**

</div>