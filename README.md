# TFG
TFG Practical Work, repository group


#GitLab indicaciones
La estructura del repositorio contiene dependencias a los repositorios de los micro-servicios, utilizando [git submudules]([Git - Submodules](https://git-scm.com/book/en/v2/Git-Tools-Submodules))

Para clonar el repositorio ejecutamos:

```shell
git clone https://github.com/Wcamaly/TFG.git
```

para bajar los submodulos ejecutamos:

```shell
git submodule init && git submodule update
```

para visualizar los branch en los submodules ejecutamos:

```shell
git submodule status
```

una vez descargado nos posicionamos en el branch

```shell
git checkout develop
```

Si queremos cambiar el branch de todos los submodules:

```shell
git submodule foreach --recursive git pull
```

para cambiar el branch de un solo proyecto hacemos lo siguiente

```shell
cd $PROJECT-service/
git checkout $BRANCH
```

Para agregar un nuevo submodulo se ejecuta el siguiente comando:

```shell
git submodule add https://github.com/Wcamaly/$REPO_NAME
```
---
# Kubernetes Hyperladger

Despliegue de kubernetes y hyperladger en un cluster local con kind. En este repositorio podemos encontra la crecion de un cluster de kubernetes y los distintos artefactos para que hyperladger funcione. 


---
# ApiGateway

Estas es la capa de ingreso de los request que nos ayuda a tener mayor control de los servicio

---
# User Service

Servicio de usuarios y roles, nos permite crear usuarios editarlos cambios de contrasenas, asigancion de roles entre otras cosas

---
# Person Service

Servicio de Personas, ABM para almacenar informacion de las personas como individudos, con una relacion one a one con User service

---
# auth Service

Servicio para manejo de autorizacion de usuarios, control de permisos y accesos, control de session.

---
# Medical Service

Servicio para almacenar la informacion relevante de un doctor, relaciones con centros medicos, legajo medico especialidades, etc...

---
# Medical Center Service

Servicio para gestionar centros medicos, guardar informacion, dar de alta pacientes entre otras cosas

---
# Blockchain Service

Api integracion con blockchain para ejecutar chaincodes 

---
# Organization service 

Api para crear nuevas organizaciones en la blockchain 