package tempo.simulator;

import java.util.List;

import tempo.transports.RecvValue;
import tempo.transports.SendValue;

public record Result(Object returnValue, List<SendValue> sends, List<RecvValue> receives) {
}
