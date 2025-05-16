# API Gerenciador de Livros

Esta API permite criar, listar, atualizar e deletar livros.

---

## Endpoints

### Listar todos os livros

- **URL:** `/books`  
- **Método:** `GET`  
- **Resposta:**
  ```json
  [
    {
      "id": 1,
      "title": "Título do livro",
      "author": "Autor do livro"
    },
    ...
  ]
  ```

- **Exemplo curl:**
  ```bash
  curl -X GET http://localhost:8080/books
  ```

---

### Criar um novo livro

- **URL:** `/books`  
- **Método:** `POST`  
- **Corpo (JSON):**
  ```json
  {
    "title": "Nome do livro",
    "author": "Nome do autor"
  }
  ```
- **Resposta:** O livro criado

- **Exemplo curl:**
  ```bash
  curl -X POST http://localhost:8080/books \
    -H "Content-Type: application/json" \
    -d '{"title":"O Alquimista", "author":"Paulo Coelho"}'
  ```

---

### Atualizar um livro

- **URL:** `/books/:id`  
- **Método:** `PUT`  
- **Parâmetros da URL:**  
  `id` — ID do livro a ser atualizado

- **Corpo (JSON):**
  ```json
  {
    "title": "Novo título",
    "author": "Novo autor"
  }
  ```
- **Resposta:** O livro atualizado

- **Exemplo curl:**
  ```bash
  curl -X PUT http://localhost:8080/books/1 \
    -H "Content-Type: application/json" \
    -d '{"title":"Novo Título", "author":"Novo Autor"}'
  ```

---

### Deletar um livro

- **URL:** `/books/:id`  
- **Método:** `DELETE`  
- **Parâmetros da URL:**  
  `id` — ID do livro a ser deletado

- **Resposta:** Mensagem de confirmação

- **Exemplo curl:**
  ```bash
  curl -X DELETE http://localhost:8080/books/1
  ```