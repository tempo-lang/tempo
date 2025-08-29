package tempo.transports;

import java.util.HashMap;

import tempo.runtime.Transport;

public class LocalQueue {
    protected HashMap<String, LocalChannel<Object>> channels;

    public LocalQueue() {
        this.channels = new HashMap<>();
    }

    public synchronized LocalChannel<Object> get(String from, String to) {
        String key = from + "." + to;

        var chan = this.channels.get(key);
        if (chan != null) {
            return chan;
        } else {
            LocalChannel<Object> newChan = new LocalChannel<>();
            this.channels.put(key, newChan);
            return newChan;
        }
    }

    public Transport role(String role) {
        return new LocalTransport(this, role);
    }
}
