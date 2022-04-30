## prueba para mercado-libre
esta es la prueba hecha para el proceso de reclutamiento de mercado libre
## Instalacion
para poder abrir los tres niveles necesita esta ubicados en la raiz de cada proyecto las cuales son:
- level-one
- level-two
- level-three
a los tres niveles primero se le tiene que instalar la paqueteria yo uso 
```sh
go mod tidy
```
en el nivel 1 existe un dato archivo llamado data.json el cual se usa para el funcioamiento de dicho ejercicio, este no cuenta con la validacion de nxn de la cadena adn.
en el nivel 2 y 3 si cuenta con la validacion de la cadena nxn , y en estos se tiene que revisar la configuraciones dentro de data/config,yml, en el level 2 no importa la configuracion de la db, en el 3 si es importante.

para correr los tres procesos es necesario hacer uso del comando

```sh
go run main.go
```
## link de las pruebas en la capa gratuita de heroku

| nivel | link |
| ------ | ------ |
|  level-2 | https://mercado-libre-test-2.herokuapp.com/ |
| level-3 | https://mercado-libre-test-juan.herokuapp.com/ |

## en point pedidos en la prueba level 3

| method | link |  end point  | 
| ------ | ------ | ------ |
|  mutant | https://mercado-libre-test-2.herokuapp.com/ | /mutant |
| stats | https://mercado-libre-test-juan.herokuapp.com/ | /stats |




## en point pedidos en la prueba level 2

| method | link |  end point  | 
| ------ | ------ | ------ |
|  mutant | https://mercado-libre-test-2.herokuapp.com/ | /mutant |
