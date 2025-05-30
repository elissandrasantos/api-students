# API Students

API RESTful para gerenciamento de estudantes desenvolvida em Go com Echo Framework e GORM.

## 🚀 Tecnologias

- **Go** - Linguagem de programação
- **Echo** - Framework web
- **GORM** - ORM para Go
- **Swagger** - Documentação da API
- **SQLite** - Banco de dados

## 📋 Funcionalidades

- ✅ Criar estudante
- ✅ Listar todos os estudantes
- ✅ Buscar estudante por ID
- ✅ Atualizar dados do estudante
- ✅ Deletar estudante
- ✅ Filtrar por status ativo/inativo
- ✅ Documentação Swagger

## 🛠️ Instalação

1. Clone o repositório
```bash
git clone https://github.com/elissandrasantos/api-students.git
cd api-students
```

2. Instale as dependências
```bash
go mod download
```

3. Gere a documentação Swagger
```bash
swag init
```

## ▶️ Como Executar

```bash
make run
```

Ou diretamente:
```bash
go run main.go
```

A API estará disponível em `http://localhost:8080`

## 📁 Estrutura do Projeto

```
api-students/
├── api/
│   ├── api.go         # Configuração da API e rotas
│   ├── handler.go     # Handlers dos endpoints
│   └── request.go     # Estruturas de request e validação
├── db/
│   └── db.go          # Conexão e operações do banco
├── schemas/
│   └── schemas.go     # Modelos de dados
├── docs/              # Documentação Swagger (gerada)
├── main.go            # Ponto de entrada
├── Makefile          # Comandos make
├── go.mod            # Dependências
└── student.db        # Banco SQLite
```

## 🔌 Endpoints

| Método | Rota | Descrição |
|--------|------|-----------|
| GET | `/students` | Lista todos os estudantes |
| GET | `/students?active=true` | Lista estudantes ativos |
| GET | `/students/{id}` | Busca estudante por ID |
| POST | `/students` | Cria novo estudante |
| PUT | `/students/{id}` | Atualiza estudante |
| DELETE | `/students/{id}` | Remove estudante |

## 📚 Documentação Swagger

Acesse a documentação interativa em:
```
http://localhost:8080/swagger/index.html
```

## 💻 Exemplos de Uso

### Criar estudante
```bash
curl -X POST http://localhost:8080/students \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "cpf": 12345678900,
    "email": "joao@email.com",
    "age": 20,
    "active": true
  }'
```

### Listar estudantes
```bash
curl http://localhost:8080/students
```

### Buscar por ID
```bash
curl http://localhost:8080/students/1
```

### Atualizar estudante
```bash
curl -X PUT http://localhost:8080/students/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva Santos",
    "age": 21
  }'
```

### Deletar estudante
```bash
curl -X DELETE http://localhost:8080/students/1
```

## 📊 Modelo de Dados

### Student
```json
{
  "id": 1,
  "name": "João Silva",
  "cpf": 12345678900,
  "email": "joao@email.com",
  "age": 20,
  "active": true,
  "createdAt": "2024-01-01T10:00:00Z",
  "updatedAt": "2024-01-01T10:00:00Z"
}
```

---

Desenvolvido por [Elissandra Santos](https://github.com/elissandrasantos)