package tempo.simulator;

import java.util.concurrent.ExecutionException;

public class ProcessExecutionException extends ExecutionException {

    public final String role;

    public ProcessExecutionException(String role, ExecutionException exception) {
        super("Process '" + role + "' threw exception: " + exception.getMessage(), exception.getCause());
        this.role = role;
    }

}
