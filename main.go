package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

// isMaintenanceEnabled devuelve true si está activa la página de mantenimiento.
// Se puede activar de dos formas:
// 1) Variable de entorno: MAINTENANCE=1
// 2) Presencia del archivo "maintenance.enabled" en el directorio del binario.
func isMaintenanceEnabled() bool {
    if os.Getenv("MAINTENANCE") == "1" {
        return true
    }
    if _, err := os.Stat("maintenance.enabled"); err == nil {
        return true
    }
    return false
}

// MaintenanceMiddleware intercepta todas las peticiones y, si hay mantenimiento,
// responde con la plantilla de mantenimiento para rutas HTML y un JSON básico para APIs.
func MaintenanceMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        if isMaintenanceEnabled() {
            accept := c.GetHeader("Accept")
            // Si el cliente acepta HTML, devuelve la página bonita
            if accept == "" || (accept != "" && (contains(accept, "text/html") || contains(accept, "*/*"))) {
                c.HTML(http.StatusServiceUnavailable, "maintenance.html", gin.H{
                    "title": "Under Maintenance",
                })
            } else {
                // Para APIs u otros clientes
                c.JSON(http.StatusServiceUnavailable, gin.H{
                    "status":  "maintenance",
                    "message": "Service temporarily unavailable due to maintenance.",
                })
            }
            c.Abort()
            return
        }
        c.Next()
    }
}

func contains(haystack, needle string) bool {
    return len(haystack) >= len(needle) && (func() bool {
        // chequeo simple sin strings.Contains para evitar imports extra
        for i := 0; i+len(needle) <= len(haystack); i++ {
            if haystack[i:i+len(needle)] == needle {
                return true
            }
        }
        return false
    })()
}

func main() {
    // Usa modo release para producción
    gin.SetMode(gin.ReleaseMode)

    r := gin.Default()

    // Carga de plantillas
    r.LoadHTMLGlob("templates/*")

    // Archivos estáticos (CSS/JS/imagenes)
    r.Static("/static", "./static")

    // Middleware de mantenimiento delante de todo
    r.Use(MaintenanceMiddleware())

    // Rutas de tu web estática
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "Mi Web Estática",
        })
    })

    // Ejemplo de ruta adicional
    r.GET("/about", func(c *gin.Context) {
        c.String(http.StatusOK, "Acerca de: esta es una página de ejemplo.")
    })

    // Puerto configurable por env (PORT) o 8080 por defecto
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Servidor iniciado en http://localhost:%s (mantenimiento=%v)\n", port, isMaintenanceEnabled())
    if err := r.Run(":" + port); err != nil {
        log.Fatal(err)
    }
}
