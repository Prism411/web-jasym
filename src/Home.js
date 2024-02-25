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
  const [errorMessage, setErrorMessage] = useState(''); // Estado para armazenar a mensagem de erro

  const navigate = useNavigate();

  const handleSearchChange = (event) => {
    setSearchTerm(event.target.value);
  };

  const toggleMenu = () => {
    setIsMenuVisible(!isMenuVisible);
  };

  const handleLogin = () => {
    fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ login, password }),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Usuário ou senha inválidos');
        }
        return response.json();
    })
    .then(data => {
        console.log('Success:', data);
        setIsLoggedIn(true);
        setErrorMessage(''); // Limpa a mensagem de erro em caso de sucesso
        // Certifique-se de que esta linha está corretamente atualizando o estado `login`
        // setLogin(login); // Isso parece redundante ou incorreto se 'login' já é o estado atual.
    })
    .catch((error) => {
        console.error('Error:', error);
        setIsLoggedIn(false);
        setErrorMessage(error.message);
    });
};

  const handleRegister = () => {
    console.log("Abrir formulário de registro");
  };

  const handleSearch = () => {
    if (!searchTerm.trim()) {
      console.log('Termo de pesquisa vazio, pesquisa não realizada.');
      return;
    }
    fetch(`http://localhost:8080/search?query=${encodeURIComponent(searchTerm)}`, {
      method: 'GET',
    })
    .then(response => response.json())
    .then(data => {
      console.log('Search Results:', data);
      navigate('/searcher', { state: { searchResults: data } });
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  };
  
  // Função para enviar a chave API
  const handleSendApiKey = () => {
    console.log('Login atual:', login); // Verifique o valor de login aqui
    console.log('Enviando chave API:', apiKey);
  
    // Substitua 'http://localhost:8080/api-endpoint' pelo endpoint correto: '/update-api-key'
    fetch('http://localhost:8080/update-api-key', {
      method: 'POST', 
      headers: {
        'Content-Type': 'application/json',
        // Aqui você pode incluir um token ou outras informações de autenticação se necessário
        // 'Authorization': 'Bearer seu_token_aqui',
      },
      // Certifique-se de enviar tanto a apiKey quanto o login (ou outro identificador do usuário) no corpo da requisição
      body: JSON.stringify({ apiKey, login }), // Substitua 'oLoginDoUsuario' pelo estado ou propriedade correta
    })
    .then(response => {
      if (!response.ok) {
        // Caso a resposta não seja OK, lança um erro com o status para tratamento no .catch
        throw new Error(`Falha ao enviar a chave API: ${response.status}`);
      }
      return response.json();
    })
    .then(data => {
      console.log('Chave API enviada com sucesso:', data);
      // Trate a resposta do servidor aqui, por exemplo, exibindo uma mensagem de sucesso
    })
    .catch((error) => {
      console.error('Erro ao enviar chave API:', error);
      // Aqui você pode atualizar o estado da aplicação para informar o usuário sobre o erro
    });
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
      <button className="search-button" onClick={handleSearch}>Pesquisar</button>
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
          {errorMessage && <div className="error-message">{errorMessage}</div>} {/* Exibe a mensagem de erro aqui */}
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
          <button onClick={handleSendApiKey}>Enviar Chave API</button> 
        </div>
      )}
      <div className="wave-container">
        <img src={Wave} alt="Wave" />
      </div>
    </div>
  );
}

export default Home;