async function handleSubmit() {
    const raw_code = document.getElementById('raw-code').value;
  
    fetch("http://localhost:8080/compile", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ raw_code: raw_code })
    })
    .then(response => {
      if (!response.ok) {
        throw new Error("Network response was not ok " + response.statusText);
      }
      return response.json();
    })
    .then(data => {
      const keys = ["raw-code", "tokens", "ast", "res"];
      const titles = {
        "raw-code": "Raw Code",
        "tokens": "Tokens",
        "ast": "Abstract Syntax Tree",
        "res": "Compiled Result"
      };
      const responseTitle = document.getElementById('response-title');
      const responseText = document.getElementById('response-text');
  
      keys.forEach((key, index) => {
        setTimeout(() => {
          responseTitle.innerHTML = titles[key];  // Corrected to titles[key]
          responseText.innerHTML = data[key];
        }, 2000 * (index + 1));  // Delay increases with each key
      });
    })
    .catch(error => {
      console.error("Error:", error);
    });
  }
  