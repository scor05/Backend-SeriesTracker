# Backend-SeriesTracker 

## Descripción
Este es el repositorio con el código server-side para mi proyecto #1 de la clase de Sistemas y Tecnologías Web: un tracker de series de películas o de cualquier otra índole.

### Instrucciones de ejecución 
Este repositorio está diseñado para ir en conjunto con [este otro repositorio del código client-side hecho con JS, CSS y HTML nativo](https://github.com/scor05/Frontend-SeriesTracker), el cual también dentro de su `README.md` tiene instrucciones para su ejecución.

Para ejecutar el código de este repositorio, se recomienda tener instalado el lenguaje `Go`, el cual permite ejecutar el código de dos maneras: corriendolo localmente con `go run ./src/api` o bien compilando directamente a binario con `go build -o seriesTracker ./src/api` (desde la carpeta raíz obligatoriamente) y luego ejecutando el archivo `./seriesTracker`, el cual automáticamente empezará a manejar las requests en el puerto `:8080`.

También es de notar de que la base de datos guardada en `src/database/sql/series.db` debe de estar en ese mismo path cuando se ejecute el binario, si no retornará un error al no encontrar ese archivo.

El código también es visible en [este servidor](https://joelsiervas.online/24472/SeriesTracker/)

### CORS
CORS, acrónimo de Cross Origin Resource Sharing, es un protocolo de seguridad establecido por los browsers en sí que básicamente bloquea (o en este caso permite) acceso para transferir recursos entre distintos dominios/puertos o demás, únicamente mediante el uso de headers de HTTP normales. La configuración que hice para este proyecto fue de utilizar un middleware (idea que me surgió por mi proyecto de otra clase, donde hicimos casi lo mismo pero con laravel), el cual ejecuta las funciones de `net/http` en vez del mismo `http` de por sí. Para ello, le configuré una función `enableCORS`, que en resumidas cuentas solo edita esos headers de CORS para que tanto el origen (puerto del frontend), los métodos que se pueden utilizar y el content-type no sean bloqueados por el browser. Luego se modifica la clásica línea `http.Listen()` con el puerto para simplemente agregarle como segundo parámetro el `http.HandlerFunc` que retorna `enableCORS` 

---
## Challenges Implementados
Para este proyecto, habían varios retos que se podían implementar para obtener la nota completa, de los cuales decidí implementar los siguientes:


---
## Reflexión de Tecnología Utilizada
Para esta parte del proyecto, decidí utilizar Go como la tecnología escogida para el backend, principalmente para refinar aún más mis habilidades para el lenguaje. Pensé en utilizar php junto con Laravel en un inicio, pero como ese stack lo escogimos para otro proyecto de otra clase, decidí quedarme con el good ol' reliable `Go`.
