# 📰 API de Noticias con Go

Este proyecto es una API en Go que consume la API de noticias y permite filtrar por país, categoría e idioma. La documentación está generada con Swagger.

## 🚀 Características
- Obtiene fuentes de noticias de la API externa.
- Filtrado por país, categoría e idioma.
- Documentación con Swagger.
- Uso de variables de entorno para la API Key.

## 🛠️ Tecnologías
- **Go** (Gin Framework)
- **Swagger** para documentación
- **dotenv** para variables de entorno

## 📌 Instalación y ejecución
### 1️⃣ Clona el repositorio:
```bash
git clone https://github.com/CamiloHernandez0910/News-Api.git
cd News-Api
```

### 2️⃣ Crea un archivo `.env` y agrega la API Key:
- Obten la api key GRATIS aquí: [NewsApi](https://newsapi.org/)
```env
API_KEY=tu_api_key_aqui
```

### 3️⃣ Instala las dependencias:
```bash
go mod tidy
```

### 4️⃣ Genera la documentación Swagger:
```bash
swag init
```

### 5️⃣ Ejecuta el servidor:
```bash
go run main.go
```

## 📖 Uso de la API
### 📝 Endpoints principales

#### 🔍 Obtener fuentes de noticias filtradas
```http
GET /api/noticias?sources={source}&country={country}&category={category}&language={language}
```
- **Parámetros:**
  - `country`: Código del país (ej. `fr` para Francia).
  - `category`: Categoría de noticias (ej. `business`, `sports`).
  - `language`: Categoría de noticias (ej. `es`, `en`, `fr`).

### 📑 Documentación Swagger
Accede a la documentación y prueba la api en:
```
http://localhost:8080/swagger/index.html
```

## 👨‍💻 Contribuciones
¡Las contribuciones son bienvenidas! Por favor, sigue estos pasos:
1. Haz un **fork** del repositorio
2. Crea una nueva rama (`git checkout -b feature-nueva`)
3. Realiza los cambios y haz **commit** (`git commit -m 'Agrega nueva funcionalidad'`)
4. Haz **push** a tu rama (`git push origin feature-nueva`)
5. Abre un **Pull Request**

---
**Desarrollado con ❤️ por [HERTUQ](https://github.com/CamiloHernandez0910)** 🚀

