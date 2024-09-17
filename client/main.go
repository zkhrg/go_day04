package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	candyTypes = []string{"CE", "AA", "NT", "DE", "YR"}
)

const (
	domainURL = "candy.ltd:3333"
	localURL  = "127.0.0.1:3333"
)

type BuyCandyRequest struct {
	Money      int32  `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int32  `json:"candyCount"`
}

type BuyCandy201Response struct {
	Thanks string `json:"thanks,omitempty"`
	Change int32  `json:"change,omitempty"`
}

type BuyCandy400Response struct {
	Error string `json:"error,omitempty"`
}

func main() {
	// Загрузка корневого сертификата (CA)
	executablePath, _ := os.Executable()
	executableDir := filepath.Dir(executablePath)
	candyTypesStr := strings.Join(candyTypes, " ")

	candyType := flag.String("k", "", fmt.Sprintf("provide candy type %s", candyTypesStr))
	candyCount := flag.Int("c", 0, "provide candy count")
	orderMoney := flag.Int("m", 0, "provide money for order")

	flag.Parse()

	caCert, err := os.ReadFile(executableDir + "/../certs/minica.pem")
	if err != nil {
		log.Fatalf("Ошибка чтения CA сертификата: %v", err)
	}

	// Создание пула сертификатов для проверки
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal("Не удалось добавить CA сертификат в пул")
	}

	// Загрузка клиентского сертификата и ключа
	clientCert, err := tls.LoadX509KeyPair(executableDir+"/../certs/candy.ltd/cert.pem", executableDir+"/../certs/candy.ltd/key.pem")
	if err != nil {
		log.Fatalf("Ошибка загрузки клиентского сертификата и ключа: %v", err)
	}

	// Настройка TLS-конфигурации клиента
	tlsConfig := &tls.Config{
		RootCAs:      caCertPool,                    // Указываем наш самоподписанный CA
		Certificates: []tls.Certificate{clientCert}, // Добавляем клиентский сертификат
	}

	// Настройка кастомного DialContext для резолва candy.ltd -> 127.0.0.1
	dialer := &net.Dialer{
		Timeout: 5 * time.Second,
	}

	// Кастомный DialContext для переопределения резолва домена
	customDialContext := func(ctx context.Context, network, addr string) (net.Conn, error) {
		if addr == domainURL {
			// Подменяем резолв на локальный IP
			addr = localURL
		}
		return dialer.DialContext(ctx, network, addr)
	}

	// Создаем HTTPS-клиент с кастомной TLS-конфигурацией и DialContext
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
			DialContext:     customDialContext, // Используем наш кастомный резолвер
		},
	}
	SendPostBuyCandy(client, *candyType, *orderMoney, *candyCount)
}

func SendPostBuyCandy(client *http.Client, candyType string, orderMoney, candyCount int) {
	// Подготовка данных для запроса
	requestData := BuyCandyRequest{
		Money:      int32(orderMoney),
		CandyType:  candyType,
		CandyCount: int32(candyCount),
	}

	// Преобразование данных в JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		log.Printf("Ошибка преобразования данных в JSON: %v", err)
	}

	// Создание POST-запроса
	resp, err := client.Post(
		"https://"+domainURL+"/buy_candy",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Printf("Ошибка запроса: %v", err)
		return
	}
	defer resp.Body.Close()

	// Чтение ответа
	body, err := io.ReadAll(resp.Body) // Используем io.ReadAll для чтения тела ответа
	if err != nil {
		log.Printf("Ошибка чтения ответа: %v", err)
	}
	var res string
	switch resp.StatusCode {
	case 201:
		var respp BuyCandy201Response
		json.Unmarshal(body, &respp)
		res = fmt.Sprintf("%s \nYour change is %d", respp.Thanks, respp.Change)
	case 402:
		var respp BuyCandy400Response
		json.Unmarshal(body, &respp)
		res = fmt.Sprintf("you cant buy now, reason is: %s", respp.Error)
	default:
		res = fmt.Sprintf("error occured, status code: %d | %s", resp.StatusCode, body)
	}
	fmt.Println(res)
}
