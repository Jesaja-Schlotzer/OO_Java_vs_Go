# OO_Java_vs_Go

# Inhalt
- [Einleitung](#einleitung)
    - [Aufgabenstellung](#aufgabenstellung)
    - [Motivation](#Motivation)
- [OO Konzepte](#oo-konzepte)
    - [Objekte / Klassen](#objekte--klassen)
    - [Vererbung](#vererbung)
    - [Abstraktion](#abstraktion)
    - [Interfaces](#interfaces)
    - [Kapselung](#kapselung)
    - [Polymorphie](#polymorphie)
- [Fazit](#fazit)


# Einleitung
In dieser Arbeit wird das Konzept der Objektorientierung untersucht und anhand von zwei Sprachen, Java und Go, angeschaut, inwiefern sich dieses Konzept in eine Sprache einbinden lässt und wie praktisch dieses Konzept in den jeweiligen Sprachen ist. Dafür werden die einzelnen Teile der Objektorientierung in beiden Sprachen angeschaut und verglichen, für diese Vergleiche findet sich in diesem Repository der entsprechende Code, welcher jedes einzelne Konzept beispielhaft darstellt. Am Ende wird dann ein Fazit gezogen, wie gut beide Sprachen das Konzept umsetzen und inwiefern es nützlich ist Objektorierierung in Go oder Java zu verwenden.

## Aufgabenstellung 
OO in Java versus Go

We compare inheritance in Java (nominal subtyping, virtual methods) versus Go interfaces and structural subtyping

1. Consider several Java examples (of your own choice) that make use of nominal subtyping and virtual methods

2. Re-implement these examples in Go

3. Draw a comparison between Java and Go

4. Summarize your findings in a short document (you could set up a GitHub repo).

## Motivation
Wir haben uns für dieses Thema entschieden, da die Objektorientierung ein essenzielles Konzept in der Programmierung ist und da Java und Go stark unterschiedliche Ansätze haben wie dieses Konzept implementiert werden kann. Entsprechend wollten wir uns etwas genauer mit diesem Thema beschäftigen, um die Ähnlichkeiten der objektorientierten Programmierung in Go und Java festzustellen.


# OO Konzepte 
## Objekte / Klassen
Objekte und Klassen sind die Zentralen Bestandteile der objektorientieren Programmierung. Klassen sind Vorlagen für Objekte und stellen Daten (Properties) dar, an welche ein Verhalten (Methoden) gebunden wird. Ein Objekt hingegen repräsentiert einen konkreten Zustand mit konkreten Daten und Verhalten. Im Gegensatz zur funktionalen Programmierung ist das Konzept in objektorientierter Programmierung, dass es immer und überall einen State gibt.

### Java
In Java sollte der Konvention nach jede Klasse eine eigene Datei bekommen, welchen den Namen der Klasse trägt. Außerdem ist in Java fast alles ein Object, die Ausnahme stellen die primitiven Datentypen wie ``int`` oder ``boolean`` dar.

### Go 
In Go gibt es keine Klassen wie in Java, allerdings gibt es Structs welche die Funktionalität von Klassen imitieren können. Entsprechend gibt es auch keine klassischen Methoden, man kann diese allerdings mithilfe von Funktionen und einem speziellen Receiver ebenfalls imitieren.

```go
type Vehicle struct {
	topSpeed          float32
	weight            float32
	manufacturingDate time.Time
}

// Move is an abstract method in Go
func (v *Vehicle) move() {
	fmt.Println("Moving...")
}
```

<p align="right">(<a href="#inhalt">back to top</a>)</p>

## Vererbung
Vererbung beschreibt ein wichtiges Konzept der Objektorientierung bei dem ein Typ die Eigenschaften und Verhalten von einem anderen Typen erbt.
Dies ermöglicht eine effiziente und strukturierte Art, Code zu organisieren, zu erweitern und wiederzuverwenden. Dadurch wird die Entwicklung erleichtert und die Wartbarkeit verbessert.

### Java
Java nutzt nominales Subtyping, die Vererbungshierarchie muss dementsprechend explizit angegeben werden. Im unteren Beispiel erbt die Klasse ``Car`` von der Klasse `Vehicle` und implementiert gleichzeitig das Interface `FuelDependent`. Es besteht also explizit eine Hierarchie zu `Vehicle` und `FuelDependent`. Nur Klassen, die von `Vehicle` erben, erben dessen Funktionalität, gleiches gilt für das Interface `FuelDependent`. Dadurch kann ein Objekt der Klasse ``Car`` in einem Kontext verwendet werden, wo eigentlich ein `Vehicle` erwartet werden würde.

````java
public class Car extends Vehicle implements FuelDependent {

    int doorCount;
    int passengerCount;
    private float fuelLevel = 0;
    
    //...

  @Override
  public void refuel(float liter) {

  }

  @Override
  void move() {

  }
}
````

### Go
Go nutzt strukturelles Typing. Damit ist gemeint, dass ein Typ genau dann von einem Interface erbt, wenn alle Methoden des Interfaces im jeweiligen Typ implementiert sind.

````go
type FuelDependent interface {
    refuel(liter float32)
}

// Implementing refuel method for PassengerTrain (required by FuelDependent interface)
func (pt *PassengerTrain) refuel(liter float32) {
    // ...
}
````

Typen werden hierbei anhand ihrer Eigenschaften und nicht durch explizite Vererbungshierarchien verglichen. Durch Komposition können Strukturen in Go ähnliche Effekte wie Vererbung erzielen, indem sie Eigenschaften anderer Strukturen nutzen, solange sie diese gemeinsamen Merkmale teilen. Hierarchien sind nicht notwendig – das strukturelle Typsystem erlaubt flexibles Arbeiten mit Typen basierend auf geteilten Eigenschaften, unabhängig von einer Vererbungsbeziehung. Mehr zum strukturellen Typing in [Interfaces](#interfaces).

```go
type Car struct {
Vehicle        // Embedding Vehicle to extend its fields and methods
doorCount      int
passengerCount int
fuelLevel      float32
}
```

Mit Komposition sieht eine Objekterstellung folgendermaßen aus:

````go
var car = Car{
		Vehicle: Vehicle{
			topSpeed:          160,
			weight:            2000,
			manufacturingDate: time.Now(),
		},
		doorCount:      4,
		passengerCount: 4,
		fuelLevel:      0,
	}
````

<p align="right">(<a href="#inhalt">back to top</a>)</p>


## Abstraktion 
Wenn in der Objektorientierung über Abstraktion gesprochen wird, wird dies oft anhand von abstrakten Typen oder abstrakten Methoden erläutert. Abstraktion beschreibt dabei die Idee, komplexes Verhalten zu vereinfachen, beziehungsweise Details der Implementierung zu vernachlässigen. Abstraktion in der OO ist oftmals auch ein Werkzeug um Sourcecode zu designen und zu strukturieren.

### Java
Abstraktion besitzt in Java eine sehr offensichtliche Ausprägung. Mit dem Schlüsselwort `abstract` können Typen aber auch Methoden gekennzeichnet werden. Ist ein Typ als `abstract` markiert, können keine Instanzen von diesem erstellt werden, es handelt sich also um einen strukturgebenden Typen.

Methoden können ebenfalls mit `abstract` markiert werden. Dies ist allerdings nur möglich, wenn sich die Methode innerhalb eines abstrakten Typen befindet. Eine abstrakte Methode besitzt keine Implementierung und dient wie ein abstrakter Typ lediglich als strukturierendes Element.

Im folgenden Beispiel ist die abstrakte Klasse [Vehicle](java/src/Vehicle.java) zu sehen. Die Klasse wurde als abstrakt markiert, da das Konzept eines generischen Fahrzeuges zwar Sinn ergibt, ein konkretes Fahrzeug im realen Leben allerdings eine spezifischere Ausprägung (bspw. [Auto](java/src/Car.java), [Zug](java/src/Train.java), [Fahrrad](java/src/Bike.java)) besitz.

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
Go bietet keine expliziten Schlüsselwörter wie ``abstract`` an um Abstraction zu ermöglichen. Um ein Verhalten ähnlich zu objektorientierten Sprachen wie Java zu erhalten werden in Go Interfaces benutzt. Diese funktionieren ähnlich wie Interfaces in Java. Im Kapitel [Interfaces](#interfaces) wird darauf näher eingegangen.

<p align="right">(<a href="#inhalt">back to top</a>)</p>



## Interfaces
Interfaces ermöglichen es Methoden zu definieren, die eine implementierende Klasse auf jeden Fall haben muss. In unserem Code gibt es zum Beispiel das Interface ``FuelDependent`` welches von der Klasse `Car` und `Train`, aber nicht von der Klasse `Bike` implementiert wird. Klassen, die dieses Interface implementieren, benötigen Treibstoff zur Fortbewegung. 
Eine andere Möglichkeit ein solches Verhalten zu implementieren wäre gewesen, wenn wir eine Superklasse von `Vehicle` `FuelDependentVehicle` erstellt hätten, welche dann die Funktionalität von `FuelDependent` übernimmt. Wollen wir dann aber zum Beispiel auch E-Bike-Objekte einer neuen von `Bike` erbenden `EBike` Klasse erzeugen hätten wir ein Problem. Das E-Bike benötigt ebenso wie `Car` und `Train` Treibstoff zum Fahren, weshalb wir dann entweder `EBike` nicht von `Bike` sondern von `FuelDependentVehicle` erben lassen müssten oder den Code aus `FuelDependentVehicle` in `EBike` duplizieren müssten. Beides keine optimalen Lösungen.
Mit dem Interface `FuelDependent` haben wir dieses Problem nicht, wir können einfach alle relevanten Klassen vom Interface `FuelDependent` erben lassen.



### Java
In Java werden Interfaces sehr ähnlich zu Klassen definiert. Auch Interfaces bekommen der Konvention nach eine eigene Datei mit dem Namen des Interfaces. Im Gegensatz zu Klassen sind alle Methoden eines Interfaces abstrakt, haben also keinen Methodenrumpf (Eine Ausnahme dazu stellen statische und default Methoden dar). Alle Variablen eines Interfaces sind implizit statisch.

````java
public interface FuelDependent {
    void refuel(float liter);
}

````

Will man, dass eine Klasse ein Interface implementiert, so muss man das explizit angeben. Auch hier sieht man das Java nominales Subtyping einsetzt.
````java
public class Car extends Vehicle implements FuelDependent {

    //...

    @Override
    public void refuel(float liter) {
        
    }

    //...
}
````



### Go
Interfaces spielen in Go eine sehr wichtige Rolle, wenn es um Objektorientierung geht und liefern eine Methode wie man einige der Grundfunktionen emulieren kann. Ein Interface in Go ist eine Kollektion von Methodensignaturen. Da Go structural subtyping verwendet, implementiert ein Typ ein Interface immer dann, wenn es alle Methoden des Interfaces verwendet. Es muss also nicht manuell angegeben werden, dass das Interface implementiert wird, wie es in Java der Fall ist. 
```go
type FuelDependent interface {
	refuel(liter float32)
}

// PassengerTrain is a type in Go that extends Train
type PassengerTrain struct {
	Train          // Embedding Train to extend its fields and methods
	passengerCount int
}

// Implementing refuel method for PassengerTrain (required by FuelDependent interface)
func (pt *PassengerTrain) refuel(liter float32) {
	// Implement refuel logic here
}
// 
```
Wie man am obigen Beispiel sieht, wird das Interface nie explizit vom Typ ``PassengerTrain`` implementiert. Dennoch wird, dadurch dass die Methode "refuel" implementiert wird, auch das Interface implementiert. Dies wird oft auch als "duck typing" bezeichnet: wenn es aussieht wie eine Ente und sich verhält wie eine Ente, dann wird es auch als Ente wahrgenommen. Dadurch kann man in Go flexibler mit Interfaces umgehen, als es in Sprachen mit nominellem Subtyping möglich ist. 


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
In Go gibt es keine expliziten Zugriffsmodifikatoren, diese werden durch den Variablennamen inferiert. Ist der Variablenname großgeschrieben, so ist die Variable public. Ist der Name kleingeschrieben, so ist die Variable private. 
Die korrekten Bezeichnungen dafür sind exported und unexported identifier.




<p align="right">(<a href="#inhalt">back to top</a>)</p>


## Polymorphie 
Polymorphie - oder die Vielseitigkeit des Codes - existiert in verschiedenen Ausprägungen:
- **Überladene Operatoren**: Beim Überladen von Operatoren (bspw. +, -, * usw) führt das Nutzen eines Operators innerhalb verschiedener Kontexte zu verschiedenem Verhalten. So kann `+` im Kontext von nummerischen Typen eine Addition darstellen, während es im Kontext von Strings eine Konkatenation darstellt.

- **Inklusionspolymorphie**: Die Inklusionspolymorphie hängt stark mit der Vererbung zusammen. Im Kern bezieht sich diese Art der Polymorphie darauf, dass sich abgeleitete Typen wie der Basistyp verhalten können. So erfolgen Methodenaufrufe basierend auf dem tatsächlichen Typen erst zur Laufzeit. Dies wird auch _dynamisches Binden_ oder _late binding_ genannt.

- **Funktionsüberladung**: Bei der Überladung einer Funktion existieren mehrere Funktionen mit demselben Bezeichner aber mit unterschiedlichen Parametern. So wird beim Aufruf einer Funktion während des Kompiliervorgangs festgelegt welche Version der Funktion aufgerufen werden soll. Dieses Konzept wird auch _early binding_ genannt. 

- **Parametrische Polymorphie (Generics)**: Generics repräsentieren Typen, deren Definitionen Typvariablen beinhalten. Generics ermöglichen so die Definition von Klassen, Interfaces und Methoden, die mit verschiedenen Datentypen arbeiten können, wodurch Typsicherheit und Wiederverwendbarkeit des Codes erhöht werden.

### Java
Java unterstützt viele Arten der Polymorphie. Zunächst die Funktionsüberladung, es ist möglich mehrere Methoden im gleichen Kontext mit selben Bezeichner zu deklarieren und zu nutzen, solange sich die Parameterliste unterscheidet. In diesem Fall entscheidet der Compiler automatisch, welche Funktion aufzurufen ist

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

Der hier gezeigte Code zeigt _late binding_. Dabei ist erst zur Laufzeit bekannt, welche Implementierung der Funktion `couple()` aufgerufen wird. Das Verhalten des Basistypes wird von einem abgeleiteten Typ beeinflusst.

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

Go unterstützt ebenfalls verschiedene Formen der Polymorphie, eine klassische Vererbung welche für Inklusionspolymorphie benötigt wird nicht unterstützt. Es handelt sich in Go vielmehr um das Einbetten (Komposition), um ähnliche Effekte wie bei einer klassischen Vererbung zu erzielen.

Strukturen können Eigenschaften anderer Strukturen einbetten und diese verwenden. Methoden werden statisch gebunden, basierend auf dem statischen Typen der Variablen zur Kompilierzeit. Mehr zum Thema Vererbung in Go kann unter [Vererbung](#vererbung) nachgelesen werden.


Go unterstützt ebenfalls keine Funktionenüberladung. Es erlaubt nicht, mehrere Funktionen desselben Namens mit unterschiedlichen Parametern zu definieren. Funktionen mit demselben Namen sind nicht zulässig, selbst wenn sie unterschiedliche Parameter haben. So würde das vorher beschriebene Beispiel in go wie folgt aussehen: 

```go
func (b *Bike) gearShiftUp() {
	b.currentGear++
}

// Funktion muss einen anderen Namen besitzen, auch wenn sich die Parameter unterscheiden
func (b *Bike) gearShiftUpWithCount(gearCount int) {
	b.currentGear += gearCount
}
```

Generics werden allerdings von Go unterstützt und bieten damit all die vorher genannten Vorteile. 
In Go gibt es im Gegensatz zu Java allerdings keine direkte Möglichkeit, Type Constraints für spezifische Teile einer Vererbungshierarchie festzulegen. In Java ist es beispielsweise möglich, den Typen auf solche zu beschränken, die von Vehicle erben. In Go kann jedoch nur ein Interface als Typ Constraint genutzt werden, was bedeutet, dass die Einschränkung auf Schnittstellen basiert und nicht auf einer spezifischen Vererbungsbeziehung wie in Java.
 
```go
// nutzt Any anstelle von Vehicle
type TrafficJam[T any] struct {
    queue []T
}

func (t *TrafficJam[T]) addVehicle(vehicle T) {
    t.queue = append(t.queue, vehicle)
}

func (t *TrafficJam[T]) jamLength() int {
    return len(t.queue)
}
```

Sowohl in Go als auch in Java ist es nicht möglich, Operatoren manuell zu überladen. Dennoch ist der `+`-Operator in Go für verschiedene Typen definiert und somit polymorph, ähnlich zu Java.

<p align="right">(<a href="#inhalt">back to top</a>)</p>

# Fazit

Java und Go verfolgen jeweils andere Herangehensweisen im Bezug auf das objektorientierte Programmieren. Während Java von Anfang an auf die klassischen OOP-Konzepte setzte und viel explizit durch Schlüsselwörter wie `extends`, `abstract` etc. angegeben wird, setzt Go mit seinem structural Subtyping Ansatz eher darauf Vererbung und andere Konzepte der OOP implizit umzusetzen. Dadurch entstehen zwei sehr unterschiedliche Verwendungsweisen, welche beide jeweils ihre eigenen Vor- und Nachteile mit sich bringen. Während Java den Fokus oft auf Datenstrukturen legt, steht bei Go eher der Prozess im Mittelpunkt. Außerdem legt Java dem Programmierer mehr Regeln auf, während Go durch structural Subtyping freier und moderner wirkt.



