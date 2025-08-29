📧 Email Domain Checker


A beginner-friendly tool to quickly check if a domain is ready to send and receive secure emails.

🚀 Overview

The Email Domain Checker is a lightweight web app built with Go, HTML, CSS, and JavaScript.
It analyzes a domain and tells you:
✅ If it can receive emails
🛡️ If it’s protected from spoofing
📜 If it has spam protection policies

We use DNS records behind the scenes to make this simple:

MX records → To see if a domain can receive emails.

SPF checks → To find if it has protection against forged sender addresses.

DMARC checks → To see if it has rules to handle suspicious or spam emails.

🖥️ How It Works

Enter a domain or email (e.g., hello@gmail.com or gmail.com).

The backend extracts the domain and runs DNS checks.

The app translates technical results into easy-to-read statuses like:

📬 Receives Emails: Can this domain accept incoming mail?

🛡️ Spoof Protection: Is there protection to stop hackers pretending to be this domain?

📜 Spam Protection: Does the domain specify how suspicious emails are handled?

📂 Project Structure
email-checker/
│
├── main.go             # Go server + DNS logic
├── static/
│   ├── index.html      # Frontend HTML
│   ├── style.css       # Styling
│   └── script.js       # Frontend JavaScript
└── README.md

⚡ Quick Start
1️⃣ Clone the Repository
git clone https://github.com/your-username/email-checker.git
cd email-checker

2️⃣ Run the Server
go run main.go

3️⃣ Open in Browser
http://localhost:8080

🧩 Example Results
Input	Receives Emails	Spoof Protection	Spam Protection
gmail.com	✅ Yes	✅ Yes	✅ Yes
example.com	❌ No	❌ No	❌ No
🌟 Why Use This

Quickly see if a domain is ready for email communication

Get a simple summary of security measures

Perfect for beginners, IT students, or small business owners

🛠️ Built With

Go
 - Backend server and DNS lookups

HTML, CSS, JS - Simple and clean frontend

DNS Lookup Libraries - For domain record checks

🤝 Contributing

We welcome contributions!

Fork the repo

Create a feature branch

Submit a pull request 🎉
