public class Bike extends Vehicle {

    int gearCount;
    int currentGear;

    @Override
    void move() {

    }

    void ringBell() {

    }


    void gearShiftUp() {
        this.currentGear++;
    }

    void gearShiftUp(int gearCount) {
        this.currentGear += gearCount;
    }
}
