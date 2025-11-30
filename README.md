# 🌳 MMP: Map My Project

MMP es una herramienta CLI de línea de comandos rápida y moderna escrita en **Go (Golang)**, inspirada en el comando `tree` de Unix. Permite visualizar la estructura de directorios de cualquier proyecto de forma recursiva, ofreciendo control avanzado sobre la profundidad y el filtrado.

## ✨ Características Principales

* **Velocidad:** Compilada en Go, es extremadamente rápida en el escaneo de grandes sistemas de archivos.
* **Filtro `.gitignore`:** Ignora automáticamente archivos y directorios definidos en el archivo `.gitignore` de la raíz del proyecto.
* **Control Total:** Permite gestionar la profundidad y la visualización de archivos ocultos.

---

## 🚀 Instalación y Uso

### Compilación

Para compilar y crear el ejecutable `mmp` en tu sistema:

```bash
# 1. Asegúrate de estar en el directorio raíz de tu proyecto
go build -o mmp main.go