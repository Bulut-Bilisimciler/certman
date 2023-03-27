package registers

import (
	"crypto/tls"
	"fmt"
	"net"

	"github.com/Bulut-Bilisimciler/certman/models"
)

func RegisterHTTPServer(dto *models.RegisterServerDto) {
	// Params come from http request as RegisterServerDto

	// Validate params

	// Create server and server related job
	server := &models.Server{
		ValidateFn: func() error {
			targetAddr := "www.google.com:443"

			// TCP bağlantısı oluşturun
			conn, err := net.Dial("tcp", targetAddr)
			if err != nil {
				fmt.Println("Bağlantı hatası:", err)
				return err
			}
			defer conn.Close()

			// TLS el sıkışması gerçekleştirin
			tlsConn := tls.Client(conn, &tls.Config{
				ServerName: "www.google.com",
			})
			if err := tlsConn.Handshake(); err != nil {
				fmt.Println("TLS el sıkısma hatası:", err)
				return err
			}

			// HTTP isteğini gönderin
			req := "GET / HTTP/1.1\r\nHost: www.google.com\r\n\r\n"
			if _, err := tlsConn.Write([]byte(req)); err != nil {
				fmt.Println("İstek gönderme hatası:", err)
				return err
			}

			return nil
		},
	}

	// attach trigger functions (onBeforeConnection, onAfterConnection, onCheckSuccess, onCheckFailure)

	// schedule job
}
