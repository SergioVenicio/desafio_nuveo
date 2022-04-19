package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/SergioVenicio/desafio_nuveo/domain/entity"
)

const baseUrl = "http://localhost:5000"
const customerEndpoint = "/cliente"

func TestCustomerEndPoint(t *testing.T) {
	server := HttpServer{
		Host: "0.0.0.0",
		Port: 5000,
	}
	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	go server.Run()

	t.Run("create customer", func(t *testing.T) {
		customer := entity.Customer{
			Name:    "test",
			Address: "test street, 123",
		}
		reqBody, _ := json.Marshal(customer)
		req, _ := http.NewRequest("POST", baseUrl+customerEndpoint, bytes.NewReader(reqBody))
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		if http.StatusCreated != resp.StatusCode {
			t.Errorf("want %d, get %d", http.StatusCreated, resp.StatusCode)
		}

		var respCustomer entity.Customer
		respBody, _ := io.ReadAll(resp.Body)
		json.Unmarshal(respBody, &respCustomer)
		if respCustomer.Uuid.String() == "" {
			t.Errorf("customer uuid has no value %s", respCustomer.Uuid)
		}

		if respCustomer.CreatedAt == "" {
			t.Errorf("customer createdAt has no value %s", respCustomer.CreatedAt)
		}
	})

	t.Run("list customer", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseUrl+customerEndpoint, nil)
		resp, _ := client.Do(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("want %d, get %d", http.StatusOK, resp.StatusCode)
		}

		var respCustomers []entity.Customer
		respBody, _ := io.ReadAll(resp.Body)
		json.Unmarshal(respBody, &respCustomers)

		for _, customer := range respCustomers {
			if customer.Uuid.String() == "" {
				t.Errorf("customer uuid has no value %s", customer.Uuid)
			}

			if customer.CreatedAt == "" {
				t.Errorf("customer createdAt has no value %s", customer.CreatedAt)
			}
		}
	})
}
