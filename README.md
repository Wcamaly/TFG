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