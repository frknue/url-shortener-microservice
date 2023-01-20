
const submitButton = document.getElementById('submitButton');

submitButton.addEventListener('click', function() {
  const urlInput = document.getElementById('url_input').value;

  fetch('/api/shorturl', { 
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ url: urlInput })
  })
});
