Use findhotel;

CREATE TABLE `ip_data` 
  ( 
     ip      VARCHAR(39) NOT NULL, 
     ccode   VARCHAR(3) NOT NULL, 
     country VARCHAR(50) NOT NULL, 
     city    VARCHAR(50) NOT NULL, 
     mystry  VARCHAR(50) NOT NULL,
     lat     FLOAT(10, 6) NOT NULL, 
     lon     FLOAT(10, 6) NOT NULL, 
     PRIMARY KEY (ip) 
  );