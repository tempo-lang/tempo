package tempo.simulator;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.function.Function;

import tempo.transports.RecvValue;
import tempo.transports.SendValue;

public record Result(Object returnValue, List<SendValue> sends, List<RecvValue> receives) {

    public static Builder builder() {
        return new Builder();
    }

    public static class Builder {

        HashMap<String, Result> result = new HashMap<>();

        public Builder() {
        }

        public Map<String, Result> build() {
            return Map.copyOf(result);
        }

        public Builder addProcess(String name, Function<ProcessBuilder, ProcessBuilder> process) {
            result.put(name, process.apply(new ProcessBuilder()).build());
            return this;
        }

        public static class ProcessBuilder {
            protected Object returnValue = null;
            protected ArrayList<SendValue> sends = new ArrayList<>();
            protected ArrayList<RecvValue> receives = new ArrayList<>();

            protected ProcessBuilder() {
            }

            public Result build() {
                return new Result(returnValue, sends, receives);
            }

            public ProcessBuilder returns(Object value) {
                this.returnValue = value;
                return this;
            }

            public ProcessBuilder sends(Object value, String... receivers) {
                sends.add(new SendValue(value, Arrays.asList(receivers)));
                return this;
            }

            public ProcessBuilder receives(Object value, String sender) {
                receives.add(new RecvValue(value, sender));
                return this;
            }
        }
    }

}
