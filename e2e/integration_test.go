package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

const baseURL = "http://localhost:8080"

func TestPVZReceptionFlow(t *testing.T) {
	moderatorToken := getToken(t, "moderator")
	employeeToken := getToken(t, "employee")

	pvzID := createPVZ(t, moderatorToken)

	receptionID := createReception(t, employeeToken, pvzID)

	for i := 0; i < 50; i++ {
		addProduct(t, employeeToken, pvzID, randomProductType(i))
	}

	closeReception(t, employeeToken, pvzID)

	t.Logf("Интеграционный тест пройден успешно: ПВЗ %s, Приемка %s", pvzID, receptionID)
}

func getToken(t *testing.T, role string) string {
	body := map[string]string{"role": role}
	b, _ := json.Marshal(body)
	res, err := http.Post(baseURL+"/dummyLogin", "application/json", bytes.NewBuffer(b))
	if err != nil {
		t.Fatalf("Ошибка получения токена: %v", err)
	}
	defer res.Body.Close()

	var token string
	if err := json.NewDecoder(res.Body).Decode(&token); err != nil {
		t.Fatalf("Ошибка декодирования токена: %v", err)
	}
	return token
}

func createPVZ(t *testing.T, token string) string {
	body := map[string]string{"city": "Казань"}
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", baseURL+"/pvz", bytes.NewBuffer(b))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Ошибка создания ПВЗ: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Ошибка декодирования ПВЗ: %v", err)
	}
	return result.ID
}

func createReception(t *testing.T, token, pvzID string) string {
	body := map[string]string{"pvzId": pvzID}
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", baseURL+"/receptions", bytes.NewBuffer(b))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Ошибка создания приёмки: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Ошибка декодирования приёмки: %v", err)
	}
	return result.ID
}

func addProduct(t *testing.T, token, pvzID, productType string) {
	body := map[string]string{
		"type":  productType,
		"pvzId": pvzID,
	}
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", baseURL+"/products", bytes.NewBuffer(b))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Ошибка добавления товара: %v", err)
	}
	resp.Body.Close()
}

func closeReception(t *testing.T, token, pvzID string) {
	url := fmt.Sprintf("%s/pvz/%s/close_last_reception", baseURL, pvzID)
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Ошибка закрытия приёмки: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Приёмка не закрыта, код: %d", resp.StatusCode)
	}
}

func randomProductType(i int) string {
	switch i % 3 {
	case 0:
		return "электроника"
	case 1:
		return "одежда"
	default:
		return "обувь"
	}
}
