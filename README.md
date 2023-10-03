#Sesion handler

Manejadores Disponibles
1. Hello
Un manejador de ejemplo que devuelve un mensaje de saludo.

Ruta: /hello
Método: GET
Respuesta: Retorna el mensaje "Hello sebas".
2. GetTasks
Este manejador obtiene todas las tareas existentes en la base de datos.

Ruta: /tasks
Método: GET
Respuesta: Retorna un array JSON que contiene todas las tareas almacenadas en la base de datos.
3. GetSoloTasks
Este manejador obtiene una tarea específica por su ID.

Ruta: /tasks/:id
Método: GET
Parámetro: id (ID de la tarea)
Respuesta: Retorna un objeto JSON que representa la tarea correspondiente al ID proporcionado.
4. CreateTask
Este manejador crea una nueva tarea utilizando los datos proporcionados en el cuerpo de la solicitud.

Ruta: /tasks
Método: POST
Datos requeridos: body (contenido de la tarea)
Respuesta: Retorna un objeto JSON que representa la tarea creada.
5. EditTask
Este manejador actualiza una tarea existente por su ID, utilizando los datos proporcionados en el cuerpo de la solicitud.

Ruta: /tasks/:id
Método: PUT
Parámetro: id (ID de la tarea a editar)
Datos requeridos: body (nuevos datos de la tarea)
Respuesta: Retorna un objeto JSON que representa la tarea editada.
6. DeleteTasks
Este manejador elimina una tarea existente por su ID.

Ruta: /tasks/:id
Método: DELETE
Parámetro: id (ID de la tarea a eliminar)
Respuesta: Retorna un objeto JSON con un mensaje indicando que la tarea ha sido eliminada.
