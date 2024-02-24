import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Home from './Home';

function App() {
  return (
    <div>
      <Routes>
        <Route path="/" element={<Home/>} />
        {/* Adicione mais rotas conforme necessário */}
      </Routes>
    </div>
  );
}

export default App;
