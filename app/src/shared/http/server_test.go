package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/SergioVenicio/desafio_nuveo/application/usecase"
	"github.com/SergioVenicio/desafio_nuveo/domain/entity"
	"github.com/SergioVenicio/desafio_nuveo/shared/repositories"
)

const baseUrl = "http://localhost:5001"
const customerEndpoint = "/cliente"

func TestCustomerEndPoint(t *testing.T) {
	customerRespository := repositories.MockCustomerRepository{}
	customerUseCase := usecase.CustomerUseCase{
		Repository: &customerRespository,
	}
	config := Config{
		Host:            "0.0.0.0",
		Port:            5001,
		CustomerUseCase: &customerUseCase,
	}
	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	go config.Run()

	t.Run("create customer", func(t *testing.T) {
		customer := entity.Customer{
			Name:    "test",
			Address: "test street, 123",
		}
		reqBody, _ := json.Marshal(customer)
		req, _ := http.NewRequest("POST", baseUrl+customerEndpoint, bytes.NewReader(reqBody))
		req.Header.Add("Content-Type", "application/json")
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
		req.Header.Add("Content-Type", "application/json")
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

	t.Run("delete customer", func(t *testing.T) {
		customer, _ := customerUseCase.Create("test", "test")
		req, _ := http.NewRequest("DELETE", baseUrl+customerEndpoint+"/"+customer.Uuid.String(), nil)
		req.Header.Add("Content-Type", "application/json")
		resp, _ := client.Do(req)

		if resp.StatusCode != http.StatusNoContent {
			t.Errorf("want %d, get %d", http.StatusNoContent, resp.StatusCode)
		}
	})

	t.Run("update customer", func(t *testing.T) {
		customer, _ := customerUseCase.Create("test", "test")
		newCustomer := entity.Customer{
			Name:    "new name",
			Address: "new addrs",
		}
		reqBody, _ := json.Marshal(newCustomer)
		req, _ := http.NewRequest("PUT", baseUrl+customerEndpoint+"/"+customer.Uuid.String(), bytes.NewReader(reqBody))
		req.Header.Add("Content-Type", "application/json")
		resp, _ := client.Do(req)

		if resp.StatusCode != http.StatusAccepted {
			t.Errorf("want %d, get %d", http.StatusNoContent, resp.StatusCode)
		}
	})
}
