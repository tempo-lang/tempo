package tempo.transports;

import java.util.concurrent.Future;

import tempo.runtime.Transport;

final class LocalTransport implements Transport {
    String role;
    LocalQueue queue;

    protected LocalTransport(LocalQueue queue, String role) {
        this.role = role;
        this.queue = queue;
    }

    @Override
    public <T> void send(T value, String... roles) {
        for (String to : roles) {
            this.queue.get(role, to).send(value);
        }
    }

    @Override
    public <T> Future<T> recv(String from) {
        return this.queue.get(from, role).recv().thenApply(value -> {
            @SuppressWarnings("unchecked")
            T result = (T) value;
            return result;
        });
    }

}
