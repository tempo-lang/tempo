package tempo.simulator;

import tempo.runtime.Env;

@FunctionalInterface
public interface Process {
    Object run(Env env) throws Exception;
}
