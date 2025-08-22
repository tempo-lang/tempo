package tempo.runtime.function;

@FunctionalInterface
public interface Fn4<One, Two, Three, Four> {
    public void call(One one, Two two, Three three, Four four);
}
