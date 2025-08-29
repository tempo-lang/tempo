import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import tempo.simulator.Result;
import tempo.simulator.Simulator;

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

        var expected = Result.builder()
                .addProcess("A", proc -> proc
                        .sends(4, "B").sends(2, "B")
                        .receives(3, "B").receives(1, "B"))
                .addProcess("B", proc -> proc
                        .sends(3, "A").sends(1, "A")
                        .receives(4, "A").receives(2, "A"))
                .build();

        assertEquals(result, expected);
    }
}
