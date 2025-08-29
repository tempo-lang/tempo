package tempo.transports;

import java.util.ArrayDeque;
import java.util.concurrent.CompletableFuture;

public class LocalChannel<T> {
    ArrayDeque<T> sendBuf;
    ArrayDeque<CompletableFuture<T>> recvBuf;

    public LocalChannel() {
        this.sendBuf = new ArrayDeque<>();
        this.recvBuf = new ArrayDeque<>();
    }

    public synchronized void send(T value) {
        CompletableFuture<T> fut = recvBuf.pollFirst();
        if (fut != null) {
            fut.complete(value);
        } else {
            this.sendBuf.add(value);
        }
    }

    public synchronized CompletableFuture<T> recv() {
        CompletableFuture<T> fut = new CompletableFuture<>();
        T value = this.sendBuf.pollFirst();
        if (value != null) {
            fut.complete(value);
        } else {
            this.recvBuf.add(fut);
        }
        return fut;
    }
}
