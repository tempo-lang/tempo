package tempo.transports;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.Future;

import tempo.runtime.Transport;

public class Recorder implements Transport {
    protected Transport inner;
    protected ArrayList<SendValue> sends = new ArrayList<>();
    protected ArrayList<FutureRecvValue> receives = new ArrayList<>();

    public Recorder(Transport inner) {
        this.inner = inner;
    }

    @Override
    public <T> void send(T value, String... roles) {
        this.sends.add(new SendValue(value, Arrays.asList(roles)));
        this.inner.send(value, roles);
    }

    @Override
    public <T> Future<T> recv(String role) {
        var future = this.inner.<T>recv(role);
        this.receives.add(new FutureRecvValue(future, role));
        return future;
    }

    public List<SendValue> sendValues() {
        return Collections.unmodifiableList(this.sends);
    }

    public List<RecvValue> receivedValues() {
        ArrayList<RecvValue> result = new ArrayList<>(this.receives.size());

        for (FutureRecvValue futureVal : this.receives) {
            if (!futureVal.value.isDone()) {
                throw new IllegalStateException("Not all receive futures are completed");
            }

            try {
                result.add(new RecvValue(futureVal.value.get(), futureVal.sender));
            } catch (InterruptedException | ExecutionException e) {
                throw new RuntimeException(e);
            }
        }

        return Collections.unmodifiableList(result);
    }

    protected static record FutureRecvValue(Future<?> value, String sender) {
    }
}
