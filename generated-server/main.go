package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"

	openapi "github.com/zkhrg/go_day04/go"
)

func main() {
	log.Printf("Server started")

	// Создание TLS-конфигурации
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13,
		ClientCAs:  loadCA(),                       // Загрузите корневой сертификат CA
		ClientAuth: tls.RequireAndVerifyClientCert, // Обязательное требование клиента
	}

	// Создание сервиса и контроллера API
	DefaultAPIService := openapi.NewDefaultAPIService()
	DefaultAPIController := openapi.NewDefaultAPIController(DefaultAPIService)
	router := openapi.NewRouter(DefaultAPIController)

	server := &http.Server{
		Addr:      ":3333",
		Handler:   router,
		TLSConfig: tlsConfig, // Установка конфигурации TLS
	}

	// Запуск сервера с указанием сертификатов
	err := server.ListenAndServeTLS("./certs/candy.ltd/cert.pem", "./certs/candy.ltd/key.pem")
	if err != nil {
		log.Fatalf("Failed to start TLS server: %v", err)
	}
}

// Функция для загрузки корневого сертификата CA
func loadCA() *x509.CertPool {
	caCert, err := os.ReadFile("./certs/minica.pem")
	if err != nil {
		log.Fatalf("Failed to read CA certificate: %v", err)
	}

	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatalf("Failed to append CA certificate")
	}
	return caCertPool
}
