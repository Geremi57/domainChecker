# 📧 Email Domain Checker

A lightweight tool that checks if a domain is ready to send and receive emails. Built with **Go, HTML, CSS, and JavaScript**, this project focuses on DNS lookups, email security checks, and delivering results in simple, non-technical language.

---

## 🚀 Features

- **Domain Email Check:**  
  Quickly find out if a domain can accept incoming emails.

- **Spoof Protection Check:**  
  Detects if the domain has security settings to prevent fake senders.

- **Spam Protection Check:**  
  Identifies if spam policies are in place to reduce junk mail.

- **Simple & Clear UI:**  
  All results are displayed in plain English for easy understanding.

- **Fast DNS Lookups:**  
  The backend uses Go to perform lightning-fast DNS queries.

---

## 🌐 Live Demo

_Coming soon 🚀_

---

## 🛠️ Technologies Used

- **Go** (Backend & DNS queries)  
- **HTML, CSS, JavaScript** (Frontend)  
- **DNS Lookup Libraries** (MX, SPF, DMARC checks)

---

## 🖥️ How It Works

1. Enter a domain or email (e.g., `example.com` or `hello@example.com`).  
2. The app extracts the domain and checks for DNS records:
   - **MX Records** → Email server setup check  
   - **SPF Records** → Spoofing prevention  
   - **DMARC Records** → Spam filtering policies  
3. The results are shown in clear text, not technical jargon.

---

## ⚡ Quick Start

### 1. Clone the Repo
```bash
git clone https://github.com/your-username/email-domain-checker.git
cd email-domain-checker
2. Run the Server
go run main.go
```

3. Open in Browser
http://localhost:8080

📂 Project Structure
email-domain-checker/
│

├── main.go             # Go server + DNS logic

├── static/

│   ├── index.html      # Frontend HTML

│   ├── style.css       # Styling

│   └── script.js       # Frontend logic

└── README.md



🧩 Example Results
Domain	Receives Emails	Spoof Protection	Spam Protection
gmail.com	✅ Yes	✅ Yes	✅ Yes
example.com	❌ No	❌ No	❌ No
🤝 Contributing

We welcome contributions!

Fork the repo

Create a feature branch

Submit a pull request 🎉
