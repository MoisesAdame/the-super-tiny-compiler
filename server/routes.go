package server

import (
	"compiler/compiler"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Index() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/compile", compileHandler)

	fmt.Println("Starting Server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type JsonResponse struct {
	RawCode string `json:"raw_code"`
}

func compileHandler(res http.ResponseWriter, req *http.Request) {
	// Handling wrong paths
	if req.URL.Path != "/compile" {
		http.Error(res, "404 NOT FOUND", http.StatusNotFound)
	}

	// Handling wrong HTTP methods
	if req.Method != "POST" {
		http.Error(res, "METHOD NOT SUPPORTED", http.StatusMethodNotAllowed)
	}

	// Decode raw code from the request
	var result JsonResponse
	err := json.NewDecoder(req.Body).Decode(&result)
	if err != nil {
		http.Error(res, "Invalid JSON", http.StatusBadRequest)
	}

	rawCode := result.RawCode

	// Tokenization, Parsing, and Compilation
	tokens := compiler.Tokenizer(rawCode)
	tokensStr := compiler.ToString(tokens)

	ast := compiler.Parser(tokens)
	astStr := ast.ToString()

	compiledResult := compiler.Compiler(rawCode)

	data := map[string]string{
		"raw-code": rawCode,
		"tokens":   tokensStr,
		"ast":      astStr,
		"res":      compiledResult,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonData)
}
