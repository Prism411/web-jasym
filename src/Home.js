import React, { useState } from 'react';
import './Home.css';
import Wave from './gui/items/wave.svg';

function Home() {
  // State para armazenar o valor do input
  const [searchTerm, setSearchTerm] = useState('');

  // Handler para atualizar o state com o valor do input
  const handleSearchChange = (event) => {
    setSearchTerm(event.target.value);
  };

  return (
    <div className="home-container">
      <div className="home-header">
        <h1 className="home-title">JAWYN</h1>
      </div>
      <p className="home-welcome">Bem vindo ao JAWYN searcher!</p>
      <input
        type="text"
        className="search-input"
        placeholder="Digite aqui para pesquisar..."
        value={searchTerm}
        onChange={handleSearchChange}
      />
      <button className="search-button">Pesquisar</button>
      <div className="wave-container">
        <img src={Wave} alt="Wave" />
      </div>
    </div>
  );
}

export default Home;
