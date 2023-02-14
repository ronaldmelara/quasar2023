<!-- TOC -->

- [1. **Operación Fuego de Quasar**](#1-operación-fuego-de-quasar)
    - [1.1. **Desafío**](#11-desafío)
    - [1.2. **Nivel 1**](#12-nivel-1)
    - [1.3. **Nivel 2**](#13-nivel-2)
    - [1.4. **Nivel 2**](#14-nivel-2)
- [2. **Solución**](#2-solución)
    - [2.1 **Localización**](#21-localización)
    - [2.2 **Mensaje**](#22-mensaje)
- [3. **Documnetación Referencial**](#3-documnetación-referencial)

<!-- /TOC -->

# 1. **Operación Fuego de Quasar**

Han Solo ha sido recientemente nombrado General de la Alianza Rebelde y busca dar un golpe contra el Imperio Galáctico para reavivar la llama de la resistencia.
El servicio de inteligencia rebelde ha detectado un llamado de auxilio de una nave portacarga imperial a la deriva en un campo de asteroides. El manifiesto de la nave es ultra clasificado, pero se rumorea que trasporta raciones y armamento para una legión entera.

![Imagen ilustrativa del enunciado del problema](/assets/img/img_enun_problem01.png)

## 1.1. **Desafío**

Como jefe de comunicaciones rebelde, tu misión es crear un programa en Golang que **retorne la fuente y contenido del mensaje de auxilio**. Para esto, cuentas con tres satélites que te permitirán triangular la posición, ¡pero cuidado! el mensaje puede no llegar completo a cada satélite debido al campo de asteroides frente a la nave.

**Posición de los satélites actualmente en servicio**

- Kenobi:    [-500,-200]
- Skywalter: [100, -100]
- Sato:      [500,  100]

## 1.2. **Nivel 1**

Crea un programa con las siguientes firmas:

*// input: distancia al emisor tal cual se recibe en cada satélite.*

*// output: las coordenadas 'x' e 'y' del emisor del mensaje.*

**func GetLocation(distances ...float32) (x, y float32)**

*// input: el mensaje tal cual se recibe en cada satélite*
*// output: el mensaje tal cual lo genera el emisor del mensaje*

**func GetMenssage(messages ...[]string) (msg string)**

Consideraciones:

* La unidad de distancia en los parámetros de *GetLocation* es la misma que la que se utilizan para indicar la posicion de cada satélite.
* El mensaje recibido en cada satélite se recibe en forma de arreglo de string.
* Cuando una palabra del mensaje no puede ser determinada, se reemplaza por un string en blanco en el array.
  * Ejemplo: ["este", "es", "", "mensaje"]
* Considerar que existe un desfasaje (a determinar) en el mensaje que se recibe en cada satélite
  * Ejemplo:
    * Kenobi: ["", "este", "es", "un", "mensaje"]
    * Skywalker: ["este", "", "un", "mensaje"]
    * Sato: ["", "", "es", "", "mensaje"]

## 1.3. **Nivel 2**

Crear una API REST, hostear esa API en un cloud computing libre (Google APP Engine, Amazon AWS, etc), crear el servicio /topsecret/ en donde se puede obtener la ubicación de la nave y el mensaje que emite.

El servicio recibirá la información de la nave a través de un HTTP POST con un payload con el
siguiente formato:

POST → /topsecret/

```yaml

{
   "satellites": [
      {
        “name”: "kenobi",
        “distance”: 100.0,
        “message”: ["este", "", "", "mensaje", ""]
      },
      {
        “name”: "skywalker",
        “distance”: 115.5
        “message”: ["", "es", "", "", "secreto"]
      },
      {
        “name”: "sato",
        “distance”: 142.7
        “message”: ["este", "", "un", "", ""]
      }
   ]
}
```

La respuesta, por otro lado, deberá tener la siguiente forma:
RESPONSE CODE: 200

```yaml
{
   "position": {
     "x": -100.0,
     "y": 75.5
   },
   "message": "este es un mensaje secreto"
}
```

En caso que no se pueda determinar la posición o el mensaje, retorna:
RESPONSE CODE: 404

## 1.4. **Nivel 2**

Considerar que el mensaje ahora debe poder recibirse en diferentes POST al nuevo servicio
/topsecret_split/ , respetando la misma firma que antes. Por ejemplo:
POST → /topsecret_split/{satellite_name}

```yaml
{
   "distance": 100.0,
   "message": ["este", "", "", "mensaje", ""]
}
```

Crear un nuevo servicio /topsecret_split/ que acepte POST y GET. En el GET la
respuesta deberá i ndicar l a posición y el mensaje en caso que sea posible determinarlo y tener
la misma estructura del ejemplo del Nivel 2. Caso contrario, deberá responder un mensaje de
error indicando que no hay suficiente información.

# 2. **Solución**

## 2.1 **Localización**

Para resolver el punto de GetLocation es necesario poder entender y aplicar el concepto de **"Trilateración"**. La definición nos dice que es una **técnica geométrica para determinar la posición de un objecto conociendo su distancia a tres puntos de referencia**.
Para poder calcular la posición del objecto en estudio se requiere tener las posicion X e Y del Punto 1 (P1), Punto 2 (P2) y Punto 3(P3) y que cada punto es el centro de una circunferencia. Además, se necesita el radio o distancia desde el centro de cada circunferencia hacia el borde de esta, el cual es el lugar donde se encuentra nuestro objecto a estudiar. Ahora, hay que imaginar que estas 3 circunferencias convergen en un punto, el cual indica el lugar en el espacio donde se encuentra nuestro objecto a estudiar.
![Imagen circunferencias](/assets/img/01_sol.jpeg)

![Imagen circunferencias](/assets/img/02_sol.jpeg)

Entonces tenemos la siguiente información:


| Satélite | Posición X | Posición Y | Distancia |
| ----------- | ------------- | ------------- | ----------- |
| Kenobi    | -500        | -200        | 100       |
| Skywalker | 100         | -200        | 115.5     |
| Sato      | 500         | 100         | 142.7     |

La teoría analítica para cada uno de los puntos en cuestión nos indica que debemos de usar la siguiente ecuación:
![Ecuación](/assets/img/03_sol.jpeg)

Sacamos esta ecuación para cada circunferencia y punto dado:
![Ecuación por circunferencia](/assets/img/04_sol.jpeg)

Ahora igualamos la ecuación del P1 en P2, y agrupamos los resultados nombrandolos como A, B y C
![Ecuación por circunferencia](/assets/img/05_sol.jpeg)

Ahora igualamos la ecuación de P2 en P3 y agrupamos los resultados nombrandolos como E, F y G
![Ecuación por circunferencia](/assets/img/06_sol.jpeg)

De lo anterior tenemos como resultado 2 ecuaciones:
![Ecuación por circunferencia](/assets/img/07_sol.jpeg)

Finalmente resolvemos por Determinantes o regla de Cramer para obtener las ecuaciones que nos permitirán conocer el punto X e Y donde se intecepta de la nave enemiga con los 3 satélites.
![Ecuación por circunferencia](/assets/img/08_sol.jpeg)


## 2.2 **Mensaje**
El problema del mensaje implica hacer un **merge** de los 3 arrays del mensaje que ha recibido cada satélite. Según se indica que debido al defase de la señal puede que algunas palabras no lleguen al satélite pero queda registrado como un input vacío, lo cual hará que en cada satélite existan 3 colecciones del mismo largo, permitiendo asi el merge y posterior obtener los valores unicos que desifrarán el mensaje.

# 3. **Documnetación Referencial**
Youtube
- How does GPS work? https://www.youtube.com/watch?v=FU_pY2sTwTA&t=22s
- Trilateración vs Triangulación https://www.youtube.com/watch?v=WzCXNIDbw7w&t=256s
- Trilateración (PDF file), Autor Mauricio Gende e Ivana Molina from academia.edu
- How Does Your GPS Device Know Where You Are? https://www.youtube.com/watch?v=4fXjc9uibGM
- Intersección de dos circunferencias https://www.youtube.com/watch?v=_BhsWxtGDog
- Golang https://www.golangprograms.com
- Google Cloud (onboarding) https://cloud.google.com/run/docs/quickstarts/build-and-deploy/deploy-go-service?hl=es-419
- Google Cloud CLI https://cloud.google.com/sdk/docs/install?hl=es-419
