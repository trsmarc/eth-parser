package server

import (
	"encoding/json"
	"eth-parser/provider"
	store "eth-parser/txstore"
	"io"
	"net/http"
)

type Server struct {
	indexer provider.Indexer
	txStore store.TxStore
}

func NewServer(indexer provider.Indexer, txStore store.TxStore) Server {
	return Server{
		indexer: indexer,
		txStore: txStore,
	}
}

func (s *Server) StartServer() {
	// Start the server
	http.Handle("/subscribe", httpMethodCheck(http.HandlerFunc(s.subscribe), http.MethodPost))
	http.Handle("/block", httpMethodCheck(http.HandlerFunc(s.getCurrentBlock), http.MethodGet))
	http.Handle("/subscriber", httpMethodCheck(http.HandlerFunc(s.getSubscribedAddress), http.MethodGet))
	http.Handle("/transactions", httpMethodCheck(http.HandlerFunc(s.getTransactions), http.MethodGet))

	http.ListenAndServe(":8080", nil)
}

func (s *Server) subscribe(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var req SubscribeRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	err = s.txStore.Add(req.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Subscribed successfully",
	})

}

func (s *Server) getCurrentBlock(w http.ResponseWriter, r *http.Request) {
	blockNumber := s.indexer.GetCurrentBlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"blockNumber": blockNumber,
	})
}

func (s *Server) getTransactions(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var req GetTransactionRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	txs, err := s.txStore.List(req.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"total":        len(txs),
		"transactions": txs,
	})
}

func (s *Server) getSubscribedAddress(w http.ResponseWriter, r *http.Request) {
	addrs := s.txStore.Keys()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"total":     len(addrs),
		"addresses": addrs,
	})
}

// Middleware function to check HTTP methods
func httpMethodCheck(next http.Handler, allowedMethod string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowed := false

		if allowedMethod == r.Method {
			allowed = true
		}

		if !allowed {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// If the method is allowed, call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
