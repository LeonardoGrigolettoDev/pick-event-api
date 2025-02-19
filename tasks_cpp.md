# RecognizeYou

## 📌 Objetivo
Entender e utilizar, métodos de alto nível em desenvolvimento IOT para reconhecimento de faces em níveis médios, mas principalmente com o objetivo de controlar e manusear o envio de dados para o servidor.

## ✅ Requisitos Funcionais
1. Reconhecer rostos humanos, mas sem a necessidade de quem é essa identificação inicialmente.
2. Algoritmo otimizado para o envio dados ao servidor WEB, sabendo lidar com concorrência e insistência (eventos).
3. Algoritmo de controle de filas.

## ⚙️ Requisitos Técnicos
- C++.
- Banco de dados dinâmico para o armazenamento temporário até o envio de dados ao servidor WEB.
- Orientado a objetos (classes).
- Tipagem forte (se necessário, forçada).
- Controle de logs (também fará parte dos dados temporários, até à posse do servidor).

## 📌 Tasks
(Em todas as tasks deveremos pensar no nível de abstração que o algoritmo está adotando (incluindo construção de classes), pensar em necessidades padrões de cada função e memória, será essencial)
(Testes unitários e integração (é possível ver coverage?) deverão ser implementados de acordo com a abstração micro/macro construída, deverá ser implementado antes e depois (se necessário) de cada implementação)
- 0: Entendimentos e definições:
    - 0.1: Entender como controlar funcionalidades básicas do hardware com C++ (como por exemplo: energizar pinos - acender leds - e mais para frente visualizar a câmera com controle de eventos e comandos - visualização em tempo real das câmeras).
    - 0.1: Definir qual será o banco de dados interno para o IOT, visando a dinamicidade e tipagem do algoritmo. 
    - 0.2: Definir modelos e classes da regra de negócio (assim como reflexos às tabelas dinâmicas).
    - 0.3: Definir métodos de classes.
    - 0.4: Definir comportamento de eventos.
    - 0.5: Tratamento de erros e pensar em possíveis retorno de resultados indesejados.
    - 0.6: Definir estruturas e controle padrão para o algoritmo.
- 1: Criação do algoritmo:
    - 1.1: Construir uma estrutura básica do funcionamento do serviço, podendo ser capaz de mandar uma requisição HTTP para o servidor (inicialização da construção de diretórios como services/models - dependendo do design pattern decidido).
    - 1.2: Construir modelo de conexão com o banco.
    - 1.3: Construir modelos e serviços de manuseio padrão com o banco de dados (CRUD e bulk methods).
    - 1.4: Construir modelos e serviços de middleware para a necessidade de utilização de lazy loading (concorrência, alta demanda, etc).
    - 1.5: Construir modelos, serviços e rotas (locais) de configurações:
        - 1.5.1: Necessitar de um usuário para manipulação de dados (classe de usuário (com chamadas a métodos), middleware de autenticação).
        - 1.5.2: Construir serviços de manipulação e armazenamento de dados no banco para o cadastro de usuários.
        - 1.5.3: Incluir nos serviços de middlewares funções de autenticação e logging ao servidor.
        - 1.5.4: Construir serviços e modelos de configurações do hardware.
        - 1.5.5: Construir serviços e modelos de comunicação com o servidor (absorvendo os serviços do passo 1.5.2).
            - 1.5.5.6: GetConfigs;
            - 1.5.5.7: SendLogConfig;
            - 1.5.5.8: InitializeIOT;


## ⏳ Dependências
- (Banco de dados (não decidido ainda))
- Libraries
