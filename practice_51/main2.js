const socket = new WebSocket('ws://localhost:8080/ws');

socket.onopen = () => {
  socket.send('Hello, WebSocket!');
};

socket.onmessage = (event) => {
  console.log('Message from server:', event.data);
};
