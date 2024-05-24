-- Cria Banco de dados
CREATE DATABASE goChat;

USE goChat;

-- Criando os esquemas
CREATE SCHEMA messages;

GO

-- Criando tabelas no esquema dbo
CREATE TABLE dbo.usuario (
    id INT PRIMARY KEY IDENTITY(1,1),
    email VARCHAR(100) UNIQUE NOT NULL,
    senha VARCHAR(255) NOT NULL,
    nome VARCHAR(30),
    cargo VARCHAR(30),
    fotoPath VARCHAR(255),
    recado VARCHAR(30),
    eAcessoUsuario INT NOT NULL,
    eStatusUsuario INT,
    delet DATETIME
);

CREATE TABLE dbo.config_geral (
    id INT PRIMARY KEY IDENTITY(1,1),
    chave VARCHAR(30) NOT NULL,
    valor VARCHAR(255) NOT NULL
);


CREATE TABLE dbo.conexao (
    id INT PRIMARY KEY IDENTITY(1,1),
    idUsuario INT NOT NULL,
	tabelaMensagem VARCHAR(51) NOT NULL,
    eTipoConexao INT NOT NULL,
    delet DATETIME,
    FOREIGN KEY (idUsuario) REFERENCES dbo.usuario(id),
);

--------------------------
-- CRIAÇÃO DE TABELA MENSAGEM O OBTENÇÃO DE SEQUENCIA PARA TABELAS DE MENSAGENS

-- Criação da tabela SequenciaTabela
CREATE TABLE dbo.SequenciaTabela (
    Valor INT NOT NULL
);

-- Inserção do primeiro valor na tabela SequenciaTabela
INSERT INTO dbo.SequenciaTabela (Valor) VALUES (0);

-- Criação da stored procedure para criação de conexão
CREATE PROCEDURE criarConexao (
    @idUsuario INT,  
    @tipoConexao INT,
    @idUsuario2 INT,
    @Erro INT OUTPUT
)
AS
BEGIN	
    DECLARE @ProximoValor INT;
    DECLARE @ProximoSequencial VARCHAR(50);
    DECLARE @SQL NVARCHAR(MAX);
    DECLARE @NomeTabela VARCHAR(51);

    BEGIN TRY
        -- Início de uma transação para garantir a atomicidade das operações
        BEGIN TRANSACTION;

        -- Obtenção do último valor gerado
        SELECT @ProximoValor = Valor FROM dbo.SequenciaTabela;

        -- Incremento do último valor e geração do próximo sequencial
        SET @ProximoValor = @ProximoValor + 1;
        SET @ProximoSequencial = 'M' + CAST(@ProximoValor AS NVARCHAR(50));

        -- Atualização do último valor na tabela de controle
        UPDATE dbo.SequenciaTabela SET Valor = @ProximoValor;

        -- Definir o nome da tabela usando o próximo valor sequencial
        SET @NomeTabela = @ProximoSequencial;

        -- Criação dinâmica da tabela de mensagens dos usuários
        SET @SQL = 'CREATE TABLE messages.' + QUOTENAME(@NomeTabela) + ' (
        id INT PRIMARY KEY,
        idConexao INT,
        texto VARCHAR(MAX),
        status_msg INT,
        data_envio DATETIME,
        delet DATETIME,
        FOREIGN KEY (idConexao) REFERENCES dbo.conexao(id))';

        -- Execução da consulta dinâmica
        EXEC sp_executesql @SQL;

        -- Inserção dos detalhes da conexão na tabela de controle
        INSERT INTO dbo.conexao (
            idUsuario,
            tabelaMensagem, 
            eTipoConexao
        )
        VALUES (
            @idUsuario, 
            @NomeTabela, 
            @tipoConexao
        );

        -- Verificação e inserção condicional do segundo usuário
        IF @idUsuario2 IS NULL OR @idUsuario2 <> 0
        BEGIN
            INSERT INTO dbo.conexao (
                idUsuario,
                tabelaMensagem, 
                eTipoConexao
            )
            VALUES (
                @idUsuario2, 
                @NomeTabela, 
                @tipoConexao
            );
        END

        -- Comitar a transação se tudo ocorrer bem
        COMMIT TRANSACTION;

        -- Retorno 0 indicando sucesso
        SET @Erro = 0;
    END TRY
    BEGIN CATCH
        -- Rollback da transação em caso de erro
        ROLLBACK TRANSACTION;

        -- Retorno 1 indicando erro
        SET @Erro = 1;
    END CATCH;
END;
