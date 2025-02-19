# 🚀 RecognizeYou

📌 **Projeto de Reconhecimento Facial para IoT**  
📅 **Data de Criação:** _[Definir]_  
👤 **Responsável:** _[Seu Nome]_  

---

## 📌 Objetivo  

📖 **Descrição:**  
Desenvolver uma solução IoT para **reconhecimento facial**, priorizando:  
- Métodos de **alto nível** em C++ para manipulação de hardware.  
- **Eficiência no envio de dados** para o servidor, considerando concorrência e persistência.  
- Implementação de um **algoritmo robusto de controle de filas**.  

---

## ✅ Requisitos Funcionais  

🔹 **RF-01**: Reconhecer rostos humanos, independentemente da identificação inicial.  
🔹 **RF-02**: Implementar um **algoritmo otimizado** para envio de dados ao servidor, lidando com concorrência e eventos de tentativa.  
🔹 **RF-03**: Criar um **sistema de filas** para garantir consistência na comunicação.  

---

## ⚙️ Requisitos Técnicos  

🔹 **Linguagem:** C++  
🔹 **Banco de Dados:** Temporário para armazenamento intermediário dos dados antes do envio.  
🔹 **Paradigma:** Orientação a Objetos (OOP).  
🔹 **Tipagem:** Tipagem forte (forçada, se necessário).  
🔹 **Logs:** Registro de eventos para depuração e auditoria antes do envio ao servidor.  

---

## 📌 Tasks  

⚠ **Nota:** Todas as implementações devem considerar:  
- **Nível de abstração** adequado.  
- **Uso eficiente de memória**.  
- **Testes unitários e de integração** _(incluindo cobertura de código)_.  

### 🔹 **🛠 0: Planejamento e Definições**  

📌 **0.1**: Controle básico de hardware com C++ (ex: acionar LEDs, manipular GPIOs).  
📌 **0.2**: Escolha do banco de dados IoT (_dinâmico e tipado_).  
📌 **0.3**: Definir **modelos de classes** e estrutura de dados.  
📌 **0.4**: Especificar métodos e comportamentos dos eventos.  
📌 **0.5**: Planejar **tratamento de erros** e cenários inesperados.  
📌 **0.6**: Criar estrutura base e padrões de controle.  

---

### 🔹 **🚀 1: Desenvolvimento do Algoritmo**  

#### 🏗 **1.1: Estruturação do Serviço**  
✅ Criar um **sistema mínimo funcional** capaz de enviar requisições HTTP ao servidor.  
✅ Configuração inicial do **design pattern** (ex: `services/models`).  

#### 🗄 **1.2: Banco de Dados e Persistência**  
✅ Implementação da **conexão** com o banco de dados.  
✅ Desenvolvimento de um **CRUD eficiente** para persistência temporária.  
✅ Implementação de **bulk methods** para otimizar gravações massivas.  

#### ⚡ **1.3: Middleware e Concorrência**  
✅ Criar **middlewares** para manipulação de alta demanda (**lazy loading**).  

#### 🔧 **1.4: Serviços e Configurações**  

✅ Criar **sistema de usuários** para controle de acesso.  
✅ Desenvolver **middleware de autenticação**.  
✅ Criar serviços de **cadastro e manipulação de usuários** no banco.  
✅ Implementar **autenticação e logging** no servidor.  

##### 📡 **1.5: Comunicação com o Servidor**  
✅ Criar serviços para sincronização e envio de dados:  
- **GetConfigs** (_Obtenção de configurações do servidor_).  
- **SendLogConfig** (_Envio de logs e status do dispositivo_).  
- **InitializeIOT** (_Inicialização do dispositivo IoT_).  

---

## ⏳ Dependências  

📌 **Banco de Dados:** _[Definir: SQLite, PostgreSQL, ou outra solução embedded]_  
📌 **Libraries:**  
- _[Listar dependências como OpenCV, Boost, Poco, etc.]_  
- _[Frameworks de comunicação como MQTT, HTTP client]_  

---

✍ **Última Atualização:** _[Definir Data]_  
📂 **Versão do Documento:** `v1.0`  

---
