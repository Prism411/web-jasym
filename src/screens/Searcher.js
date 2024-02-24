import React from 'react';
import './Searcher.css';

// Mock de dados para representar resultados de pesquisa
const mockResults = [
  { id: 1, title: "Site 1", description: "Descrição do site 1" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  { id: 2, title: "Site 2", description: "Descrição do site 2" },
  
  // Adicione mais objetos conforme necessário para simular resultados
];

function Searcher() {
  return (
    <div className="searcher-container">
      <h2>Resultados da Pesquisa: {mockResults.length} encontrados</h2>
      {mockResults.map(result => (
        <div key={result.id} className="result-item">
          <h3>{result.title}</h3>
          <p>{result.description}</p>
        </div>
      ))}
    </div>
  );
}

export default Searcher;
