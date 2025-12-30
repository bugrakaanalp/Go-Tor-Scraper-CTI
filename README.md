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
