import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom'; // Importando useNavigate
import './Home.css';
import Wave from './gui/items/wave.svg';
import JawynLogo from './gui/icons/jawyn.png';


function Home() {
  const [searchTerm, setSearchTerm] = useState('');
  const [isMenuVisible, setIsMenuVisible] = useState(false);
  const [login, setLogin] = useState('');
  const [password, setPassword] = useState('');
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [apiKey, setApiKey] = useState('');

  const navigate = useNavigate(); // Instância de useNavigate

  const handleSearchChange = (event) => {
    setSearchTerm(event.target.value);
  };

  const toggleMenu = () => {
    setIsMenuVisible(!isMenuVisible);
  };

  const handleLogin = () => {
    console.log({ login, password });
    setIsLoggedIn(true);
  };

  const handleRegister = () => {
    console.log("Abrir formulário de registro");
  };

  const handleSearch = () => {
    navigate('/searcher'); // Navega para o componente Searcher
  };
  return (
    <div className="home-container">
      <div className="home-header">
        <h1 className="home-title">JAWYN</h1>
      </div>
      <div className="logo-container">
        <img src={JawynLogo} alt="JAWYN Logo" />
      </div>
      <button className="menu-button" onClick={toggleMenu}>Menu</button> 
      <input
        type="text"
        className="search-input"
        placeholder="Digite aqui para pesquisar..."
        value={searchTerm}
        onChange={handleSearchChange}
      />
      <button className="search-button"onClick={handleSearch}>Pesquisar</button> {/* Botão de pesquisa agora com evento onClick */}
      {isMenuVisible && !isLoggedIn && (
        <div className="side-menu">
          <input
            type="text"
            className="text_input"
            placeholder="Login"
            value={login}
            onChange={(e) => setLogin(e.target.value)}
          />
          <input
            type="password"
            className="password_input"
            placeholder="Senha"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <button onClick={handleLogin}>Entrar</button>
          <button onClick={handleRegister} className="register-button">Registrar</button>
        </div>
      )}
      {isMenuVisible && isLoggedIn && (
        <div className="side-menu">
          <input
            type="text"
            className="api-key-input"
            placeholder="Insira sua chave de API aqui"
            value={apiKey}
            onChange={(e) => setApiKey(e.target.value)}
          />
        </div>
      )}
      <div className="wave-container">
        <img src={Wave} alt="Wave" />
      </div>
    </div>
  );
} 

export default Home;
