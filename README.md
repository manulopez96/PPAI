# 📄 Proyecto Práctico de Aplicación Integrador (PPAI) 2025

## 🖥️ Sistema PPAI Red Sísmica

---

## 📌 Información General

- **Materia:** Diseño de Sistemas de Información
- **Universidad:** Universidad Tecnológica Nacional – Facultad Regional Córdoba
- **Año:** 2025
- **Grupo:** G13
- **Integrantes:**
  - López, David Emanuel
  - Jorge Lorenzo, Francisco
  - Garro, Dana Brenda
  - Savarino, Sofía
  - ⁠Errigo, Mayra
  - Nieva, Agustín
  - Murua Ortiz, Facundo
  - Mazzalay, Francisco Tomas
  - Posse, Gonzalo Adonai

---

## 🎯 Objetivo del Proyecto

- Dar conocimiento a los interesados respecto de la situación sísmica en diferentes puntos.
- Emitir alertas tempranas.
- Gestionar la instalación de las estaciones sísmicas.
- Monitorear el funcionamiento de las estaciones sísmicas.
- Generar información derivada de la ocurrencia de sismos y de la instalación y mantenimiento de las sísmicas

---

## 🔍 Modelo de Requerimientos

### Objetivo y Alcance

- Gestionar la generación de información de eventos sísmicos registrados en el país y regiones colindantes.
- Dar conocimiento a los interesados respecto de la situación sísmica en diferentes puntos.
- Emitir alertas tempranas.
- Gestionar la instalación de las estaciones sísmicas.
- Gestionar el monitoreo del funcionamiento de las estaciones sísmicas.
- Generar información derivada de la ocurrencia de sismos, y de la instalación y mantenimiento de las estaciones sísmicas.

### Reglas de Negocio

- Sismógrafos para la construcción de una Estación Sismológica (ES)
- Inicio de construcción de una ES
- Reclamo por falla en sismógrafo
- Respuesta del proveedor de reclamo por falla
- Habilitación de sismógrafo
- Reemplazo de sismógrafo por falla
- Inicio de inspección a una ES
- Información de cambios de estado de sismógrafos
- Registro de resultados de inspección
- Situación de funcionamiento del sismógrafo
- Envío a reparación de sismógrafo
- Configuración de Umbrales Sísmicos
- Registro automático de eventos sísmicos
- Alerta automática de eventos sísmicos
- Eventos sísmicos no revisados
- Bloqueo de evento sísmico en revisión
- Revisión de eventos sísmicos
- Información de cambios de estado de eventos sísmicos
- Ventana temporal para detección de eventos sísmicos
- Cierre de eventos sísmicos

### Casos de Uso

- Paquete: Adm. de Usuarios
  - Registrar usuario
  - Modificar usuario
  - Consultar usuario
  - Eliminar usuario
  - Registrar perfil
  - Modificar perfil
  - Consultar perfil
  - Eliminar perfil
  - Registrar permiso
  - Modificar permiso
  - Consultar permiso
  - Eliminar permiso
  - Asignar perfiles a usuario
  - Iniciar sesión
  - Cerrar sesión
- Paquete: Gestión de Eventos Sísmicos
  - Adquirir datos de sismógrafo
  - Generar sismograma
  - Notificar variación en datos sísmicos
  - Procesar fusión de datos de estaciones sismológicas
  - Ubicar evento sísmico en mapa
  - Registrar resultado de revisión manual
  - Registrar evento sísmico
  - Modificar evento sísmico
  - Consultar evento sísmico
  - Anular evento sísmico
  - Consultar magnitud de sismo
  - Enviar notificación de ocurrencia de sismo
  - Verificar eventos sísmicos auto detectados
  - Cerrar evento sísmico
  - Registrar clasificación de sismos
  - Modificar clasificación de sismos
  - Consultar clasificación de sismos
  - Eliminar clasificación de sismos
  - Registrar origen de generación de sismos
  - Modificar origen de generación de sismos
  - Consultar origen de generación de sismos
  - Eliminar origen de generación de sismos
  - Registrar valor de magnitud Richter
  - Modificar valor de Magnitud Richter
  - Consultar valor de Magnitud Richter
  - Eliminar valor de Magnitud Richter
  - Registrar alcance de sismos
  - Modificar alcance de sismos
  - Consultar alcance de sismos
  - Eliminar alcance de sismos
  - Registrar resultado de revisión de eventos derivados
