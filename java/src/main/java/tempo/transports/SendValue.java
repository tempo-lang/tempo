package tempo.transports;

import java.util.List;

public record SendValue(Object value, List<String> receivers) {
}