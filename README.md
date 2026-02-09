# maintenance
Under Maintenance web page

#Cómo ejecutar
## 1) Inicializa el módulo si no lo hiciste
go mod init tu-proyecto
go get github.com/gin-gonic/gin@v1.10.0

## 2) Ejecuta
go run main.go
### Abre http://localhost:8080

# Activar/desactivar mantenimiento
## Opción A: Variable de entorno
### Activar
MAINTENANCE=1 go run main.go

### O en Linux/macOS con server ya corriendo (usando systemd o export por línea de comando antes de lanzar)
export MAINTENANCE=1
go run main.go

## Opción B: Archivo "flag"
### Activar
touch maintenance.enabled

### Desactivar
rm -f maintenance.enabled

