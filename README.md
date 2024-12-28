
# Candidate Management API

Este proyecto proporciona una API para gestionar candidatos, donde podrás obtener, crear y autenticar usuarios. Está basado en **Go** y expone una serie de endpoints que permiten interactuar con la base de datos de candidatos.

## Requisitos

Antes de ejecutar la aplicación, asegúrate de tener instalados los siguientes requisitos:

- **Go 1.18+**
- **Postman** (para probar las rutas y las colecciones)

## Configuración

### 1. Clonar el Repositorio

Clona el repositorio en tu máquina local:

```bash
git clone <URL del repositorio>
cd candidate-management-api
```

### 2. Variables de Entorno

Copia el archivo `.env.example` y renómbralo como `.env` en la raíz del proyecto.

```bash
cp .env.example .env
```

Asegúrate de llenar las variables de entorno necesarias (como las credenciales de la base de datos y las claves secretas para el JWT).

El archivo `.env` debe contener las siguientes variables:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=candidate_db
JWT_SECRET=your_jwt_secret_key
```

### 3. Instalar Dependencias

Si el proyecto tiene dependencias en Go, instala los paquetes necesarios:

```bash
go mod tidy
```


Este comando iniciará un contenedor de PostgreSQL. Asegúrate de tener el archivo `.env` configurado correctamente con las credenciales de la base de datos.

### 4. Ejecutar la Aplicación

Ejecuta la aplicación con el siguiente comando:

```bash
go run main.go
```

Esto iniciará el servidor en **`http://localhost:8080`**.

## Endpoints

Aquí tienes la lista de endpoints disponibles en la API:

### 5. **Generar Bearer Token**

**Método**: `POST`

**URL**: `http://localhost:8080/login/get-token`

**Cuerpo**:
```json
{}
```

**Respuesta**:

- **200 OK**: Se genera un token JWT válido.
  ```json
  {
    "token": "<jwt_token>"
  }
  ```

- **500 Internal Server Error**: Error generando el token.
  ```json
  {
    "error": "Error generating token"
  }
  ```

### 2. **Obtener Todos los Candidatos**

**Método**: `GET`

**URL**: `http://localhost:8080/candidates`

**Cabecera**:
```
Authorization: Bearer <token>
```

**Respuesta**:

- **200 OK**: Lista de candidatos.
  ```json
  [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "gender": "Male",
      "salary_expected": 50000,
      "created_at": "2024-12-27T10:00:00Z"
    }
  ]
  ```

- **500 Internal Server Error**: Error al obtener los candidatos.
  ```json
  {
    "error": "Error fetching candidates"
  }
  ```

### 3. **Crear un Candidato**

**Método**: `POST`

**URL**: `http://localhost:8080/candidates`

**Cabecera**:
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Cuerpo**:
```json
{
  "name": "Jane Smith",
  "email": "jane@example.com",
  "gender": "Female",
  "salary_expected": 60000
}
```

**Respuesta**:

- **201 Created**: El candidato fue creado exitosamente.
  ```json
  {
    "message": "Candidate created successfully"
  }
  ```

- **400 Bad Request**: Datos inválidos en la solicitud.
  ```json
  {
    "error": "Invalid request"
  }
  ```

- **500 Internal Server Error**: Error al crear el candidato.
  ```json
  {
    "error": "Error creating candidate"
  }
  ```

## Pruebas

Puedes realizar pruebas de la API con **Postman**. Te recomendamos importar la colección de Postman desde el siguiente enlace:

[**Candidate Management API - Postman Collection**](<enlace-al-archivo-de-colección>)

### Cómo Importar la Colección en Postman:

1. Abre **Postman**.
2. Haz clic en **Importar**.
3. Selecciona **Archivo** y carga el archivo **`Candidate_Management_API.postman_collection.json`**.
4. Usa las rutas proporcionadas en la colección para probar la API.

## Envío de la Configuración por Correo

Si necesitas enviar la configuración por correo electrónico, asegúrate de adjuntar los siguientes archivos:

1. El archivo `.env` configurado.
2. La colección de Postman (`Candidate_Management_API.postman_collection.json`).

Envíalos a la dirección de correo de tu equipo o cliente, proporcionando las instrucciones adecuadas.

## link de test

1. go test ./test/controllers

## Licencia

Este proyecto está bajo la licencia MIT. Ver el archivo [LICENSE](LICENSE) para más detalles.
