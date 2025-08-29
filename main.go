package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

// Response structure for frontend
type EmailCheckResponse struct {
	Domain                 string `json:"domain"`
	CanReceiveEmails       string `json:"can_receive_emails"`
	HeaderProtection       string `json:"header_protection"`
	SpamProtection         string `json:"spam_protection"`
	Error                  string `json:"error,omitempty"`
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/check", handleCheck)

	fmt.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCheck(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")
	if domain == "" {
		http.Error(w, "Missing domain parameter", http.StatusBadRequest)
		return
	}

	resp := checkDomain(domain)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func checkDomain(domain string) EmailCheckResponse {
	var (
		canReceiveEmails = "❌ No"
		headerProtection = "❌ Not configured"
		spamProtection   = "❌ Not configured"
	)

	// Check MX (mail server) records
	mxRecords, err := net.LookupMX(domain)
	if err == nil && len(mxRecords) > 0 {
		canReceiveEmails = "✅ Yes"
	}

	// Check SPF records
	txtRecords, _ := net.LookupTXT(domain)
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			headerProtection = "✅ Configured correctly"
			break
		}
	}

	// Check DMARC records
	dmarcRecords, _ := net.LookupTXT("_dmarc." + domain)
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			spamProtection = "✅ Configured correctly"
			break
		}
	}

	return EmailCheckResponse{
		Domain:           domain,
		CanReceiveEmails: canReceiveEmails,
		HeaderProtection: headerProtection,
		SpamProtection:   spamProtection,
	}
}
