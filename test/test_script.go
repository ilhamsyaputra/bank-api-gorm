package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"test/data/response"
)

var (
	BASE_URL = "http://localhost:2525/v1"
)

func main() {
	token, err := login()
	if err != nil {
		fmt.Println(err)
	}

	testScenario1(token)
	testScenario2(token)
	testScenario3(token)
	testScenario4(token)
}

func login() (token string, err error) {
	url := BASE_URL + "/login"

	loginBody := struct {
		NoNasabah string `json:"no_nasabah"`
		Pin       string `json:"pin"`
	}{
		NoNasabah: "001000001",
		Pin:       "0852",
	}

	body, err := json.Marshal(loginBody)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var response response.LoginResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if response.Code != 201 {
		err = fmt.Errorf("terjadi kesalahan")
		return
	}

	return response.Data.Token, nil
}

func testScenario1(token string) {
	// Request tabung nomor rekening dikenali
	fmt.Println("TEST SCENARIO 1 START")
	fmt.Println("Test Scenario: Request tabung nomor rekening dikenali")

	defer func() {
		fmt.Println("TEST SCENARIO 1 END")
		fmt.Println("========================")
	}()

	url := BASE_URL + "/tabung"

	tabungBody := struct {
		NoRekening string  `json:"no_rekening"`
		Nominal    float64 `json:"nominal"`
	}{
		NoRekening: "9900100000001",
		Nominal:    1000000,
	}

	body, err := json.Marshal(tabungBody)
	if err != nil {
		fmt.Println("Gagal marshal, error:", err.Error())
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		fmt.Println("Gagal menyiapkan request, error:", err.Error())
		return
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Gagal melakukan request, error:", err.Error())
		return
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Gagal membaca response, error:", err.Error())
		return
	}

	fmt.Println(string(respBody))

	var response response.TabungResponseSuccess
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Println("Gagal unmarshal response, error:", err.Error())
		return
	}

	if response.Code == 401 {
		fmt.Println("FAILED:", response.Remark)
		return
	}

	fmt.Println("SCENARIO PASSED")
}

func testScenario2(token string) {
	// Request tabung nomor rekening dikenali
	fmt.Println("TEST SCENARIO 2 START")
	fmt.Println("Test Scenario: Request tabung nomor rekening tidak dikenali")

	defer func() {
		fmt.Println("TEST SCENARIO 2 END")
		fmt.Println("========================")
	}()

	url := BASE_URL + "/tabung"

	tabungBody := struct {
		NoRekening string  `json:"no_rekening"`
		Nominal    float64 `json:"nominal"`
	}{
		NoRekening: "9900100000001222",
		Nominal:    1000000,
	}

	body, err := json.Marshal(tabungBody)
	if err != nil {
		fmt.Println("Gagal marshal, error:", err.Error())
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		fmt.Println("Gagal menyiapkan request, error:", err.Error())
		return
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Gagal melakukan request, error:", err.Error())
		return
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Gagal membaca response, error:", err.Error())
		return
	}

	fmt.Println(string(respBody))

	var response response.TabungResponseSuccess
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Println("Gagal unmarshal response, error:", err.Error())
		return
	}

	if response.Code == 401 && response.Remark == "tidak memiliki izin untuk melanjutkan proses" {
		fmt.Println("SCENARIO PASSED")
		return
	}

	fmt.Println("FAILED:", response.Remark)
}

func testScenario3(token string) {
	// Request tabung nomor rekening dikenali
	fmt.Println("TEST SCENARIO 3 START")
	fmt.Println("Test Scenario: Request tarik saldo lebih dari yang dimiliki")

	defer func() {
		fmt.Println("TEST SCENARIO 3 END")
		fmt.Println("========================")
	}()

	url := BASE_URL + "/tarik"

	tarikBody := struct {
		NoRekening string  `json:"no_rekening"`
		Nominal    float64 `json:"nominal"`
	}{
		NoRekening: "9900100000001",
		Nominal:    50000,
	}

	body, err := json.Marshal(tarikBody)
	if err != nil {
		fmt.Println("Gagal marshal, error:", err.Error())
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		fmt.Println("Gagal menyiapkan request, error:", err.Error())
		return
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Gagal melakukan request, error:", err.Error())
		return
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Gagal membaca response, error:", err.Error())
		return
	}

	fmt.Println(string(respBody))

	var response response.TabungResponseSuccess
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Println("Gagal unmarshal response, error:", err.Error())
		return
	}

	if response.Code == 201 && response.Status == "success" {
		fmt.Println("SCENARIO PASSED")
		return
	}

	fmt.Println("FAILED:", response.Remark)
}

func testScenario4(token string) {
	// Request tabung nomor rekening dikenali
	fmt.Println("TEST SCENARIO 4 START")
	fmt.Println("Test Scenario: Request tarik saldo kurang dari yang dimiliki")

	defer func() {
		fmt.Println("TEST SCENARIO 4 END")
		fmt.Println("========================")
	}()

	url := BASE_URL + "/tarik"

	tarikBody := struct {
		NoRekening string  `json:"no_rekening"`
		Nominal    float64 `json:"nominal"`
	}{
		NoRekening: "9900100000001",
		Nominal:    5000000000,
	}

	body, err := json.Marshal(tarikBody)
	if err != nil {
		fmt.Println("Gagal marshal, error:", err.Error())
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		fmt.Println("Gagal menyiapkan request, error:", err.Error())
		return
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Gagal melakukan request, error:", err.Error())
		return
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Gagal membaca response, error:", err.Error())
		return
	}

	fmt.Println(string(respBody))

	var response response.TabungResponseSuccess
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Println("Gagal unmarshal response, error:", err.Error())
		return
	}

	if response.Code == 400 && response.Remark == "saldo tidak mencukupi untuk melakukan transaksi tarik" {
		fmt.Println("SCENARIO PASSED")
		return
	}

	fmt.Println("FAILED:", response.Remark)
}
