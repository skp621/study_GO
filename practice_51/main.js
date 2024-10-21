fetch('http://localhost:8080/api/hello')
  .then(response => response.json())
  .then(data => console.log(data.message))  // "Hello, World!"がコンソールに出力される
  .catch(error => console.error('Error:', error));
