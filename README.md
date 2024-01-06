# OO_Java_vs_Go

# Inhalt
- [Einleitung](#einleitung)
    - [Aufgabenstellung](#aufgabenstellung)
    - [Motivation](#Motivation)
    - [Abgrenzung](#Abgrenzung)
- [OO Konzepte](#oo-konzepte)
    - [Objekte / Klassen](#objekte--klassen)
    - [Kapselung](#kapselung)
    - [Abstraktion](#abstraktion)
    - [Polymorphie](#polymorphie)
    - [Vererbung](#vererbung)
    - [Interfaces](#interfaces)
- [Fazit](#fazit)


# Einleitung
In dieser Arbeit wird das Konzept der Objektorientierung untersucht und anhand von zwei Sprachen, Java und Go, angeschaut, inwiefern sich dieses Konzept in eine Sprache einbinden lässt und wie praktisch dieses Konzept in den jeweiligen Sprachen ist. Dafür werden die einzelnen Teile der Objektorientierung in beiden Sprachen angeschaut und verglichen, für diese Vergleiche findet sich in diesem Repository der entsprechende Code, welcher jedes einzelne Konzept beispielhaft darstellt. Am Ende wird dann ein Fazit gezogen, wie gut beide Sprachen das Konzept umsetzen und inwiefern es nützlich ist Objektorierierung in Go oder Java zu verwenden.

## Aufgabenstellung 
OO in Java versus Go

We compare inheritance in Java (nominal subtyping, virtual methods) versus Go interfaces and structural subtyping

1. Consider several Java examples (of your own choice) that make use of nominal subtyping and virtual methods

2. Re-implement these examples in Go

3. Draw a comparison between Java and Go

4. Summarize your findings in a short document (you could set up a github repo).

## Motivation
Wir haben uns für dieses Thema entschieden, da die Objektorientierung ein essentielles Konzept in der Programmierung ist und da Java und Go stark unterschiedliche Ansätze haben wie dieses Konzept implementiert werden kann. Entsprechend wollten wir uns etwas genauer mit diesem Thema beschäftigen um festzustellen inwiefern die Objektorientierung in Go ähnlich der in java ist.

## Abgrenzung


# OO Konzepte 
## Objekte / Klassen
Objekte und Klassen sind die Zentralen Bestandteile einer Objektorientieren Programmierung. Klassen sind die Vorlagen für Objekte und stellen Daten (Properties)  da, an welche ein Verhalten (Methoden) gebunden wird. Ein Objekt hingegen repräsentiert einen konkreten zustand mit konkreten daten und verhalten. Im Gegensatz zur funktionalen Programmierung ist das Konzept in objektorientierter Programmierung, dass es immer und überall einen State gibt.

TODO: Auf operatoren eingehen
### Java
### Go 

<p align="right">(<a href="#inhalt">back to top</a>)</p>



## Kapselung
Kapselung ist ein Konzept in der objektorientierten Programmierung, das es ermöglicht, den Zugriff auf die internen Zustände und Funktionalitäten eines Objekts zu kontrollieren und zu verbergen.

### Java
In Java wird Kapselung durch die Verwendung von Zugriffsmodifikatoren auf Klassen, Methoden und Variablen erreicht. Es gibt vier Arten von Zugriffsmodifikatoren:

- **public**: Mit dem Schlüsselwort public markierte Klassen, Methoden oder Variablen sind von überall aus zugänglich.

- **private**: Das Schlüsselwort private beschränkt den Zugriff auf die gekapselten Elemente auf die Klasse selbst. Diese sind nicht direkt von außerhalb der Klasse zugänglich.

- **protected**: Zugriff auf geschützte Elemente ist ähnlich wie bei private, aber auch für abgeleitete Klassen (Subklassen) zugänglich, die sich im selben Paket oder in einer anderen Paketklasse befinden.

- **package-private**: Wenn keine der oben genannten Modifikatoren verwendet wird, ist das Element standardmäßig im Paket sichtbar. Es ist nicht von außerhalb des Pakets zugänglich.

Zusätzlich existiert in Java das Konzept von Gettern und Settern, welche den Zugriff auf private Instanzvariablen ermöglichen.
Der Einsatz von Getter- und Setter-Methoden bietet außerdem mehr Kontrolle über den Zugriff auf die Variablen einer Klasse, da Eingaben beispielsweise validiert werden können. 

### Go 

<p align="right">(<a href="#inhalt">back to top</a>)</p>



## Abstraktion 
Wenn in der Objektorientierung über Abstraktion gesprochen wird, wird dies oft anhand von abstrakten Typen oder abstrakten Methoden erläutert. Abstraktion beschreibt dabei die Idee, komplexes Verhalten zu vereinfachen, beziehungsweise Details der Implementierung zu vernachlässigen. Abstraktion in der OO ist oftmals auch ein Werkzeug um Sourcecode zu designen und zu Strukturieren.

### Java
Abstraktion besitzt in Java eine sehr offensichtliche Ausprägung. Mit dem Schlüsselwort `abstract` können Typen aber auch Methoden gekennzeichnet werden. Ist ein Typ als `abstract` makiert, können keine Instanzen von diesem erstellt werden, es handelt sich also um einen strukturgebenden Typen.

Methoden können ebenfalls mit `abstract` markiert werden. Dies ist allerdings nur möglich, wenn sich die Methode innerhalb eines abstrakten Typen befindet. Eine abstrakte Methode besitzt keine Implementierung und dient wie ein abstrakter Typ lediglich als strukturierendes Element.

Im folgenden Beispiel ist die abstrakte Klasse [Vehicle](java/src/Vehicle.java) zu sehen. Die Klasse wurde als abstrakt markiert, da das Konzept eines generischen Fahrzeuges zwar Sinn ergibt, ein konkretes Fahrzeug im realen leben allerdings eine speziefischere Ausprägung (bspw. [Auto](java/src/Car.java), [Zug](java/src/Train.java), [Fahrrad](java/src/Bike.java)) besitz.

```java
public abstract class Vehicle {

    protected float topSpeed;
    protected float weight;
    protected Date manufacturingDate;

    abstract void move();
}
```
Ebenfalls fällt die `move()`-Methode auf, welche als abstrakte Methode keine Implementierung innerhalb von Vehicle besitzt. Mit diesen abstrakten Elementen wird hier Verhalten ausgedrückt, welches zu konkretisieren ist. Ein `Vehicle` kann sich bewegen (`move()`), wie dies allerdings genau aussieht ist nicht definiert.

> Diese art von Abstraktion hängt stark mit dem Konzept der Vererbung zusammen, dazu mehr im Abschnitt [Vererbung](#vererbung).
### Go 

<p align="right">(<a href="#inhalt">back to top</a>)</p>



## Polymorphie 
Polymorphie - oder die Vielseitigkeit des Codes - existiert in verschiedenen Ausprägungen:
- **Überladene Operatoren**: Beim überladen von Operatoren (bspw. +, -, * usw) fürt das Nutzen eines Operators innerhalb verschiedener Kontexte zu verschiedenem verhalten. So kann `+` im Kontext von nummerischen Typen eine Addition darstellen, während es im Kontext von Strings eine Konkatenation  darstellt.

- **Inklusionspolymorphie**: Die Inklusionspolymorphie hängt stark mit der Vererbung zusammen. Im Kern bezieht sich diese Art der Polymorphie darauf, dass sich abgeleitete Typen wie der Basistyp verhalten können. So erfolgen Methodenaufrufe basierend auf dem tatsächlichen Typen erst zur Laufzeit. Dies wird auch _dynamisches Binden_ oder _late binding_ genannt.

- **Funktionsüberladung**: Bei der Überladung einer Funktion existieren mehrere Funktionen mit demselben Bezeichner aber mit unterschiedlichen Parametern. So wird beim Aufruf einer Funktion während des Kompiliervorgangs festgelegt welche Version der Funktion aufgerufen werden soll. Dieses Konzept wird auch _early binding_ genannt. 

- **Parametrische Polymorphie (Generics)**: Generics repräsentieren Typen, deren Definitionen Typvariablen beinhalten. Generics ermöglichen so die Definition von Klassen, Interfaces und Methoden, die mit verschiedenen Datentypen arbeiten können, wodurch Typsicherheit und Wiederverwendbarkeit des Codes erhöht werden.

### Java
Java unterstützt viele Arten der Polymorphie. Zunächst die Funktionsüberladung, es ist möglich mehrere Methoden im gleichen Kontext mit selben Bezeichner zu deklarieren und zu nutzen, solange sich die Parameterliste unterscheidet. In diesem Fall entscheidet der Compiler automatisch, welche funktion aufzurufen ist

```java
// Bike.java

void gearShiftUp() {
    this.currentGear++;
}

void gearShiftUp(int gearCount) {
    this.currentGear += gearCount;
}

----------------------------------
// Aufruf:

Bike b = new Bike();
b.gearShiftUp();
b.gearShiftUp(3);
```
> Beispiel aus [Bike.java](java/src/Bike.java)

Ebenfalls ist es in Java möglich _late binding_ zu nutzen. Dafür ist allerdings Vererbung notwendig.

```java
// Train.java

void couple() {
    this.wagonCount++;
}

----------------------------------
// PassengerTrain.java

@Override
void couple() {
    // Only empty trains can couple wagons
    if(this.passengerCount == 0) {
        super.couple();
    }
}

----------------------------------
// Aufruf

Train[] trains = {new PassengerTrain(0),
                  new FreightTrain()};

trains[0].couple();
trains[1].couple();

```

> Beispiel aus [Train.java](java/src/Train.java) und [PassengerTrain.java](java/src/PassengerTrain.java)

Der hier gezeigte Code zeigt _late binding_. Dabei ist erst zur Laufzeit bekannt, welche Implementierungen der Funktion `couple()` aufgerufen werden. Das Verhalten des Basistypes wird von einem abgeleiteten Typ beeinflusst.

Auch die Generik ist seit längerer Zeit in Java möglich. Folgender Code-Ausschnitt zeigt, wie ein Stau (realisiert als _Queue_) Typen als Parameter annehmen kann, sodass bspw. ein Auto-Stau oder Zug-Stau realisiert werden kann. Es ist auch möglich den Typen einzuschränken, sodass hier nur Staus von Typen angelegt werden können, welche von Vehicle erben.

```java
// TrafficJam.java

public class TrafficJam<T extends Vehicle> {

    private Queue<T> queue;


    public void addVehicle(T vehicle) {
        queue.add(vehicle);
    }

    public int jamLength() {
        return queue.size();
    }
}

----------------------------------
// Aufruf
TrafficJam<Car> carJam = new TrafficJam<>();
carJam.addVehicle(new Car());
carJam.addVehicle(new Bike()); // compile time error

```
> Beispiel aus [TrafficJam.java](java/src/TrafficJam.java)


Zuletzt ist die Polymorphie bezüglich der Operatoren. Diese existiert teilweise in Java. Der `+`-Operator beispielsweise fungiert abhängig der Typen einmal als Operator für die Addition, allerdings auch im Kontext von Strings als Konkatenation. Ein eigen definiertes Verhalten von Operatoren existiert in Java jedoch nicht.

### Go 

<p align="right">(<a href="#inhalt">back to top</a>)</p>



## Vererbung 
TODO: 
- Java (nominal)
- Go (structual)
### Java
### Go 

<p align="right">(<a href="#inhalt">back to top</a>)</p>



## Interfaces  
### Java
### Go 

<p align="right">(<a href="#inhalt">back to top</a>)</p>

# Fazit
TODO:
- Welche Konzepte funktionieren in Go besonders gut / schlecht? 
