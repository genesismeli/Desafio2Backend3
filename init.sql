-- Eliminar la base de datos si existe
DROP DATABASE IF EXISTS odontologos;

-- Crear la base de datos
CREATE DATABASE odontologos;

-- Usar la base de datos
USE odontologos;

-- Crear la tabla dentista
-- Crear la tabla dentista
CREATE TABLE IF NOT EXISTS dentistas (
  id INT NOT NULL AUTO_INCREMENT,
  matricula VARCHAR(50) NOT NULL,
  apellido VARCHAR(100) NOT NULL,
  nombre VARCHAR(100) NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Crear un índice en la columna "matricula" en la tabla "dentistas"
CREATE INDEX idx_dentistas_matricula ON dentistas (matricula);

-- Modificar la tabla dentistas para agregar un índice único a la columna "id"
ALTER TABLE dentistas
ADD UNIQUE INDEX dentistas_id_unique (id);

-- Crear la tabla paciente
CREATE TABLE IF NOT EXISTS pacientes (
  id INT NOT NULL AUTO_INCREMENT,
  DNI VARCHAR(50) NOT NULL,
  nombre VARCHAR(50) NOT NULL,
  apellido VARCHAR(50) NOT NULL,
  domicilio VARCHAR(100) NOT NULL,
  fecha_alta DATE NOT NULL,  
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Crear un índice en la columna "DNI" en la tabla "pacientes"
CREATE INDEX idx_pacientes_DNI ON pacientes (DNI);


-- Crear la tabla turno
CREATE TABLE IF NOT EXISTS turnos (
  id INT NOT NULL AUTO_INCREMENT,
  dentista_matricula VARCHAR(50) NOT NULL,
  paciente_DNI VARCHAR(20) NOT NULL,
  fecha_hora DATETIME NOT NULL,
  descripcion VARCHAR(1000),
  PRIMARY KEY (id),
  FOREIGN KEY (dentista_matricula) REFERENCES dentistas(matricula) ON DELETE CASCADE  ON UPDATE CASCADE,
  FOREIGN KEY (paciente_DNI) REFERENCES pacientes(DNI) ON DELETE CASCADE  ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Insertar datos ficticios en la tabla dentistas 
INSERT INTO dentistas (matricula, apellido, nombre) 
VALUES
  ('12345', 'Pérez', 'Juan'),
  ('67890', 'González', 'María'),
  ('54321', 'López', 'Carlos');

-- Insertar datos ficticios en la tabla pacientes
INSERT INTO pacientes (DNI, nombre, apellido, domicilio, fecha_alta)
VALUES
  ('111111', 'Ana', 'Martínez', 'Calle A 123', '2023-01-15'),
  ('222222', 'Luis', 'Rodríguez', 'Calle B 456', '2022-11-30'),
  ('333333', 'Laura', 'Gómez', 'Calle C 789', '2023-03-05');

-- Insertar datos ficticios en la tabla turnos 
INSERT INTO turnos (dentista_matricula, paciente_DNI, fecha_hora)
VALUES
  ('12345', '111111', '2023-09-16 10:00:00'),
  ('67890', '222222', '2023-09-17 14:30:00'),
  ('54321', '333333', '2023-09-18 09:15:00');

  -- Insertar dos turnos adicionales en la tabla turnos con descripción
INSERT INTO turnos (dentista_matricula, paciente_DNI, fecha_hora, descripcion)
VALUES
  ('12345', '111111', '2023-09-19 11:30:00', 'Revisión anual'),
  ('67890', '222222', '2023-09-20 15:00:00', 'Limpieza y consulta');
  
  select * from pacientes;