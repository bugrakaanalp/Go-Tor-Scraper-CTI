\# 🧅 Tor Scraper (CTI Collection Tool)



\*\*Tor Scraper\*\*, siber tehdit istihbaratı (CTI) süreçlerinde "Collection" (Veri Toplama) aşamasını otomatize etmek için geliştirilmiş Go tabanlı bir araçtır. 



Bu araç, belirlenen .onion uzantılı hedefleri veya IP adreslerini Tor ağı (SOCKS5 Proxy) üzerinden anonim olarak ziyaret eder, kaynak kodlarını arşivler ve erişim durumlarını raporlar.



\## 🚀 Özellikler

\- \*\*Tor Proxy Entegrasyonu:\*\* Tüm trafik Tor ağı üzerinden geçer (IP Sızıntısı Koruması).

\- \*\*OpSec (Operasyonel Güvenlik):\*\* User-Agent spoofing ile kendini normal bir tarayıcı gibi gösterir.

\- \*\*Hata Toleransı:\*\* Kapanmış (dead) siteleri tespit eder, loglar ve taramaya devam eder.

\- \*\*Raporlama:\*\* Tarama sonuçlarını detaylı bir log dosyasına ve HTML formatında arşivler.



\## 🛠️ Kurulum



1\. Bu depoyu klonlayın:

&nbsp;  ```bash

&nbsp;  git clone https://github.com/bugrakaanalp/Go-Tor-Scraper-CTI.git

&nbsp;  cd TorScraper


2. Gerekli kütüphaneyi indirin:

&nbsp;  ```bash
   go get golang.org/x/net/Proxy



3\. Tor Servisinin (veya Tor Browser) çalıştığından emin olun (Port: 9150 veya 9050).



\## 💻 Kullanım

1. targets.yaml dosyasına hedef URL'leri ekleyin.

2. Aracı çalıştırın:
   ```bash
   go run main.go

Sonuçlar scraped\_data/ klasörüne kaydedilecektir.

## ⚠️ Yasal Uyarı

Bu araç sadece \*\*eğitim ve siber güvenlik araştırmaları\*\* (CTI) amacıyla geliştirilmiştir. Yasadışı amaçlarla kullanılması durumunda sorumluluk kullanıcıya aittir.

