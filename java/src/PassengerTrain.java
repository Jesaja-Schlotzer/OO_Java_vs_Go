public class PassengerTrain extends Train{

    int passengerCount;

    public PassengerTrain(int passengerCount) {
        this.passengerCount = passengerCount;
    }

    @Override
    public void refuel(float liter) {

    }

    @Override
    void move() {

    }


    @Override
    void couple() {
        // Only empty trains can couple wagons
        if(this.passengerCount == 0) {
            super.couple();
        }
    }
}
