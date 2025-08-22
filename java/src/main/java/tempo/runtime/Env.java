package tempo.runtime;

import java.util.Arrays;
import java.util.HashMap;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.Future;

/**
 * Hello world!
 */
public class Env implements Cloneable {
    private final Transport trans;
    private HashMap<String, String> roleSubst;

    public Env(Transport transport) {
        this.trans = transport;
        this.roleSubst = new HashMap<>();
    }

    public Env subst(String... roles) {
        throw new RuntimeException("not implemented yet");
    }

    public <T> Future<T> send(T value, String... roles) {
        this.trans.send(value, this.substituteRoles(roles));
        return CompletableFuture.completedFuture(value);
    }

    public <T> Future<T> recv(String role) {
        String subRole = this.role(role);
        return this.trans.recv(subRole);
    }

    public String[] substituteRoles(String[] roles) {
        return Arrays
                .stream(roles)
                .map(r -> this.role(r))
                .toArray(String[]::new);
    }

    public String role(String name) {
        String role = this.roleSubst.get(name);
        if (role != null) {
            return role;
        } else {
            return name;
        }
    }

    @Override
    public Env clone() {
        Env copy = new Env(trans);
        copy.roleSubst = new HashMap<>(this.roleSubst);
        return copy;
    }
}
