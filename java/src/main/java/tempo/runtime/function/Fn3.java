package tempo.runtime.function;

@FunctionalInterface
public interface Fn3<One, Two, Three> {
    public void call(One one, Two two, Three three) throws Exception;
}
