async function checkDomain() {
  const domain = document.getElementById("domainInput").value.trim();
  const resultDiv = document.getElementById("result");

  if (!domain) {
    alert("Please enter a domain!");
    return;
  }

  resultDiv.classList.remove("hidden");
  resultDiv.innerHTML = "🔍 Checking...";

  try {
    const response = await fetch(`/check?domain=${encodeURIComponent(domain)}`);
    const data = await response.json();

    if (data.error) {
      resultDiv.innerHTML = `⚠️ Error: ${data.error}`;
    } else {
      resultDiv.innerHTML = `
        <p><strong>Domain:</strong> ${data.domain}</p>
        <p>📬 Can receive emails: ${data.can_receive_emails}</p>
        <p>🛡️ Protected from header manipulation: ${data.header_protection}</p>
        <p>🚫 Protected against spam & fake emails: ${data.spam_protection}</p>
      `;
    }
  } catch (err) {
    resultDiv.innerHTML = `❌ Error: Could not check domain`;
  }
}
