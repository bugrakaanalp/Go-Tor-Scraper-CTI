# 🧅 Tor Scraper (CTI Collection Tool)

**Tor Scraper** is a **Go (Golang)** based tool developed to automate the **Collection** phase of Cyber Threat Intelligence (CTI) processes.

This tool anonymously visits specified **.onion domains or IP addresses** via the **Tor network (SOCKS5 Proxy)**, archives their **source code**, and reports their **accessibility status**.

---

## 🚀 Features

- **Tor Proxy Integration**  
  All traffic is routed through the Tor network (IP leak protection).

- **OpSec (Operational Security)**  
  Uses User-Agent spoofing to behave like a regular web browser.

- **Fault Tolerance**  
  Detects unreachable (dead) sites, logs them, and continues scanning.

- **Reporting & Archiving**  
  Stores scan results as detailed log files and HTML outputs.

---

## 🛠️ Installation

### 1️⃣ Clone the Repository
```bash
git clone https://github.com/bugrakaanalp/Go-Tor-Scraper-CTI.git
cd TorScraper
```

### 2️⃣ Install Required Dependency  
````bash
go get golang.org/x/net/proxy
````

### 3️⃣ Ensure Tor Service Is Running

Make sure Tor Service or Tor Browser is running.

Default ports:
9050 (Tor Service)
9150 (Tor Browser)

### 💻 Usage

Add target URLs or IP addresses to the targets.yaml file.

Run the tool:
````bash
go run main.go
````

📁 Scan results will be saved to the scraped_data/ directory.

### ⚠️ Legal Disclaimer

This tool is developed strictly for educational and cybersecurity research (CTI) purposes.
Any illegal or unauthorized use is solely the responsibility of the user.
