package controllers

import (
	"encoding/json"
	"fmt"
	"main/config"
	"main/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	// Decode request
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validation
	if product.Name == "" {
		http.Error(w, "Product Name is required", http.StatusBadRequest)
		return
	}
	if product.Description == "" {
		http.Error(w, "Product Description is required", http.StatusBadRequest)
		return
	}
	if product.Price <= 0 {
		http.Error(w, "Price must be greater than zero", http.StatusBadRequest)
		return
	}

	// Create product
	product.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	query := `
		INSERT INTO Products (Name, Description, Price, CreatedAt)
		OUTPUT INSERTED.Id
		VALUES (@p1, @p2, @p3, @p4)
	`
	err = config.DB.QueryRow(query, product.Name, product.Description, product.Price, product.CreatedAt).Scan(&product.ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create product: %v", err), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Product created successfully",
		"product": product,
		"status":  http.StatusOK,
	})
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	// Fetch products from DB
	var products []models.Product
	query := "SELECT Id, Name, Description, Price, CreatedAt FROM Products"
	rows, err := config.DB.Query(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch products: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to scan product: %v", err), http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Products fetched successfully",
		"products": products,
		"status":   http.StatusOK,
	})
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	// Get product ID from URL path parameter
	params := mux.Vars(r)
	id := params["id"]

	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Product ID is required",
			"status":  http.StatusBadRequest,
			"error":   true,
		})
		return
	}

	// Fetch product from DB
	var product models.Product
	query := "SELECT Id, Name, Description, Price, CreatedAt FROM Products WHERE Id = @p1"
	err := config.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt)
	if err != nil {
		// return json response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Product not found",
			"status":  http.StatusNotFound,
			"error":   true,
		})
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Product fetched successfully",
		"product": product,
		"status":  http.StatusOK,
	})
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Get product ID from URL path parameter
	params := mux.Vars(r)
	id := params["id"]

	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Product ID is required",
			"status":  http.StatusBadRequest,
			"error":   true,
		})
		return
	}

	// Fetch product from DB
	var product models.Product
	query := "SELECT Id, Name, Description, Price, CreatedAt FROM Products WHERE Id = @p1"
	err := config.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt)
	if err != nil {
		// return json response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Product not found",
			"status":  http.StatusNotFound,
			"error":   true,
		})
		return
	}

	// Decode request
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update product: %v", err), http.StatusInternalServerError)
		return
	}

	// Update product in DB
	query = "UPDATE Products SET Name = @p1, Description = @p2, Price = @p3 WHERE Id = @p4"
	_, err = config.DB.Exec(query, product.Name, product.Description, product.Price, product.ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update product: %v", err), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Product updated successfully",
		"product": product,
		"status":  http.StatusOK,
	})
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Get product ID from URL path parameter
	params := mux.Vars(r)
	id := params["id"]

	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Product ID is required",
			"status":  http.StatusBadRequest,
			"error":   true,
		})
		return
	}

	// Delete product from DB
	query := "DELETE FROM Products WHERE Id = @p1"
	_, err := config.DB.Exec(query, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete product: %v", err), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Product deleted successfully",
		"status":  http.StatusOK,
	})
}
