# OO_Java_vs_Go


## Einleitung
### Aufgabenstellung 
### Motivation
### Abgrenzung

 
## OO Konzepte 
### Objekte / Klassen
Objekte und Klassen sind die Zentralen Bestandteile einer Objektorientieren Programmierung. Klassen sind die Vorlagen für Objekte und stellen Daten (Properties)  da, an welche ein Verhalten (Methoden) gebunden wird. Ein Objekt hingegen repräsentiert einen konkreten zustand mit konkreten daten und verhalten. Im Gegensatz zur funktionalen Programmierung ist das Konzept in objektorientierter Programmierung, dass es immer und überall einen State gibt.

TODO: Auf operatoren eingehen
#### Java
#### Go 

### Kapselung
Kapselung ist ein Konzept in der objektorientierten Programmierung, das es ermöglicht, den Zugriff auf die internen Zustände und Funktionalitäten eines Objekts zu kontrollieren und zu verbergen.

#### Java
In Java wird Kapselung durch die Verwendung von Zugriffsmodifikatoren auf Klassen, Methoden und Variablen erreicht. Es gibt vier Arten von Zugriffsmodifikatoren:

- **public**: Mit dem Schlüsselwort public markierte Klassen, Methoden oder Variablen sind von überall aus zugänglich.

- **private**: Das Schlüsselwort private beschränkt den Zugriff auf die gekapselten Elemente auf die Klasse selbst. Diese sind nicht direkt von außerhalb der Klasse zugänglich.

- **protected**: Zugriff auf geschützte Elemente ist ähnlich wie bei private, aber auch für abgeleitete Klassen (Subklassen) zugänglich, die sich im selben Paket oder in einer anderen Paketklasse befinden.

- **package-private**: Wenn keine der oben genannten Modifikatoren verwendet wird, ist das Element standardmäßig im Paket sichtbar. Es ist nicht von außerhalb des Pakets zugänglich.

Zusätzlich existiert in Java das Konzept von Gettern und Settern, welche den Zugriff auf private Instanzvariablen ermöglichen.
Der Einsatz von Getter- und Setter-Methoden bietet außerdem mehr Kontrolle über den Zugriff auf die Variablen einer Klasse, da Eingaben beispielsweise validiert werden können. 

#### Go 


### Abstraktion 
#### Java
#### Go 


### Polymorphismus 
TODO: 
 - databinding? 
 - überladen
#### Java
#### Go 


### Vererbung 
TODO: 
- Java (nominal)
- Go (structual)
#### Java
#### Go 


### Interfaces  
#### Java
#### Go 

## Fazit
TODO:
- Welche Konzepte funktionieren in Go besonders gut / schlecht? 


