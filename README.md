# Backend-SeriesTracker 

## Descripción
Este es el repositorio con el código server-side para mi proyecto #1 de la clase de Sistemas y Tecnologías Web: un tracker de series de películas o de cualquier otra índole.

### Instrucciones de ejecución 
Este repositorio está diseñado para ir en conjunto con [este otro repositorio del código client-side hecho con JS, CSS y HTML nativo](https://github.com/scor05/Frontend-SeriesTracker), el cual también dentro de su `README.md` tiene instrucciones para su ejecución.

Para ejecutar el código de este repositorio, se recomienda tener instalado el lenguaje `Go`, el cual permite ejecutar el código de dos maneras: corriendolo localmente con `go run .` o bien compilando directamente a binario con `go build -o seriesTracker ./src/api` (desde la carpeta raíz obligatoriamente) y luego ejecutando el archivo `./seriesTracker`, el cual automáticamente empezará a manejar las requests en el puerto `:8080`.

También es de notar de que la base de datos guardada en `src/database/sql/series.db` debe de estar en ese mismo path cuando se ejecute el binario, si no retornará un error al no encontrar ese archivo.

El código también es visible en [este servidor](https://joelsiervas.online/24472/SeriesTracker/)

---
## Challenges Implementados
Para este proyecto, habían varios retos que se podían implementar para obtener la nota completa, de los cuales decidí implementar los siguientes:


---
## Reflexión de Tecnología Utilizada
Para esta parte del proyecto, decidí utilizar Go como la tecnología escogida para el backend, principalmente para refinar aún más mis habilidades para el lenguaje. Pensé en utilizar php junto con Laravel en un inicio, pero como ese stack lo escogimos para otro proyecto de otra clase, decidí quedarme con el good ol' reliable `Go`.
