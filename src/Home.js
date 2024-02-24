import React from 'react';
import './Home.css';
// Importa o SVG
import Wave from './gui/items/wave.svg';

function Home() {
  return (
    <div className="home-container">
      <div className="home-header">
        <h1 className="home-title">Main Page</h1>
      </div>
      <p className="home-welcome">Welcome to the main page! Explore our features by searching below.</p>
      <button className="search-button">Search</button>
      {/* Adiciona o SVG ao final do conteúdo */}
      <div className="wave-container">
        <img src={Wave} alt="Wave" />
      </div>
    </div>
  );
}

export default Home;
