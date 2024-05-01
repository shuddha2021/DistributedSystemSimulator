package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"
	"time"
)

// NodeData represents the data structure for a single node in the distributed system.
type NodeData struct {
	ID    int       `json:"id"`
	Name  string    `json:"name"`
	Value int       `json:"value"`
	Time  time.Time `json:"time"`
}

// Simulate a set of nodes in a distributed system.
var (
	nodeCount = 5            // Number of nodes in the simulated system.
	nodes     []NodeData     // Slice to hold node data.
	mutex     sync.RWMutex   // RWMutex for thread-safe data access.
	wg        sync.WaitGroup // WaitGroup for goroutine synchronization.
)

// InitNodes initializes a set of nodes with random data.
func InitNodes() {
	mutex.Lock()
	defer mutex.Unlock()

	nodes = make([]NodeData, nodeCount)
	for j := 0; j < nodeCount; j++ {
		nodes[j] = NodeData{
			ID:    j,
			Name:  fmt.Sprintf("Node-%d", j),
			Value: rand.Intn(100),
			Time:  time.Now(),
		}
	}
}

// GetNodeData handles HTTP requests to retrieve node data.
func GetNodeData(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	defer mutex.RUnlock()

	data, err := json.Marshal(nodes)
	if err != nil {
		log.Printf("Failed to marshal data: %v", err)
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to write data: %v", err)
	}
}

// RootHandler provides a welcome message at the root endpoint.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	message := map[string]string{
		"message": "Welcome to the Distributed System Simulator! Visit /nodes to get node data.",
	}

	data, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to marshal message: %v", err)
		http.Error(w, "Failed to marshal message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to write message: %v", err)
	}
}

// UpdateNode updates a random node with new data.
func UpdateNode() {
	mutex.Lock()
	defer mutex.Unlock()

	index := rand.Intn(nodeCount)
	nodes[index].Value = rand.Intn(100)
	nodes[index].Time = time.Now()
}

func main() {
	// Initialize the nodes with random data.
	InitNodes()

	// HTTP server setup.
	http.HandleFunc("/", RootHandler)      // Root endpoint with a welcome message
	http.HandleFunc("/nodes", GetNodeData) // Endpoint for node data

	// Periodically update a random node using goroutines.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			UpdateNode()
			time.Sleep(5 * time.Second)
		}
	}()

	// Start the HTTP server.
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// Wait for the goroutine to finish.
	wg.Wait()
}

// TestGetNodeData tests the behavior of the GetNodeData function.
func TestGetNodeData(t *testing.T) {
	// Initialize test data.
	InitNodes()

	// Create a test HTTP request.
	req, err := http.NewRequest("GET", "/nodes", nil)
	if err != nil {
		t.Fatalf("Failed to create test request: %v", err)
	}

	// Create a ResponseRecorder to capture the response.
	rr := httptest.NewRecorder()

	// Call the handler function.
	handler := http.HandlerFunc(GetNodeData)
	handler.ServeHTTP(rr, req)

	// Check the response status code.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, status)
	}

	// Check the response body.
	var respNodes []NodeData
	err = json.Unmarshal(rr.Body.Bytes(), &respNodes)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	// Check if the response nodes match the expected nodes.
	if !reflect.DeepEqual(respNodes, nodes) {
		t.Errorf("Response nodes do not match expected nodes")
	}
}

// TestRootHandler tests the behavior of the RootHandler function.
func TestRootHandler(t *testing.T) {
	// Create a test HTTP request.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create test request: %v", err)
	}

	// Create a ResponseRecorder to capture the response.
	rr := httptest.NewRecorder()

	// Call the handler function.
	handler := http.HandlerFunc(RootHandler)
	handler.ServeHTTP(rr, req)

	// Check the response status code.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, status)
	}

	// Check the response body.
	var respMessage map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &respMessage)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	expectedMessage := map[string]string{
		"message": "Welcome to the Distributed System Simulator! Visit /nodes to get node data.",
	}

	// Check if the response message matches the expected message.
	if !reflect.DeepEqual(respMessage, expectedMessage) {
		t.Errorf("Response message does not match expected message")
	}
}

// TestUpdateNode tests the behavior of the UpdateNode function.
func TestUpdateNode(t *testing.T) {
	// Initialize test data.
	InitNodes()

	// Store the initial state of the nodes.
	initialNodes := make([]NodeData, len(nodes))
	copy(initialNodes, nodes)

	// Call the UpdateNode function.
	UpdateNode()

	// Check if at least one node has been updated.
	updated := false
	for i := range nodes {
		if !reflect.DeepEqual(nodes[i], initialNodes[i]) {
			updated = true
			break
		}
	}

	if !updated {
		t.Error("No node was updated by the UpdateNode function")
	}
}
