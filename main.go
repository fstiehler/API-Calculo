package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Estrutura para a solicitação de cálculo
type CalculationRequest struct {
	Num1     float64 `json:"num1"`
	Num2     float64 `json:"num2"`
	Operator string  `json:"operator"`
}

// Estrutura para a resposta de cálculo
type CalculationResponse struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

// Funções de operação
func add(a float64, b float64) float64 {
	return a + b
}

func subtract(a float64, b float64) float64 {
	return a - b
}

func multiply(a float64, b float64) float64 {
	return a * b
}

func divide(a float64, b float64) (float64, string) {
	if b != 0 {
		return a / b, ""
	} else {
		return 0, "Erro: Divisão por zero"
	}
}

// Função de cálculo
func calculate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var req CalculationRequest
	var res CalculationResponse

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch req.Operator {
	case "+":
		res.Result = add(req.Num1, req.Num2)
	case "-":
		res.Result = subtract(req.Num1, req.Num2)
	case "*":
		res.Result = multiply(req.Num1, req.Num2)
	case "/":
		res.Result, res.Error = divide(req.Num1, req.Num2)
	default:
		res.Error = "Operador inválido"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Estrutura para a solicitação de conversão
type ConversionRequest struct {
	Liters float64 `json:"liters"`
}

// Função de conversão de litros para mililitros
func convertToMilliliters(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var req ConversionRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func main() {
	http.HandleFunc("/calculate", calculate)
	http.HandleFunc("/convertToMilliliters", convertToMilliliters)
	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}
