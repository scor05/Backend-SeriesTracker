# Backend-SeriesTracker 

## Descripción
Este es el repositorio con el código server-side para mi proyecto #1 de la clase de Sistemas y Tecnologías Web: un tracker de series de películas o de cualquier otra índole.

### Instrucciones de ejecución 
Este repositorio está diseñado para ir en conjunto con [este otro repositorio del código client-side hecho con JS, CSS y HTML nativo](https://github.com/scor05/Frontend-SeriesTracker), el cual también dentro de su `README.md` tiene instrucciones para su ejecución.

Para ejecutar el código de este repositorio, se requiere tener instalado el lenguaje `Go`, el cual permite ejecutar el código de dos maneras: corriendolo localmente con `go run ./src/api` o bien compilando directamente a binario con `go build -o seriesTracker ./src/api` (desde la carpeta raíz obligatoriamente) y luego ejecutando el archivo `./seriesTracker`, el cual automáticamente empezará a manejar las requests en el puerto `:8080`. Es de notar que para que ambos de estos métodos funcionen, el archivo `router.go` debe de estar dentro de la carpeta `src/api`, tal y como está actualmente en el repositorio. También es de notar de que la base de datos guardada en `src/database/sql/series.db` debe de estar en dicho path cuando se ejecute el binario, si no retornará un error al no encontrar ese archivo.

El código también es visible en [este servidor](https://joelsiervas.online/24472/SeriesTracker/)

### CORS
CORS, acrónimo de Cross Origin Resource Sharing, es un protocolo de seguridad establecido por los browsers en sí que básicamente bloquea (o en este caso permite) acceso para transferir recursos entre distintos dominios/puertos o demás, únicamente mediante el uso de headers de HTTP normales. La configuración que hice para este proyecto fue de utilizar un middleware (idea que me surgió por mi proyecto de otra clase, donde hicimos casi lo mismo pero con laravel), el cual ejecuta las funciones de `net/http` en vez del mismo `http` de por sí. Para ello, le configuré una función `enableCORS`, que en resumidas cuentas solo edita esos headers de CORS para que tanto el origen (puerto del frontend), los métodos que se pueden utilizar y el content-type no sean bloqueados por el browser. Luego se modifica la clásica línea `http.Listen()` con el puerto para simplemente agregarle como segundo parámetro el `http.HandlerFunc` que retorna `enableCORS` 

---
## Challenges Implementados
Para este proyecto, habían varios retos que se podían implementar para obtener la nota completa, de los cuales decidí implementar los siguientes:
1. Manejo de códigos correctos de HTTP en la API (como 201 al crear, 204 al eliminar)
2. Búsqueda por nombre con `?q=`

---
## Reflexión de Tecnología Utilizada
Para esta parte del proyecto, decidí utilizar Go como la tecnología escogida para el backend, principalmente para refinar aún más mis habilidades para el lenguaje. Pensé en utilizar php junto con Laravel en un inicio, pero como ese stack lo escogimos para otro proyecto de otra clase, decidí quedarme con el good ol' reliable `Go`.

El haber usado Go como tecnología para el backend terminó siendo una muy buena idea, pues este es es uno de los lenguajes de programación más versátiles y a la vez amigables que hay (hasta el momento). En general tuve una muy buena DX utilizando este lenguaje, y no necesariamente por todo el tooling que tiene, sino que por la modularización en paquetes que Go permite. Digo esto porque desafortunadamente he vivido el tener que manejar proyectos con Gradle, lo cual me dejó un leve PTSD al oír la palabra "paquete", pero Go la verdad me curó un poco el trauma. Simplemente poder editar `package database` o `package routes` me hizo mil veces más fácil la vida al tener que importar las funciones que tenía hechas en dichas carpetas, forzándome también a hacer obvio qué funciones se exportan colocándoles una letra mayúscula al inicio.

Tuve realmente dos problemas que vi realmente al utilizar el lenguaje de Go, el primero de ellos siendo el tener que estar manejando los imports de cada programa manualmente. Sin embargo, esto lo puedo atribuír a una mala configuración de mi entorno de programación que no agrega automáticamente esos imports, pero de todas maneras fue una inconveniencia. Y segundo, lo que más sentí que era un problema, es el error handling. Todas las funciones que llamaba en cualquier punto del código tenían que tener un "if err != nil", lo cual se vuelve bastante inconveniente, especialmente cuando se declaran nuevas variables. Digo esto último porque no se puede reutilizar el mismo error para otra declaración de una función, lo cual llevó a que en ciertas partes del código tuviera que declarar variables tontas como `err2` y `err3` para no repetir `err`.

En resumen, desde que aprendí Go, este se volvió uno de mis lenguajes favoritos, pero al enterarme bien al hacer este proyecto de cómo funcionan sus paquetes y módulos, definitivamente se volvió mi preferencia. Si pudiera utilizarlo para algún otro proyecto en el futuro que sea a aún mayor escala que el mío, definitivamente tomaría la oportunidad para ver código de Go ya perfeccionado y que pueda manejar esos inconvenientes de error handling que me costaron un poco a mí.

---
## Captura del sistema funcionando
<img width="1200" height="700" alt="image" src="https://github.com/user-attachments/assets/5824c2e7-6204-454c-b8a2-01ca8520b76c" />
