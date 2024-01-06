public class Main {
    public static void main(String[] args) {
        PassengerTrain passengerTrain = new PassengerTrain(300);
        Bike bike = new Bike();

        Car car1 = new Car(4, 4);
        Car car2 = new Car(4, 8);
        Car car3 = new Car(2, 4);

        TrafficJam<Car> carJam = new TrafficJam<>();

        carJam.addVehicle(car1);
        carJam.addVehicle(car2);
        carJam.addVehicle(car3);
    }
}