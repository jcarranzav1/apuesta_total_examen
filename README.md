# Resolución del Reto Técnico de Apuesta Total

Este proyecto aborda la resolución del desafío técnico propuesto por Apuesta Total, implementando una solución robusta y escalable utilizando la arquitectura hexagonal y una arquitectura de microservicios basada en el patrón Saga.


## Tecnologías Utilizadas
- **Golang:** Lenguaje de programación utilizado para implementar la lógica del negocio y los microservicios.
- **PostgreSQL:** Sistema de gestión de bases de datos relacional utilizado para almacenar datos de manera persistente.
- **GORM:** Biblioteca de mapeo objeto-relacional (ORM) para Golang, utilizada para interactuar con la base de datos PostgreSQL de manera eficiente y sencilla.

## Características Principales del Proyecto
- Implementación de la arquitectura hexagonal para una separación clara de las capas de aplicación.
- Desarrollo de microservicios independientes para cada funcionalidad específica.
- Uso del patrón Saga para gestionar transacciones distribuidas y mantener la consistencia de los datos.
- Integración de Golang con PostgreSQL mediante GORM para una interacción eficiente con la base de datos.
- Escalabilidad y mantenibilidad mejoradas gracias a la modularidad y el desacoplamiento de los componentes.
