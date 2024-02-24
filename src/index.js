import React from 'react';
import ReactDOM from 'react-dom/client'; // Importa de react-dom/client
import { BrowserRouter } from 'react-router-dom';
import App from './App';

// Use createRoot da maneira correta
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </React.StrictMode>
);
