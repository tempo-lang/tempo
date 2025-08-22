package tempo.runtime.function;

@FunctionalInterface
public interface Fn2<One, Two> {
    public void call(One one, Two two);
}