- Paquete: Adm. de ES y Sismógrafos
  - Registrar estación sismológica
  - Modificar estación sismológica
  - Consultar estación sismológica
  - Eliminar estación sismológica
  - Registrar fabricante
  - Modificar fabricante
  - Consultar fabricante
  - Eliminar fabricante
  - Registrar modelo de sismógrafo
  - Modificar modelo de sismógrafo
  - Consultar modelo de sismógrafo
  - Eliminar modelo de sismógrafo
  - Registrar sismógrafo
  - Modificar Sismógrafo
  - Consultar Sismógrafo
  - Eliminar Sismógrafo
  - Registrar Motivo Tipo
  - Modificar Motivo Tipo
  - Consultar Motivo Tipo
  - Eliminar Motivo Tipo
- Paquete: Gestión de Reportes y Estadísticas
  - Generar estadística de ocurrencia de sismos
  - Generar reporte de suscripciones
  - Generar reporte de progreso de plan de instalación
  - Generar reporte de órdenes de inspección
  - Generar informe de gestión de sismos
- Paquete: Gestión de Mantenimiento de ES
  - Dar cierre a orden de inspección de ES
  - Registrar tipo de tarea de inspección
  - Modificar tipo de tarea de inspección
  - Consultar tipo de tarea de inspección
  - Eliminar tipo de tarea de inspección
  - Registrar apreciación para tareas
  - Modificar apreciación para tareas
  - Consultar apreciación para tareas
  - Eliminar apreciación para tareas
  - Registrar diagramación de inspección de ES
  - Iniciar inspección de ES
  - Enviar a reparación un sismógrafo
  - Registrar respuesta de reparación
- Paquete: Adm. de Empleados
  - Registrar Empleado
  - Modificar Empleado
  - Consultar Empleado
  - Eliminar Empleado
  - Registrar Rol de Empleado
  - Modificar Rol de Empleado
  - Consultar Rol de Empleado
  - Eliminar Rol de Empleado
- Paquete: Gestión de Instalaciones
  - Recibir certificación de terreno
  - Generar plan de construcción de ES
  - Modificar plan de construcción de ES
  - Consultar plan de construcción de ES
  - Cancelar plan de construcción de ES
  - Registrar inicio de obra de construcción
  - Registrar avance de plan de construcción
  - Finalizar construcción de ES
  - Registrar tipo de trabajo
  - Modificar tipo de trabajo
  - Consultar tipo de trabajo
  - Eliminar tipo de trabajo
  - Generar reclamo a proveedor
  - Registrar respuesta de reclamo
  - Solicitar certificación de Terreno
- Paquete: Gestión de Suscripciones
  - Realizar suscripción para recibir novedades sobre eventos sísmicos
  - Anular suscripción para recibir novedades sobre eventos sísmicos
  - Consultar eventos sísmicos ocurridos

---

## 🧩 Modelo de Dominio

- Diagrama de clases conceptual incluido en el documento principal.


---

## 🧠 Modelo de Análisis


### Vista Estática

- Diagrama de clases con aplicación de GRASP.
- Roles como _Controlador_, _Experto_ y _Alta Cohesión_ fueron considerados.

### Vista Dinámica

- Diagramas de secuencia.
  

### Máquina de Estados

- Estados:  →

---

## 🏛️ Modelo Arquitectónico

### Estilos aplicados

- Arquitectura en capas (presentación, lógica, persistencia).
- Cliente-servidor.

### Vistas

- Vista lógica: casos de uso claves y subsistemas.
- Vista de componentes: división por módulos.
- Vista de despliegue: distribución de la app en servidor web + base de datos Postgresql.

---

## 🧩 Modelo de Diseño

### Interfaces de Usuario

- Módulo para operador.
- Reportes.
- Parámetros de configuración.

---

## 🗃️ Persistencia

- Lenguaje: Golang
- Tecnología: Web Templates
- Base de datos: Postgresql

---

## 🧪 Pruebas y Evaluación

- Flujo de casos de uso ejecutado correctamente.
- Módulo de encuestas probado con simulación de envío a 200 llamadas.
- Generación de reportes en Excel validada con datos simulados.

---


**Última actualización:** 2025-05-15
