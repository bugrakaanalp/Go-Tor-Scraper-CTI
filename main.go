package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

// Ayarlar
const (
	proxyAddr     = "127.0.0.1:9150"
	targetFile    = "targets.yaml"
	outputDir     = "scraped_data"
	reportLogFile = "scan_report.log"
)

func main() {
	setupEnvironment()

	client := createTorClient()

	// Hedefleri oku
	targets, err := readTargets(targetFile)
	if err != nil {
		log.Fatalf("[FATAL] Hedef dosyası okunamadı: %v", err)
	}

	logFile, err := os.OpenFile(reportLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("[FATAL] Log dosyası oluşturulamadı: %v", err)
	}
	defer logFile.Close()

	fmt.Println("------------------------------------------------")
	fmt.Println(" CTI TOR SCRAPER - BAŞLATILIYOR... ")
	fmt.Println("------------------------------------------------")

	for _, url := range targets {
		processURL(client, url, logFile)
	}

	fmt.Println("\n[BİLGİ] Tarama tamamlandı. Sonuçlar 'scan_report.log' dosyasına işlendi.")
}

// setupEnvironment
func setupEnvironment() {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("[FATAL] Klasör hatası: %v", err)
	}
}

// createTorClient
func createTorClient() *http.Client {
	dialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
	if err != nil {
		log.Fatalf("[FATAL] Proxy bağlantı hatası: %v", err)
	}

	return &http.Client{
		Transport: &http.Transport{Dial: dialer.Dial},
		Timeout:   30 * time.Second, // (Timeout)
	}
}

// readTargets
func readTargets(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var targets []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			targets = append(targets, line)
		}
	}
	return targets, scanner.Err()
}

// processURL
func processURL(client *http.Client, url string, logFile *os.File) {
	// Ben Google Chrome olarak doğdum abi kısmı
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logError(logFile, url, "Request Oluşturma Hatası")
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	fmt.Printf("[SCAN] %s ... ", url)

	// İsteği gönder
	resp, err := client.Do(req)
	// Hata Toleransı
	if err != nil {
		fmt.Println("-> [FAIL]")
		logError(logFile, url, fmt.Sprintf("Bağlantı Hatası: %v", err))
		return
	}
	defer resp.Body.Close()

	// Veriyi kaydet
	if resp.StatusCode == 200 {
		saveHTML(url, resp.Body)
		fmt.Println("-> [SUCCESS]")
		logSuccess(logFile, url, resp.ContentLength)
	} else {
		fmt.Printf("-> [HTTP %d]\n", resp.StatusCode)
		logError(logFile, url, fmt.Sprintf("HTTP Status: %d", resp.StatusCode))
	}
}

// saveHTML: HTML içeriğini dosyaya yazar.
func saveHTML(url string, body io.Reader) {
	safeName := strings.ReplaceAll(url, "http://", "")
	safeName = strings.ReplaceAll(safeName, "https://", "")
	safeName = strings.ReplaceAll(safeName, "/", "_")
	safeName = strings.ReplaceAll(safeName, ":", "") // Windows dosya adı uyumu

	filename := filepath.Join(outputDir, safeName+".html")

	data, _ := io.ReadAll(body)
	os.WriteFile(filename, data, 0644)
}

// logSuccess
func logSuccess(file *os.File, url string, size int64) {
	msg := fmt.Sprintf("[%s] SUCCESS | URL: %s | Size: %d bytes\n", time.Now().Format(time.RFC3339), url, size)
	file.WriteString(msg)
}

// logError
func logError(file *os.File, url string, reason string) {
	msg := fmt.Sprintf("[%s] FAIL    | URL: %s | Reason: %s\n", time.Now().Format(time.RFC3339), url, reason)
	file.WriteString(msg)
}
