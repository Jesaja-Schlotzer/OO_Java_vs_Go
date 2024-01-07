public class Car extends Vehicle implements FuelDependent{

    int doorCount;
    int passengerCount;
    private float fuelLevel = 0;


    public Car(int doorCount, int passengerCount) {
        this.doorCount = doorCount;
        this.passengerCount = passengerCount;
    }

    @Override
    public void refuel(float liter) {

    }

    @Override
    void move() {

    }



}
