# Actividad 03

## Sección crítica e interbloqueos

Para esta actividad se plantea el siguiente escenario, se tiene un método llamado **incContaor** el cual se ejecuta dos ocasiones de manera concurrente, este método se contiene un ciclo que ejecuta 2 ocasiones el incremento del contador, el programa principal requiere conocer el incremento acumulado de todos los hilos e imprimirlo al finalizar la ejecución.

Al ejecutarse de manera concurrente la variable contador necesita ser incrementada de manera ordenada por cada hilo y por cada iteración de ciclo.

Para este problema se hizo utilización de la función de **go** **mutex** dentro del paquete **sync** del lenguaje. La cual provee de funcionalidad para generar interbloqueos al compartir variables en la sección crítica.
