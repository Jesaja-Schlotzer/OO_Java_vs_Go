public abstract class Train extends Vehicle implements FuelDependent{

    int wagonCount;


    void couple() {
        this.wagonCount++;
    }
}
