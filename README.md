# 🛠️ Maintenance — Under Maintenance Web Page

Este proyecto muestra una página de **"Under Maintenance"** utilizando **Golang + Gin**.  
Permite activar o desactivar el modo mantenimiento mediante una **variable de entorno** o un **archivo flag**.

---

## 🚀 Cómo ejecutar

### 1️⃣ Inicializa el módulo (solo si aún no lo hiciste)

```bash
go mod init tu-proyecto
go get github.com/gin-gonic/gin@v1.10.0
```
### 2️⃣ Ejecuta el servidor

```bash
go run main.go
```
## 🔧 Activar / Desactivar mantenimiento

### 🟣 Opción A — Variable de entorno
####🔹 Activar con variable en línea
```bash
MAINTENANCE=1 go run main.go
```
#### 🔹 Activar exportando la variable (Linux/macOS)
```bash
export MAINTENANCE=1
go run main.go
```

### 🟢 Opción B — Archivo flag

#### 🔹 Activar
```bash
touch maintenance.enabled
```

#### 🔹 Desactivar
```bash
rm -f maintenance.enabled
```

