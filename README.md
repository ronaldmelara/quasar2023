# **Operación Fuego de Quasar**

Han Solo ha sido recientemente nombrado General de la Alianza Rebelde y busca dar un golpe contra el Imperio Galáctico para reavivar la llama de la resistencia.
El servicio de inteligencia rebelde ha detectado un llamado de auxilio de una nave portacarga imperial a la deriva en un campo de asteroides. El manifiesto de la nave es ultra clasificado, pero se rumorea que trasporta raciones y armamento para una legión entera.

# **Desafío**
Como jefe de comunicaciones rebelde, tu misión es crear un programa en Golang que **retorne la fuente y contenido del mensaje de auxilio**. Para esto, cuentas con tres satélites que te permitirán triangular la posición, ¡pero cuidado! el mensaje puede no llegar completo a cada satélite debido al campo de asteroides frente a la nave.

**Posición de los satélites actualmente en servicio**
- Kenobi:    [-500,-200]
- Skywalter: [100, -100]
- Sato:      [500,  100]

# **Nivel 1**
Crea un programa con las siguientes firmas:

*input: distancia al emisor tal cual se recibe en cada satélite
output: las coordenadas 'x' e 'y' del emisor del mensaje*

**func GetLocation(distances ...float32) (x, y float32)**

*input: el mensaje tal cual se recibe en cada satélite
output: el mensaje tal cual lo genera el emisor del mensaje*

**func GetMenssage(messages ...[]string) (msg string)**
