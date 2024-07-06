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

El desarrollo del ejercicio se puede ver reflejado en cada commit, ya que he usado estos a modo de histórico para poder ver la evolución del desarrollo. La forma en la que he procedido para la optimización del código, ha sido solucionando primero los errores más generales y más obvios para limpiar el código, para después centrarme en 

### For changed to range

La primera modifcación del código ha sido cambiar el `for` indexado a un `for` usando `range` de el `array` items. De esta manera, se puede acceder direcamente al obejto sin tener que navegar los indices del `array`, haciendolo más optimo y dejando el código más limpio.

### item.Qualit issue solved

Este `commit` resuelve la redundancia de la línea 48, en la que `item.Quality` se resta a el mismo. Esto se puede solucionar igualando la variable a 0, ya que un número restado por el mismo siempre dará 0.

### Idiomatic sums and minus improved and two ifs mixed as one

Aquí se resuelve otra redundancia en el incremento y decremento del valor de `item.Quality` y `item.SellIn`. Ya que, anteriormente el valor se restaba o sumaba así: `item.Quality = item.Quality + 1`. Siendo en programación (y especialmente en GO) más optimo hacerlo así: `item.Quality++`.

### Code modulerized into small methods, only the increase and decrease quality parts

En esta actualización del código, he convertido en función partes 