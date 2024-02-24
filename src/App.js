import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Home from './Home';
import Searcher from './screens/Searcher'; 

function App() {
  return (
    <div>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/searcher" element={<Searcher />} /> 
      </Routes>
    </div>
  );
}

export default App;
