# 🌳 MMP: Map My Project

**MMP** es una herramienta de línea de comandos (CLI) moderna y ultrarrápida escrita en **Go**, inspirada en el clásico comando `tree`. Está diseñada para ayudarte a visualizar estructuras de directorios complejas con control total sobre lo que deseas ver.

---

## ✨ Características Principales

* 🚀 **Alto Rendimiento:** Gracias a Go, escanea miles de archivos en milisegundos.
* 🙈 **Respeto al `.gitignore`:** Filtra automáticamente archivos y carpetas innecesarias (como `node_modules` o `.git`) basándose en tus reglas de Git.
* 🛠 **Personalización:** Controla la profundidad, archivos ocultos y tipos de visualización mediante flags sencillos.

---

## 🚀 Instalación y Compilación

Para compilar y generar el ejecutable en tu sistema local:

```bash
# 1. Clona el repositorio e ingresa a la carpeta
git clone https://github.com/victor-romero-martinez/mmp.git && cd mmp

# 2. Compila el binario
go build -o mmp main.go

# 3. (Opcional) Muévelo a tu PATH para usarlo en cualquier sitio
mv mmp /usr/local/bin/

```

---

## 🛠 Parámetros y Uso

La sintaxis básica es:

`mmp [flags]`

| Flag | Descripción |
| --- | --- |
| `-L <int>` | **Límite de profundidad**: Define cuántos niveles quieres bajar (0 = sin límite). |
| `-a` | **Mostrar todo**: Incluye archivos y carpetas ocultas (aquellos que empiezan con punto). |
| `-d` | **Solo directorios**: Omite los archivos y muestra únicamente la jerarquía de carpetas. |
| `-v` | **Versión**: Muestra la versión actual de la herramienta. |

### Ejemplos de uso:

**Ver solo 2 niveles de profundidad:**

```bash
mmp -L 2

```

**Ver todas las carpetas del proyecto (incluyendo ocultas) sin archivos:**

```bash
mmp -a -d

```

---
