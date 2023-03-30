package ssl

import (
	"crypto/tls"
	"errors"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/Bulut-Bilisimciler/certman/models"
)

func CheckSSL(s *models.Server) error {
	// trigger before connection event
	s.OnBeforeConnection()

	// stringfy port with strconv
	port := strconv.Itoa(s.Port)

	// try to open tcp conn
	conn, err := net.Dial("tcp", s.IpAddress+":"+port)
	if err != nil {
		return err
	}
	defer conn.Close()

	// trigger after connection event
	s.OnAfterConnection()

	// prepare handshake config
	tlsConn := tls.Client(conn, &tls.Config{
		ServerName: s.ServerName,
	})

	// make handshake
	if err := tlsConn.Handshake(); err != nil {
		return err
	}

	// get certificate (count first)
	length := len(tlsConn.ConnectionState().PeerCertificates)
	if length == 0 {
		return errors.New("no certificate found")
	}

	// get first certificate
	cert := tlsConn.ConnectionState().PeerCertificates[0]

	// check certificate
	if cert.Subject.CommonName != s.ServerName {
		return errors.New("certificate subject name does not match")
	}

	// check certificate expiration
	if cert.NotAfter.Before(time.Now()) {
		return errors.New("certificate is expired")
	}

	log.Println("Certificate is valid")
	s.OnCheckSuccess()

	return nil
}
