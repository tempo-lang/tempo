package tempo.runtime;

import java.util.concurrent.Future;

public interface Transport {
    <T> void send(T value, String... roles);

    <T> Future<T> recv(String role);
}
