import React, { useState, useEffect } from 'react';
import './Searcher.css';

function Searcher() {
  // State para armazenar os resultados da pesquisa
  const [searchResults, setSearchResults] = useState([]);

  useEffect(() => {
    // Substitua 'http://localhost:8080/search' pelo endpoint correto da sua API
    fetch('http://localhost:8080/search?query=suaConsulta')
      .then(response => response.json())
      .then(data => {
        // Assumindo que 'data' é um array de resultados semelhante ao mockResults
        setSearchResults(data);
      })
      .catch(error => console.error('Erro ao buscar dados:', error));
  }, []); // O array vazio [] como segundo argumento significa que este useEffect será executado apenas uma vez, quando o componente for montado

  return (
    <div className="searcher-container">
      <h2>Resultados da Pesquisa: {searchResults.length} encontrados</h2>
      {searchResults.map(result => (
        <div key={result.id} className="result-item">
          <h3>{result.title}</h3>
          <p>{result.description}</p>
        </div>
      ))}
    </div>
  );
}

export default Searcher;
