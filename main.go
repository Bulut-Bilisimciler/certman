package main

import (
	"crypto/tls"
	"fmt"
	"net"
)

func main() {
	// Hedef adresi ve portu belirleyin
	targetAddr := "www.google.com:443"

	// TCP bağlantısı oluşturun
	conn, err := net.Dial("tcp", targetAddr)
	if err != nil {
		fmt.Println("Bağlantı hatası:", err)
		return
	}
	defer conn.Close()

	// TLS el sıkışması gerçekleştirin
	tlsConn := tls.Client(conn, &tls.Config{
		ServerName: "www.google.com",
	})
	if err := tlsConn.Handshake(); err != nil {
		fmt.Println("TLS el sıkısma hatası:", err)
		return
	}

	// HTTP isteğini gönderin
	req := "GET / HTTP/1.1\r\nHost: www.google.com\r\n\r\n"
	if _, err := tlsConn.Write([]byte(req)); err != nil {
		fmt.Println("İstek gönderme hatası:", err)
		return
	}

	// Sunucudan gelen cevabı okuyun
	buf := make([]byte, 1024)
	if _, err := tlsConn.Read(buf); err != nil {
		fmt.Println("Cevap alma hatası:", err)
		return
	}

	// Cevabı yazdırın
	fmt.Println(string(buf))

	// SSL sertifikası hakkında bilgi alın
	cert := tlsConn.ConnectionState().PeerCertificates[0]
	fmt.Println("Sertifika konusu:", cert.Subject.CommonName)
	fmt.Println("Sertifika son kullanma tarihi:", cert.NotAfter)
}
