import React, { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import './Searcher.css';

function Searcher() {
  const [searchResults, setSearchResults] = useState([]);
  const location = useLocation();

  useEffect(() => {
    // Verifique se há algum estado passado para este componente.
    if (location.state && location.state.searchResults) {
      setSearchResults(location.state.searchResults);
    }
  }, [location]);

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
