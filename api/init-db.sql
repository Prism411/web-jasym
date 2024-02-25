    -- Criação da tabela de usuários
    CREATE TABLE IF NOT EXISTS usuarios (
        id SERIAL PRIMARY KEY,
        login VARCHAR(255) UNIQUE NOT NULL,
        senha VARCHAR(255) NOT NULL,
        imagem_perfil VARCHAR(255),
        api_key VARCHAR(255) UNIQUE NOT NULL
    );

    -- Criação da tabela de sites
    CREATE TABLE IF NOT EXISTS sites (
        id SERIAL PRIMARY KEY,
        titulo VARCHAR(255) NOT NULL,
        descricao TEXT,
        link VARCHAR(255) NOT NULL
    );
