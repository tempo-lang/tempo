import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import static diffie_hellman.Choreography.*;
import diffie_hellman.Choreography;
import java.lang.Math;
import tempo.runtime.Env;
import tempo.simulator.Result;
import tempo.simulator.Simulator;

public class DiffieHellmanTest {
    @Test
    public void testDiffieHellmanSim() throws Exception {

        Choreography.Math math = new Choreography.Math() {
            @Override
            public Integer Exp(Env env, Integer base, Integer exp) throws Exception {
                return (int) Math.pow((double) base, (double) exp);
            }
        };

        Simulator sim = new Simulator()
                .addProcess("A", env -> {
                    return DiffieHellman_A(env, math);
                })
                .addProcess("B", env -> {
                    return DiffieHellman_B(env, math);
                });

        var result = sim.run();

        var expected = Result.builder()
                .addProcess("A", proc -> proc
                        .returns(new Secret_A(18))
                        .sends(4, "B")
                        .receives(10, "B"))
                .addProcess("B", proc -> proc
                        .returns(new Secret_B(18))
                        .sends(10, "A")
                        .receives(4, "A"))
                .build();

        assertEquals(result, expected);
    }
}
