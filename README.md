# GO Starter

- Run :

```shell
go run texttest_fixture.go [<number-of-days>; default: 2]
```

- Run tests :

```shell
go test ./...
```

- Run tests and coverage :

```shell
go test ./... -coverprofile=coverage.out

go tool cover -html=coverage.out
```

# Resumen del desarrollo

El desarrollo del ejercicio se puede ver reflejado en cada commit, ya que he utilizado estos a modo de histórico para poder observar la evolución del desarrollo. La forma en la que he procedido para la optimización del código ha sido solucionando primero los errores más generales y obvios para limpiar el código, y después centrarme en otros detalles más generales de organización del mismo.

### For changed to range

La primera modificación del código ha sido cambiar el `for` indexado a un `for` usando `range` del array `items`. De esta manera, se puede acceder directamente al objeto sin tener que navegar los índices del array, haciéndolo más óptimo y dejando el código más limpio.

### item.Qualit issue solved

Este commit resuelve la redundancia de la línea 48, en la que `item.Quality` se restaba a sí mismo. Esto se puede solucionar igualando la variable a 0, ya que un número restado por el mismo siempre dará 0.

### Idiomatic sums and minus improved and two ifs mixed as one

Aquí se resuelve otra redundancia en el incremento y decremento del valor de `item.Quality` y `item.SellIn`. Anteriormente, el valor se restaba o sumaba así: `item.Quality = item.Quality + 1`. Siendo en programación (y especialmente en Go) más óptimo hacerlo así: `item.Quality++`.

### Code modulerized into small methods, only the increase and decrease quality parts

En esta actualización del código, he convertido en función partes que se repiten como el incremento y el decremento de la variable `Quality` del objeto `item`. De esta forma, el código queda más limpio y se puede modificar mejor.

### Modularized to one method, but there is a problem with Backstage items

Esta modificación del código es la que lo ha dejado más limpio. En esta se ha juntado toda la lógica de comprobación de nombres del objeto, ya que esta podía ser unida en una sola función que se llama varias veces, una al principio y la otra cuando `SellIn` es menor que 0. En esta actualización ha habido un problema con las variables de tipo `Backstage...`, ya que no se tuvo en cuenta la gestión de estas para cuando el `SellIn` es menor que 0.

### Modularized method works, needs some tests to be implemented

En este commit se ha solucionado el problema del anterior, creando una variable de tipo `bool` llamada `isMinusSellIn`, que será `true` en el caso de que `SellIn` sea menor que 0, y `false` si es mayor que 0. Esta comprobación se da para no igualar a 0 variables innecesarias. También se ha refactorizado la comprobación `if item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert"`.

### Some small refactoring of the code

Aquí se ha hecho una pequeña mejora al código, cambiando el nombre de la variable `isMinusSellIn` a `isSellInNegative` para hacerlo más lógico. También se ha eliminado el `item.SellIn < 6` de la comprobación, ya que era redundante e innecesario.

### tests updated

Este es el último commit relevante para el desarrollo del código, donde se han implementado los códigos de test correctos para comprobar si la aplicación funciona correctamente. Esta parte se podría mejorar.