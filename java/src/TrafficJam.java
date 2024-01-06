import java.util.Queue;

public class TrafficJam<T extends Vehicle> {

    private Queue<T> queue;


    public void addVehicle(T vehicle) {
        queue.add(vehicle);
    }

    public int jamLength() {
        return queue.size();
    }
}
