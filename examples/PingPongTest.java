import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertSame;

import java.util.List;
import java.util.Map;

import org.junit.jupiter.api.Test;

import tempo.simulator.Result;
import tempo.simulator.Simulator;
import tempo.transports.RecvValue;
import tempo.transports.SendValue;

public class PingPongTest {
    @Test
    public void testPingPongSim() throws Exception {
        Simulator sim = new Simulator()
                .addProcess("A", env -> {
                    ping_pong.Choreography.Start_A(env);
                    return null;
                })
                .addProcess("B", env -> {
                    ping_pong.Choreography.Start_B(env);
                    return null;
                });

        var result = sim.run();

        var expected = Map.<String, Result>of(
                "A", new Result(
                        null,
                        List.of(new SendValue(4, List.of("B")), new SendValue(2, List.of("B"))),
                        List.of(new RecvValue(3, "B"), new RecvValue(1, "B"))),
                "B", new Result(
                        null,
                        List.of(new SendValue(3, List.of("A")), new SendValue(1, List.of("A"))),
                        List.of(new RecvValue(4, "A"), new RecvValue(2, "A"))));

        assertEquals(result, expected);
    }
}
