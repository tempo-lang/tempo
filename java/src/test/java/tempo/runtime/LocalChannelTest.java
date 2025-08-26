package tempo.runtime;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;

import org.junit.jupiter.api.Test;

import tempo.transports.LocalChannel;

public class LocalChannelTest {

    @Test
    public void testChannel() throws InterruptedException, ExecutionException {
        LocalChannel<String> channel = new LocalChannel<>();

        CompletableFuture<String> value = channel.recv();
        assertFalse(value.isDone());

        channel.send("A");
        assertTrue(value.isDone());

        assertEquals(value.get(), "A");

        channel.send("B");
        CompletableFuture<String> value2 = channel.recv();
        assertTrue(value2.isDone());
        assertEquals(value2.get(), "B");
    }
}
